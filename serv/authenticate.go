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
	SessionID string `json:"sessionId,omitempty"`
}

func Authenticate(pmnt *data.Payment, pwd string) bool {
	hash := sha3.Sum256([]byte(pwd + fmt.Sprint(pmnt.Solt)))
	return base64.StdEncoding.EncodeToString(hash[:]) == pmnt.Password
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

	if !Authenticate(&pmnt, req.PaymentPassword) {
		s.logger.Warn("payment password mismatch")
		s.reply(w, errorReply{accessDenied})
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

	s.reply(w, AuthReply{sess.ID})
	tx.Commit()
}
