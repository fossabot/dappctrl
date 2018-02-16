package vpn

type Config struct {
	ByteCountSecs int
	CertFile      string
	ControlAddr   string
	KeyFile       string
	ServerAddr    string
	ServerTLS     bool
}

func NewConfig() *Config {
	return &Config{
		ByteCountSecs: 5,
		CertFile:      "pxctrl.cert",
		ControlAddr:   "localhost:7505",
		KeyFile:       "pxctrl.key",
		ServerAddr:    "localhost:8080",
	}
}
