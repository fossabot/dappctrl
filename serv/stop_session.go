package serv

import (
	"github.com/AlekSi/pointer"
	"net/http"
	"pxctrl/data"
	"time"
)

type StopSessionRequest struct {
	SessionID string `json:"sessionId"`
	DownKiBs  uint   `json:"downKibs"`
	UpKiBs    uint   `json:"upKibs"`
}

type StopSessionReply struct {
	ErrorReply
}

func (s *Server) handleStopSession(w http.ResponseWriter, r *http.Request) {
	var req StopSessionRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("session: %s", req.SessionID)

	sess := data.Session{ID: req.SessionID}
	if !s.findByPrimaryKey(w, &sess) {
		return
	}

	vsess := data.VPNSession{SessionID: req.SessionID}
	if !s.findByPrimaryKey(w, &sess) {
		return
	}

	sess.Ended = pointer.ToTime(time.Now())
	vsess.DownKiBs = pointer.ToUint(req.DownKiBs)
	vsess.UpKiBs = pointer.ToUint(req.UpKiBs)

	s.logger.Info("session duration: %s", sess.Ended.Sub(*sess.Started))

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

	s.reply(w, StopSessionReply{})
}
