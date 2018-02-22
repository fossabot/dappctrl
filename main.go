package main

import (
	"database/sql"
	"flag"
	"github.com/privatix/dappctrl/data"
	"github.com/privatix/dappctrl/util"
	"github.com/privatix/dappctrl/vpn"
	"log"
)

//go:generate go generate github.com/privatix/dappctrl/data

type config struct {
	DB        *data.DBConfig
	Log       *util.LogConfig
	VPNServer *vpn.ServerConfig
}

func newConfig() *config {
	return &config{
		DB:        data.NewDBConfig(),
		Log:       util.NewLogConfig(),
		VPNServer: vpn.NewServerConfig(),
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
		logger.Fatal("failed to open DB connection: %s\n", err)
	}
	defer db.DBInterface().(*sql.DB).Close()

	server := vpn.NewServer(conf.VPNServer, logger, db)

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("failed to start VPN session server: %s\n", err)
	}
}
