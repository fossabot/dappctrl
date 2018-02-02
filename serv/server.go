package serv

import (
	"encoding/json"
	"net/http"
	"pxctrl/data"
	"pxctrl/util"
)

const (
	internalFailure      = "INTERNAL_FAILURE"
	failedToParseRequest = "FAILED_TO_PARSE_REQUEST"
	accessDenied         = "ACCESS_DENIED"
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
	http.HandleFunc("/auth", s.handleAuth)
	return http.ListenAndServe(s.conf.Addr, nil)
}

func (s *Server) parseRequest(
	w http.ResponseWriter, r *http.Request, req interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Warn("failed to parse request: %s", err)
		s.reply(w, r, errorReply{failedToParseRequest})
		return false
	}
	r.Body.Close()

	return true
}

func (s *Server) reply(
	w http.ResponseWriter, r *http.Request, rep interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(rep); err != nil {
		s.logger.Warn("failed to send reply: %s", err)
	}
}
