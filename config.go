package main

import (
	"pxctrl/data"
	"pxctrl/serv"
)

type Config struct {
	Data *data.Config
	Serv *serv.Config
}

func NewConfig() *Config {
	return &Config{
		Data: data.NewConfig(),
		Serv: serv.NewConfig(),
	}
}
