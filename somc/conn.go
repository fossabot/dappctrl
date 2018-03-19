package somc

import (
	"sync"

	"github.com/gorilla/websocket"

	"github.com/privatix/dappctrl/util"
)

// Config is a configuration for SOMC connection.
type Config struct {
	ReconnPeriod int
	URL          string
}

// NewConfig creates a default configuration for SOMC connection.
func NewConfig() *Config {
	return &Config{
		ReconnPeriod: 5,
		URL:          "ws://localhost",
	}
}

// Conn is a websocket connection to SOMC.
type Conn struct {
	sync.Mutex
	conf    *Config
	logger  *util.Logger
	conn    *websocket.Conn
	pending map[string]chan reply
	exit    bool
	id      uint32
}

// NewConn creates and starts a new SOMC connection.
func NewConn(conf *Config, logger *util.Logger) (*Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(conf.URL, nil)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		conf:    conf,
		logger:  logger,
		conn:    c,
		pending: make(map[string]chan reply),
	}

	go conn.handleMessages()

	return conn, nil
}

// Close closes a given SOMC connection.
func (c *Conn) Close() error {
	c.exit = true
	return c.conn.Close()
}
