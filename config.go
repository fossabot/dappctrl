package main

import (
	"pxctrl/data"
	"pxctrl/serv"
	"pxctrl/util"
)

type Config struct {
	Data *data.Config
	Log  *util.LogConfig
	Serv *serv.Config
}

func NewConfig() *Config {
	return &Config{
		Data: data.NewConfig(),
		Log:  util.NewLogConfig(),
		Serv: serv.NewConfig(),
	}
}
