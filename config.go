package main

import (
	"pxctrl/data"
	"pxctrl/util"
	"pxctrl/vpn"
)

type config struct {
	Data *data.Config
	Log  *util.LogConfig
	VPN  *vpn.Config
}

func newConfig() *config {
	return &config{
		Data: data.NewConfig(),
		Log:  util.NewLogConfig(),
		VPN:  vpn.NewConfig(),
	}
}
