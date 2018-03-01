package deprecated

type Config struct {
	Addr    Address
	PTCAddr Address
	PSCAddr Address
}

func NewConfig() *Config {
	return &Config{}
}
