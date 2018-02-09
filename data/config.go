package data

import (
	"strings"
)

type Config struct {
	Conn map[string]string
}

func (c Config) ConnStr() string {
	comps := []string{}
	for k, v := range c.Conn {
		comps = append(comps, k+"="+v)
	}
	return strings.Join(comps, " ")
}

func NewConfig() *Config {
	return &Config{
		Conn: map[string]string{
			"dbname":  "pxctrl",
			"sslmode": "disable",
		},
	}
}
