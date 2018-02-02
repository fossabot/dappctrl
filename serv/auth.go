package serv

import (
	"database/sql"
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
	return string(hash[:]) == pmnt.Password
}

func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("auth request: %s", r.RemoteAddr)

	var req AuthRequest
	if !s.parseRequest(w, r, &req) {
		return
	}

	s.logger.Info("auth payment: %s", req.PaymentID)

	var pmnt data.Payment
	pid := util.DecodeHex(req.PaymentID)
	err := s.db.FindByPrimaryKeyTo(&pmnt, pid)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Warn("payment not found: %s", req.PaymentID)
			s.reply(w, r, errorReply{accessDenied})
			return
		}
		s.logger.Error("failed to find payment: %s", err)
		s.reply(w, r, errorReply{internalFailure})
		return
	}

	if !Authenticate(&pmnt, req.PaymentPassword) {
		s.logger.Warn("payment password mismatch")
		s.reply(w, r, errorReply{accessDenied})
		return
	}

	tx, err := s.db.Begin()
	if err != nil {
		s.logger.Error("failed to begin transaction: %s", err)
		s.reply(w, r, errorReply{internalFailure})
	}

	sess := data.Session{
		ID:        util.NewUUID(),
		PaymentID: pid,
	}

	if err := tx.Insert(&sess); err != nil {
		s.logger.Error("failed to insert session: %s", err)
		s.reply(w, r, errorReply{internalFailure})
		tx.Rollback()
		return
	}

	vsess := data.VPNSession{
		SessionID: sess.ID,
	}

	if err := tx.Insert(&vsess); err != nil {
		s.logger.Error("failed to insert VPN session: %s", err)
		s.reply(w, r, errorReply{internalFailure})
		tx.Rollback()
		return
	}

	s.reply(w, r, AuthReply{sess.ID})
	tx.Commit()
}
