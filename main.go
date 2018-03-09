package main

import (
	"flag"
	"github.com/privatix/dappctrl/data"
	"github.com/privatix/dappctrl/util"
	vpnmon "github.com/privatix/dappctrl/vpn/mon"
	vpnsrv "github.com/privatix/dappctrl/vpn/srv"
	"log"
)

type config struct {
	DB         *data.DBConfig
	Log        *util.LogConfig
	VPNServer  *vpnsrv.Config
	VPNMonitor *vpnmon.Config
}

func newConfig() *config {
	return &config{
		DB:         data.NewDBConfig(),
		Log:        util.NewLogConfig(),
		VPNServer:  vpnsrv.NewConfig(),
		VPNMonitor: vpnmon.NewConfig(),
	}
}

func main() {
	fconfig := flag.String(
		"config", "dappctrl.config.json", "Configuration file")
	flag.Parse()

	conf := newConfig()
	if err := util.ReadJSONFile(*fconfig, &conf); err != nil {
		log.Fatalf("failed to read configuration: %s\n", err)
	}

	logger, err := util.NewLogger(conf.Log)
	if err != nil {
		log.Fatalf("failed to create logger: %s\n", err)
	}

	db, err := data.NewDB(conf.DB, logger)
	if err != nil {
		logger.Fatal("failed to open db connection: %s\n", err)
	}
	defer data.CloseDB(db)

	srv := vpnsrv.NewServer(conf.VPNServer, logger, db)
	defer srv.Close()
	go func() {
		logger.Fatal("failed to serve vpn session requests: %s\n",
			srv.ListenAndServe())
	}()

	mon := vpnmon.NewMonitor(conf.VPNMonitor, logger, db)
	defer mon.Close()
	go func() {
		logger.Fatal("failed to monitor vpn traffic: %s\n",
			mon.MonitorTraffic())
	}()

	<-make(chan bool)
}
