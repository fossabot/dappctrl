package main

import (
	"pxctrl/data"
	"pxctrl/serv"
	"pxctrl/util"
)

type config struct {
	Data *data.Config
	Log  *util.LogConfig
	Serv *serv.Config
}

func newConfig() *config {
	return &config{
		Data: data.NewConfig(),
		Log:  util.NewLogConfig(),
		Serv: serv.NewConfig(),
	}
}
