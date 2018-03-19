package somc

import (
	"encoding/json"
)

// WaitForEndpoint waits for a state channel endpoint data and returns it when
// it's ready. Can take a while.
func (c *Conn) WaitForEndpoint(channel string) ([]byte, error) {
	params := endpointParams{
		Channel: channel,
	}

	data, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	repl := c.request("subscribe", data)
	if repl.err != nil {
		return nil, err
	}

	c.Lock()

	ch := make(chan reply)
	pch := c.pending[channel]
	c.pending[channel] = ch

	c.Unlock()

	repl = <-ch

	// SOMC doesn't support simultaneous endpoint notifications,
	// so propagate the reply to all pending listeners.
	if pch != nil {
		pch <- repl
	}

	if repl.err != nil {
		return nil, err
	}

	return repl.data, nil
}
