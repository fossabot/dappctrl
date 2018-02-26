package mon

type Config struct {
	Addr       string
	BytePeriod uint
}

func NewConfig() *Config {
	return &Config{
		Addr:       "localhost:7505",
		BytePeriod: 5,
	}
}

type Monitor struct {
}
