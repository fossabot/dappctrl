package main

import (
	"flag"
	"log"
	"pxctrl/data"
	"pxctrl/serv"
	"pxctrl/util"
)

func main() {
	fconfig := flag.String("config", "config.json", "Configuration file")
	flag.Parse()

	conf := NewConfig()
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

	server := serv.NewServer(conf.Serv, logger, db)

	if err := server.ListenAndServ(); err != nil {
		logger.Fatal("failed to listen and serve: %s\n", err)
	}
}
