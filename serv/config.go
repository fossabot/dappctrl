package serv

type Config struct {
	Addr     string
	CertFile string
	KeyFile  string
	TLS      bool
}

func NewConfig() *Config {
	return &Config{
		Addr:     "localhost:8080",
		CertFile: "pxctrl.cert",
		KeyFile:  "pxctrl.key",
	}
}
