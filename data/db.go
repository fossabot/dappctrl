package data

import (
	"database/sql"
	// Load Go Postgres driver.
	_ "github.com/lib/pq"
	"github.com/privatix/dappctrl/util"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"strings"
)

// DBConfig is a DB configuration.
type DBConfig struct {
	Conn map[string]string
}

// ConnStr composes a data connection string.
func (c DBConfig) ConnStr() string {
	comps := []string{}
	for k, v := range c.Conn {
		comps = append(comps, k+"="+v)
	}
	return strings.Join(comps, " ")
}

// NewDBConfig creates a default DB configuration.
func NewDBConfig() *DBConfig {
	return &DBConfig{
		Conn: map[string]string{
			"dbname":  "dappctrl",
			"sslmode": "disable",
		},
	}
}

// NewDB creates a new data connection handle.
func NewDB(conf *DBConfig, logger *util.Logger) (*reform.DB, error) {
	conn, err := sql.Open("postgres", conf.ConnStr())
	if err == nil {
		err = conn.Ping()
	}
	if err != nil {
		return nil, err
	}

	return reform.NewDB(conn,
		postgresql.Dialect, reform.NewPrintfLogger(logger.Debug)), nil
}
