package data

import (
	"time"
)

//go:generate reform

//reform:services
type Service struct {
	ID string `reform:"id,pk"`
	SO string `reform:"so"`
}

//reform:vpn_services
type VPNService struct {
	ServiceID string `reform:"service_id,pk"`
}

//reform:clients
type Client struct {
	ID string `reform:"id,pk"`
}

//reform:payments
type Payment struct {
	ID        string `reform:"id,pk"`
	ServiceID string `reform:"service_id"`
	ClientID  string `reform:"client_id"`
	EthBlock  int    `reform:"eth_block"`
	Solt      int64  `reform:"solt"`
	Password  string `reform:"password"`
}

//reform:vpn_payments
type VPNPayment struct {
	PaymentID string `reform:"payment_id,pk"`
	TotalMiBs int    `reform:"total_mibs"`
}

//reform:sessions
type Session struct {
	ID        string    `reform:"id,pk"`
	PaymentID string    `reform:"payment_id"`
	Started   time.Time `reform:"started"`
	Ended     time.Time `reform:"ended"`
}

//reform:vpn_sessions
type VPNSession struct {
	SessionID    string `reform:"session_id,pk"`
	ConsumedMiBs int    `reform:"consumed_mibs"`
}
