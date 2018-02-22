package vpn

import (
	"github.com/privatix/dappctrl/util"
	"gopkg.in/reform.v1"
	"net/http"
)

// ServerConfig is a configuration for VPN session server.
type ServerConfig struct {
	Addr     string
	CertFile string
	KeyFile  string
	TLS      bool
}

// NewServerConfig creates a default configuration for VPN session server.
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr:     "localhost:8080",
		CertFile: "dappctrl.cert",
		KeyFile:  "dappctrl.key",
		TLS:      false,
	}
}

// Server is a VPN session server.
type Server struct {
	conf   *ServerConfig
	logger *util.Logger
	db     *reform.DB
}

// NewServer creates a new VPN session server.
func NewServer(conf *ServerConfig, logger *util.Logger, db *reform.DB) *Server {
	return &Server{conf, logger, db}
}

// VPN session API paths.
const (
	PathAuth  = "/auth"
	PathStart = "/start"
	PathStop  = "/stop"
)

// ListenAndServe starts to listen and to serve requests.
func (s *Server) ListenAndServe() error {
	http.HandleFunc(PathAuth, s.handleAuth)
	http.HandleFunc(PathStart, s.handleStart)
	http.HandleFunc(PathStop, s.handleStop)

	if s.conf.TLS {
		return http.ListenAndServeTLS(
			s.conf.Addr, s.conf.CertFile, s.conf.KeyFile, nil)
	}

	return http.ListenAndServe(s.conf.Addr, nil)
}
