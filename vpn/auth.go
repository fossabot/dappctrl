package vpn

import (
	"encoding/base64"
	"fmt"
	"github.com/privatix/dappctrl/data"
	"github.com/privatix/dappctrl/util"
	"golang.org/x/crypto/sha3"
	"net/http"
)

// AuthRequest is an authentication request.
type AuthRequest struct {
	Channel  string `json:"channel"`
	Password string `json:"password"`
}

// AuthReply is an authentication reply.
type AuthReply struct {
	errorReply
	Session string `json:"session"`
}

func checkPassword(ch *data.Channel, pwd string) bool {
	hash := sha3.Sum256([]byte(pwd + fmt.Sprint(ch.Solt)))
	return base64.URLEncoding.EncodeToString(hash[:]) == ch.Password
}

func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("channel: %s", req.Channel)

	ch := data.Channel{ID: req.Channel}
	if !s.findByPrimaryKey(w, &ch) {
		return
	}

	if !checkPassword(&ch, req.Password) {
		s.logger.Warn("channel password mismatch")
		w.WriteHeader(http.StatusForbidden)
		s.reply(w, errorReply{ErrAccessDenied})
		return
	}

	tx, ok := s.begin(w)
	if !ok {
		return
	}

	sess := data.Session{
		ID:      util.NewUUID(),
		Channel: req.Channel,
	}

	if !s.insert(w, tx, &sess) {
		return
	}

	vsess := data.VPNSession{
		ID: sess.ID,
	}

	if !s.insert(w, tx, &vsess) {
		return
	}

	if !s.commit(w, tx) {
		return
	}

	s.reply(w, AuthReply{Session: sess.ID})
}
