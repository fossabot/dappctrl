package vpn

import (
	"database/sql"
	"encoding/json"
	"gopkg.in/reform.v1"
	"net/http"
	"pxctrl/data"
	"pxctrl/util"
)

const (
	PathAuthenticate = "/client/authenticate"
	PathStartSession = "/client/start-session"
	PathStopSession  = "/client/stop-session"
)

const (
	ErrInternalFailure      = "INTERNAL_FAILURE"
	ErrFailedToParseRequest = "FAILED_TO_PARSE_REQUEST"
	ErrAccessDenied         = "ACCESS_DENIED"
	ErrMalformedRequest     = "MALFORMED_REQUEST"
	ErrObjectNotFound       = "OBJECT_NOT_FOUND"
)

type Server struct {
	conf   *Config
	logger *util.Logger
	db     *data.DB
}

type ErrorReply struct {
	Error string `json:"error,omitempty"`
}

func NewServer(conf *Config, logger *util.Logger, db *data.DB) *Server {
	return &Server{conf, logger, db}
}

func (s *Server) ListenAndServe() error {
	http.HandleFunc(PathAuthenticate, s.handleAuthenticate)
	http.HandleFunc(PathStartSession, s.handleStartSession)
	http.HandleFunc(PathStopSession, s.handleStopSession)

	if s.conf.ServerTLS {
		return http.ListenAndServeTLS(
			s.conf.ServerAddr, s.conf.CertFile, s.conf.KeyFile, nil)
	} else {
		return http.ListenAndServe(s.conf.ServerAddr, nil)
	}
}

func (s *Server) parseRequest(
	w http.ResponseWriter, r *http.Request, req interface{}) bool {
	s.logger.Info("server request %s from %s", r.URL, r.RemoteAddr)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Warn("failed to parse request: %s", err)
		s.reply(w, ErrorReply{ErrFailedToParseRequest})
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
			s.reply(w, ErrorReply{ErrObjectNotFound})
		} else {
			s.logger.Error("failed to search in %v: %s",
				rec.Table().Name(), err)
			w.WriteHeader(http.StatusInternalServerError)
			s.reply(w, ErrorReply{ErrInternalFailure})
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
		s.reply(w, ErrorReply{ErrInternalFailure})
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
		s.reply(w, ErrorReply{ErrInternalFailure})
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
		s.reply(w, ErrorReply{ErrInternalFailure})
		tx.Rollback()
		return false
	}
	return true
}

func (s *Server) commit(w http.ResponseWriter, tx *reform.TX) bool {
	if err := tx.Commit(); err != nil {
		s.logger.Error("failed to commit transaction: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		s.reply(w, ErrorReply{ErrInternalFailure})
		tx.Rollback()
		return false
	}
	return true
}