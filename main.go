package main

import (
	"flag"
	"log"
	"pxctrl/data"
	"pxctrl/util"
	"pxctrl/vpn"
)

//go:generate go generate pxctrl/data

func main() {
	fconfig := flag.String(
		"config", "pxctrl.config.json", "Configuration file")
	flag.Parse()

	conf := newConfig()
	if err := util.ReadJSONFile(*fconfig, &conf); err != nil {
		log.Fatalf("failed to read configuration: %s\n", err)
	}

	logger, err := util.NewLogger(conf.Log)
	if err != nil {
		log.Fatalf("failed to create logger: %s\n", err)
	}

	db, err := data.NewDB(conf.Data, logger)
	if err != nil {
		logger.Fatal("failed to open DB connection: %s\n", err)
	}
	defer db.Close()

	server := vpn.NewServer(conf.VPN, logger, db)

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("failed to start VPN session server: %s\n", err)
	}
}
