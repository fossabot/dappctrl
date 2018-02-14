package vpn

type Config struct {
	ServerAddr string
	CertFile   string
	KeyFile    string
	ServerTLS  bool
}

func NewConfig() *Config {
	return &Config{
		ServerAddr: "localhost:8080",
		CertFile:   "pxctrl.cert",
		KeyFile:    "pxctrl.key",
	}
}
