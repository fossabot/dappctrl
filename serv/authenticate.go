package serv

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/sha3"
	"net/http"
	"pxctrl/data"
	"pxctrl/util"
)

type AuthRequest struct {
	PaymentID       string `json:"paymentId"`
	PaymentPassword string `json:"paymentPassword"`
}

type AuthReply struct {
	ErrorReply
	SessionID string `json:"sessionId"`
}

func authenticate(pmnt *data.Payment, pwd string) bool {
	hash := sha3.Sum256([]byte(pwd + fmt.Sprint(pmnt.Solt)))
	return base64.URLEncoding.EncodeToString(hash[:]) == pmnt.Password
}

func (s *Server) handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("payment: %s", req.PaymentID)

	pmnt := data.Payment{ID: req.PaymentID}
	if !s.findByPrimaryKey(w, &pmnt) {
		return
	}

	if !authenticate(&pmnt, req.PaymentPassword) {
		s.logger.Warn("payment password mismatch")
		w.WriteHeader(http.StatusForbidden)
		s.reply(w, ErrorReply{ErrAccessDenied})
		return
	}

	tx, ok := s.begin(w)
	if !ok {
		return
	}

	sess := data.Session{
		ID:        util.NewUUID(),
		PaymentID: req.PaymentID,
	}

	if !s.insert(w, tx, &sess) {
		return
	}

	vsess := data.VPNSession{
		SessionID: sess.ID,
	}

	if !s.insert(w, tx, &vsess) {
		return
	}

	if !s.commit(w, tx) {
		return
	}

	s.reply(w, AuthReply{SessionID: sess.ID})
}
