package main

import (
	"flag"
	"pxctrl/data"
	"pxctrl/serv"
	"pxctrl/util"
)

func main() {
	fconfig := flag.String("config", "config.json", "Configuration file")
	flag.Parse()

	logger := util.NewLogger()

	var conf Config
	if err := util.ReadJSONFile(*fconfig, &conf); err != nil {
		logger.Fatal("failed to read configuration: %s\n", err)
	}

	db, err := data.NewDB(conf.Data, logger)
	if err != nil {
		logger.Fatal("failed to open DB connection: %s\n", err)
	}
	defer db.Close()

	server := serv.NewServer(conf.Serv, logger, db)

	if err := server.ListenAndServ(); err != nil {
		logger.Fatal("failed to listen and serve: %s\n", err)
	}
}
