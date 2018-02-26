package vpn

import (
	"github.com/AlekSi/pointer"
	"github.com/privatix/dappctrl/data"
	vpnutil "github.com/privatix/dappctrl/vpn/util"
	"net/http"
	"time"
)

// StopRequest is a request to stop a client session.
type StopRequest struct {
	Channel    string `json:"channel"`
	Uploaded   uint64 `json:"uploaded"`
	Downloaded uint64 `json:"downloaded"`
}

// StopReply is a reply for a session stop request.
type StopReply struct {
	errorReply
}

func (s *Server) handleStop(w http.ResponseWriter, r *http.Request) {
	var req StopRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("channel: %s", req.Channel)

	sid, err := vpnutil.FindCurrentSession(s.db, req.Channel)
	if err != nil {
		s.logger.Error("failed to find session: %s", err)
		s.replyInternalError(w)
		return
	}

	s.logger.Info("session: %s", sid)

	sess := data.Session{ID: sid}
	if !s.findByPrimaryKey(w, &sess, true) {
		return
	}

	sess.Stopped = pointer.ToTime(time.Now())

	s.logger.Info("session duration: %s", sess.Stopped.Sub(sess.Started))

	vsess := data.VPNSession{ID: sid}
	if !s.findByPrimaryKey(w, &vsess, true) {
		return
	}

	vsess.Uploaded = req.Uploaded
	vsess.Downloaded = req.Downloaded

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

	s.reply(w, StopReply{})
}
