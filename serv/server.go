package serv

import (
	"database/sql"
	"encoding/json"
	"gopkg.in/reform.v1"
	"net/http"
	"pxctrl/data"
	"pxctrl/util"
)

const (
	internalFailure      = "INTERNAL_FAILURE"
	failedToParseRequest = "FAILED_TO_PARSE_REQUEST"
	accessDenied         = "ACCESS_DENIED"
	malformedRequest     = "MALFORMED_REQUEST"
	objectNotFound       = "OBJECT_NOT_FOUND"
)

type Server struct {
	conf   *Config
	logger *util.Logger
	db     *data.DB
}

type errorReply struct {
	Error string `json:"error,omitempty"`
}

func NewServer(conf *Config, logger *util.Logger, db *data.DB) *Server {
	return &Server{conf, logger, db}
}

func (s *Server) ListenAndServ() error {
	http.HandleFunc("/client/authenticate", s.handleAuthenticate)
	http.HandleFunc("/client/start-session", s.handleStartSession)
	http.HandleFunc("/client/stop-session", s.handleStopSession)
	return http.ListenAndServe(s.conf.Addr, nil)
}

func (s *Server) parseRequest(
	w http.ResponseWriter, r *http.Request, req interface{}) bool {
	s.logger.Info("server request %s from %s", r.URL, r.RemoteAddr)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Warn("failed to parse request: %s", err)
		s.reply(w, errorReply{failedToParseRequest})
		return false
	}
	r.Body.Close()

	return true
}

func (s *Server) reply(w http.ResponseWriter, rep interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(rep); err != nil {
		s.logger.Warn("failed to send reply: %s", err)
	}
}

func (s *Server) findByPrimaryKey(
	w http.ResponseWriter, rec reform.Record) bool {
	err := s.db.FindByPrimaryKeyTo(rec, rec.PKPointer())
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Warn("primary key %v not found in %s",
				rec.PKValue(), rec.Table().Name())
			s.reply(w, errorReply{objectNotFound})
		} else {
			s.logger.Error("failed to search in %v: %s",
				rec.Table().Name(), err)
			w.WriteHeader(http.StatusInternalServerError)
			s.reply(w, errorReply{internalFailure})
		}
		return false
	}
	return true
}

func (s *Server) begin(w http.ResponseWriter) (*reform.TX, bool) {
	tx, err := s.db.Begin()
	if err != nil {
		s.logger.Error("failed to begin transaction: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		s.reply(w, errorReply{internalFailure})
		return nil, false
	}
	return tx, true
}

func (s *Server) insert(
	w http.ResponseWriter, tx *reform.TX, rec reform.Struct) bool {
	if err := tx.Insert(rec); err != nil {
		s.logger.Error("failed to insert into %s: %s",
			rec.View().Name(), err)
		w.WriteHeader(http.StatusInternalServerError)
		s.reply(w, errorReply{internalFailure})
		tx.Rollback()
		return false
	}
	return true
}

func (s *Server) save(
	w http.ResponseWriter, tx *reform.TX, rec reform.Record) bool {
	if err := tx.Save(rec); err != nil {
		s.logger.Error("failed to save in %s: %s",
			rec.View().Name(), err)
		w.WriteHeader(http.StatusInternalServerError)
		s.reply(w, errorReply{internalFailure})
		tx.Rollback()
		return false
	}
	return true
}

func (s *Server) commit(w http.ResponseWriter, tx *reform.TX) bool {
	if err := tx.Commit(); err != nil {
		s.logger.Error("failed to commit transaction: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		s.reply(w, errorReply{internalFailure})
		tx.Rollback()
		return false
	}
	return true
}
