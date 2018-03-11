package srv

import (
	"encoding/json"
	"net/http"

	"gopkg.in/reform.v1"
)

// VPN session server errors.
const (
	ErrInternalFailure      = "INTERNAL_FAILURE"
	ErrFailedToParseRequest = "FAILED_TO_PARSE_REQUEST"
	ErrAccessDenied         = "ACCESS_DENIED"
	ErrMalformedRequest     = "MALFORMED_REQUEST"
	ErrObjectNotFound       = "OBJECT_NOT_FOUND"
	ErrNonOpenChannel       = "NON_OPEN_CHANNEL"
)

type errorReply struct {
	Error string `json:"error,omitempty"`
}

func (s *Server) parseRequest(
	w http.ResponseWriter, r *http.Request, req interface{}) bool {
	s.logger.Info("server request %s from %s", r.URL, r.RemoteAddr)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Warn("failed to parse request: %s", err)
		s.reply(w, errorReply{ErrFailedToParseRequest})
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

func (s *Server) replyInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	s.reply(w, errorReply{ErrInternalFailure})
}

func (s *Server) findByPrimaryKey(w http.ResponseWriter,
	rec reform.Record, failIfNotFound bool) bool {
	err := s.db.FindByPrimaryKeyTo(rec, rec.PKPointer())
	if err == nil {
		return true
	}

	msg := "failed to find primary key %v in %v: %s"
	if failIfNotFound {
		s.logger.Error(msg, rec.PKValue(), rec.Table().Name(), err)
		s.replyInternalError(w)
	} else {
		s.logger.Warn(msg, rec.PKValue(), rec.Table().Name(), err)
		s.reply(w, errorReply{ErrObjectNotFound})
	}

	return false
}

func (s *Server) begin(w http.ResponseWriter) (*reform.TX, bool) {
	tx, err := s.db.Begin()
	if err != nil {
		s.logger.Error("failed to begin transaction: %s", err)
		s.replyInternalError(w)
		return nil, false
	}
	return tx, true
}

func (s *Server) insert(
	w http.ResponseWriter, tx *reform.TX, rec reform.Struct) bool {
	if err := tx.Insert(rec); err != nil {
		s.logger.Error("failed to insert into %s: %s",
			rec.View().Name(), err)
		s.replyInternalError(w)
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
		s.replyInternalError(w)
		tx.Rollback()
		return false
	}
	return true
}

func (s *Server) commit(w http.ResponseWriter, tx *reform.TX) bool {
	if err := tx.Commit(); err != nil {
		s.logger.Error("failed to commit transaction: %s", err)
		s.replyInternalError(w)
		tx.Rollback()
		return false
	}
	return true
}
