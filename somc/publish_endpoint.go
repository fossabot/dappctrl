package somc

import (
	"encoding/json"
)

type endpointParams struct {
	Channel  string          `json:"stateChannel"`
	Endpoint json.RawMessage `json:"endpoint,omitempty"`
}

// PublishEndpoint publishes an endpoint for a state channel in SOMC.
func (c *Conn) PublishEndpoint(channel string, endpoint []byte) error {
	params := endpointParams{
		Channel:  channel,
		Endpoint: endpoint,
	}

	data, err := json.Marshal(&params)
	if err != nil {
		return err
	}

	return c.request("connectionInfo", data).err
}
