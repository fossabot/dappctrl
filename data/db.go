package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"pxctrl/util"
)

type DB reform.DB

func NewDB(conf *Config, logger *util.Logger) (*DB, error) {
	conn, err := sql.Open("postgres", conf.ConnStr())
	if err == nil {
		err = conn.Ping()
	}
	if err != nil {
		return nil, err
	}

	return (*DB)(reform.NewDB(conn,
		postgresql.Dialect, reform.NewPrintfLogger(logger.Info))), nil
}

func (db *DB) Close() {
	(*reform.DB)(db).DBInterface().(*sql.DB).Close()
}
