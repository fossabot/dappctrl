package data

import (
	"time"
)

//go:generate reform

const (
	ServiceTypeVPN = "VPN"
)

//reform:services
type Service struct {
	ID   string `reform:"id,pk"`
	SO   string `reform:"so"`
	Type string `reform:"type"`
}

//reform:vpn_services
type VPNService struct {
	ServiceID     string `reform:"service_id,pk"`
	DownSpeedKiBs uint   `reform:"down_speed_kibs"`
	UpSpeedKiBs   uint   `reform:"up_speed_kibs"`
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
	EthBlock  uint   `reform:"eth_block"`
	Solt      uint64 `reform:"solt"`
	Password  string `reform:"password"`
}

//reform:vpn_payments
type VPNPayment struct {
	PaymentID string `reform:"payment_id,pk"`
	DownKiBs  uint   `reform:"down_kibs"`
	UpKiBs    uint   `reform:"up_kibs"`
}

//reform:sessions
type Session struct {
	ID         string     `reform:"id,pk"`
	PaymentID  string     `reform:"payment_id"`
	ServerIP   *string    `reform:"server_ip"`
	ClientIP   *string    `reform:"client_ip"`
	ClientPort *uint16    `reform:"client_port"`
	Started    *time.Time `reform:"started"`
	Ended      *time.Time `reform:"ended"`
}

//reform:vpn_sessions
type VPNSession struct {
	SessionID string `reform:"session_id,pk"`
	DownKiBs  *uint  `reform:"down_kibs"`
	UpKiBs    *uint  `reform:"up_kibs"`
}
