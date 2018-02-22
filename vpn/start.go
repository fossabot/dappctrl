package vpn

import (
	"github.com/AlekSi/pointer"
	"github.com/privatix/dappctrl/data"
	"net"
	"net/http"
	"time"
)

// StartRequest is a request to start a client session.
type StartRequest struct {
	Session    string `json:"session"`
	ServerIP   string `json:"serverIp"`
	ClientIP   string `json:"clientIp"`
	ClientPort uint16 `json:"clientPort"`
}

// StartReply is a reply for a session start request.
type StartReply struct {
	errorReply
}

func (s *Server) handleStart(w http.ResponseWriter, r *http.Request) {
	var req StartRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("session: %s", req.Session)

	sip := net.ParseIP(req.ServerIP)
	cip := net.ParseIP(req.ClientIP)

	if sip == nil || sip.IsUnspecified() ||
		cip == nil || cip.IsUnspecified() || req.ClientPort == 0 {
		s.logger.Warn("malformed request")
		s.reply(w, errorReply{ErrMalformedRequest})
		return
	}

	sess := data.Session{ID: req.Session}
	if !s.findByPrimaryKey(w, &sess) {
		return
	}

	sess.Started = pointer.ToTime(time.Now())

	vsess := data.VPNSession{ID: req.Session}
	if !s.findByPrimaryKey(w, &vsess) {
		return
	}

	s.logger.Info("session peers: %s:%d -> %s",
		req.ClientIP, req.ClientPort, req.ServerIP)

	vsess.ServerIP = pointer.ToString(req.ServerIP)
	vsess.ClientIP = pointer.ToString(req.ClientIP)
	vsess.ClientPort = pointer.ToUint16(req.ClientPort)

	tx, ok := s.begin(w)
	if !ok {
		return
	}

	if !s.save(w, tx, &sess) {
		return
	}

	if !s.save(w, tx, &vsess) {
		return
	}

	if !s.commit(w, tx) {
		return
	}

	s.reply(w, StartReply{})
}
