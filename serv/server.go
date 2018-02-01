package serv

import (
	"net/http"
	"pxctrl/data"
	"pxctrl/util"
)

type Server struct {
	conf   *Config
	logger *util.Logger
	db     *data.DB
}

func NewServer(conf *Config, logger *util.Logger, db *data.DB) *Server {
	return &Server{conf, logger, db}
}

func (s *Server) ListenAndServ() error {
	http.HandleFunc("/auth", s.handleAuth)

	return http.ListenAndServe(s.conf.Addr, nil)
}
