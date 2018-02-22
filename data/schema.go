package data

import (
	"time"
)

//go:generate reform

// Subject is a party of distributed trade.
// It can play an agent role, a client role, or both of them.
//reform:subjects
type Subject struct {
	ID         string  `reform:"id,pk"`
	PublicKey  string  `reform:"public_key"`
	PrivateKey *string `reform:"private_key"`
}

// Service types.
const (
	ServiceVPN = "vpn"
)

// Offering is a service offering.
//reform:offerings
type Offering struct {
	ID        string `reform:"id,pk"`
	Agent     string `reform:"agent"`
	Service   string `reform:"service"`
	Supply    uint   `reform:"supply"`
	Signature string `reform:"signature"`
}

// VPN protocols.
const (
	ProtocolTCP = "tcp"
	ProtocolUDP = "udp"
)

// VPNOffering is a VPN service offering.
//reform:offerings_vpn
type VPNOffering struct {
	ID               string `reform:"id,pk"`
	Hash             string `reform:"hash"`
	Country          string `reform:"country"`
	UploadPrice      string `reform:"upload_price"`
	DownloadPrice    string `reform:"download_price"`
	MinUpload        string `reform:"min_upload"`
	MaxUpload        string `reform:"max_upload"`
	MinDownload      string `reform:"min_download"`
	MaxDownload      string `reform:"max_download"`
	BillingInterval  uint   `reform:"billing_interval"`
	BillingDeviation uint   `reform:"billing_deviation"`
	FreeIntervals    uint8  `reform:"free_intervals"`
	MinUploadBPS     uint64 `reform:"min_upload_bps"`
	MinDownloadBPS   uint64 `reform:"min_download_bps"`
	TemplateVersion  uint8  `reform:"template_version"`
	Protocol         string `reform:"protocol"`
}

// Channel states.
const (
	ChannelOpen         = "open"
	ChannelClosedCoop   = "closed_coop"
	ChannelClosedUncoop = "closed_uncoop"
)

// Channel is a state channel.
//reform:channels
type Channel struct {
	ID               string `reform:"id,pk"`
	Agent            string `reform:"agent"`
	Client           string `reform:"client"`
	Offering         string `reform:"offering"`
	Block            uint   `block`
	State            string `reform:"state"`
	TotalDeposit     string `reform:"total_deposit"`
	ClosedDeposit    string `reform:"closed_deposit"`
	Solt             uint64 `reform:"solt"`
	Password         string `reform:"password"`
	ReceiptBalance   string `reform:"receipt_balance"`
	ReceiptSignature string `reform:"receipt_signature"`
}

// Session is a client session.
//reform:sessions
type Session struct {
	ID      string     `reform:"id,pk"`
	Channel string     `reform:"channel"`
	Started *time.Time `reform:"started"`
	Stopped *time.Time `reform:"stopped"`
}

// VPNSession is a client session for VPN service.
//reform:sessions_vpn
type VPNSession struct {
	ID         string  `reform:"id,pk"`
	ServerIP   *string `reform:"server_ip"`
	ClientIP   *string `reform:"client_ip"`
	ClientPort *uint16 `reform:"client_port"`
	Uploaded   uint64  `reform:"uploaded"`
	Downloaded uint64  `reform:"downloaded"`
}
