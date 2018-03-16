package somc

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/privatix/dappctrl/util"
)

type Config struct {
	ReconnPeriod int
	URL          string
}

type result struct {
	data []byte
	err  error
}

type Conn struct {
	sync.Mutex

	conf   *Config
	logger *util.Logger

	exit    bool
	conn    *websocket.Conn
	pending map[string]chan result
}

func NewConn(conf *Config, logger *util.Logger) (*Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(conf.URL, nil)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		conf:   conf,
		logger: logger,
		conn:   c,
	}

	go conn.handleMessages()

	return conn, nil
}

func (c *Conn) Close() error {
	c.exit = true
	return c.conn.Close()
}

func (c *Conn) handleMessages() {
	for !c.exit {
		var err error
		for {
			var msg jsonRPCMessage
			if err = c.conn.ReadJSON(&msg); err != nil {
				break
			}

			c.Lock()
			c.handleMessage(&msg)
			c.Unlock()
		}

		// Notify pending calls about error.
		c.Lock()
		for k, v := range c.pending {
			v <- result{nil, err}
			delete(c.pending, k)
		}
		c.Unlock()

		if c.exit {
			break
		}

		c.logger.Error("failed to receive SOMC response: %s", err)

		// Make sure following requests will fail.
		c.conn.Close()

		for !c.exit {
			c.conn, _, err =
				websocket.DefaultDialer.Dial(c.conf.URL, nil)
			if err == nil {
				break
			}

			c.logger.Error("failed to reconnect to SOMC: %s", err)
			time.Sleep(time.Second *
				time.Duration(c.conf.ReconnPeriod))
		}
	}
}

type jsonRPCMessage struct {
	Version string          `json:"jsonrpc"`
	ID      string          `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Result  json.RawMessage `json:"result"`
	Error   json.RawMessage `json:"error"`
}

func (c *Conn) handleMessage(m *jsonRPCMessage) {
	if len(m.Method) == 0 {
		if ch, ok := c.pending[m.ID]; ok {
			var err error
			if m.Error != nil {
				err = errors.New(
					"SOMC error: " + string(m.Error))
			}

			ch <- result{m.Result, err}
			delete(c.pending, m.ID)
		}

		return
	}

	// Endpoint message has arrived.
}

func (c *Conn) PublishOffering(o []byte) error {
	msg := jsonRPCMessage{
		ID: util.NewUUID(),
	}

	return nil
}

type OfferingData struct {
	Hash     string
	Offering []byte
}

func (c *Conn) RequestOfferings(hashes []string) ([]OfferingData, error) {
	return nil, nil
}

type EndpointData struct {
	DNS       string
	IPv4      string
	IPv6      string
	Signature string
}

func (c *Conn) RequestEndpoint(channel string) (*EndpointData, error) {
	return nil, nil
}
