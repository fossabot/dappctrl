package serv

type Config struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{}
}
