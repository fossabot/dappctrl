package serv

import (
	"github.com/AlekSi/pointer"
	"net"
	"net/http"
	"pxctrl/data"
	"time"
)

type StartSessionRequest struct {
	SessionID  string `json:"sessionId"`
	ServerIP   string `json:"serverIp"`
	ClientIP   string `json:"clientIp"`
	ClientPort uint16 `json:"clientPort"`
}

type StartSessionReply struct {
	ErrorReply
}

func (s *Server) handleStartSession(w http.ResponseWriter, r *http.Request) {
	var req StartSessionRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("session: %s", req.SessionID)

	sip := net.ParseIP(req.ServerIP)
	cip := net.ParseIP(req.ClientIP)

	if sip == nil || sip.IsUnspecified() ||
		cip == nil || cip.IsUnspecified() || req.ClientPort == 0 {
		s.logger.Warn("malformed request")
		s.reply(w, ErrorReply{ErrMalformedRequest})
		return
	}

	sess := data.Session{ID: req.SessionID}
	if !s.findByPrimaryKey(w, &sess) {
		return
	}

	s.logger.Info("session peers: %s:%d -> %s",
		req.ClientIP, req.ClientPort, req.ServerIP)

	sess.ServerIP = pointer.ToString(req.ServerIP)
	sess.ClientIP = pointer.ToString(req.ClientIP)
	sess.ClientPort = pointer.ToUint16(req.ClientPort)
	sess.Started = pointer.ToTime(time.Now())

	tx, ok := s.begin(w)
	if !ok {
		return
	}

	if !s.save(w, tx, &sess) {
		return
	}

	if !s.commit(w, tx) {
		return
	}

	s.reply(w, StartSessionReply{})
}
