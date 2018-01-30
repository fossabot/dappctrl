package main

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
)

type config struct {
	DBConnStr string
}

func main() {
	fconfig := flag.String("config", "config.json", "Configuration file")
	flag.Parse()

	var conf config
	if err := readJSONFile(*fconfig, &conf); err != nil {
		log.Fatalf("failed to read configuration: %s\n", err)
	}

	conn, err := sql.Open("postgres", conf.DBConnStr)
	if err == nil {
		err = conn.Ping()
	}
	if err != nil {
		log.Fatalf("failed to open DB connection: %s\n", err)
	}
	defer conn.Close()

	db := reform.NewDB(conn,
		postgresql.Dialect, reform.NewPrintfLogger(log.Printf))

	doSomething(db)
}

func doSomething(db *reform.DB) {

}
