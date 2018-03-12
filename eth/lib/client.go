package lib

// Implements client for ethereum network.
//
// For details about contracts methods calling:
// https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//
// Note:
// "abigen" is available in contract/tools, no need for building it itself.
//
// For details about events (logs in EthereumConf terminology) fetching:
// https://ethereumbuilders.gitbooks.io/guide/content/en/ethereum_json_rpc.html

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type EthereumClient struct {
	Host string
	Port uint16

	client    http.Client
	requestID uint64

	privateKeyPath string
}

func NewEthereumClient(host string, port uint16) *EthereumClient {
	return &EthereumClient{
		Host: host,
		Port: port,

		// By default, standard http-client does not uses any timeout for it's operations.
		// But, there is non zero probability, that remote geth-node would hang for a long time.
		// To avoid cascade client/agent side application hangs - timeout is used.
		client: http.Client{
			Timeout: time.Second * 5,
		},
	}
}

// Base geth API response.
type apiResponse struct {
	ID      uint64 `json:"id"`
	JsonRPC string `json:"jsonrpc"`

	// All responses also contains "result" field,
	// but from method to method it might have various different types,
	// so it is delegated to the specified response types.
}

// Fills common parameters "method" and "params",
// and calls json-RPC API of the remote geth-node.
// In case of success - tries to unmarshal received data
// to the appropriate structure type ("result" argument).
//
// Tests: this is a base method for all raw API calls
// so, it is automatically covered by the all tests of all low-level methods,
// for example, GetBlockNumber(), or GetLogs().
func (e *EthereumClient) fetch(method string, params string, result interface{}) error {
	body := fmt.Sprintf(`{"jsonrpc":"2.0","method":"%s","params":[%s]}`, method, params)
	httpResponse, err := e.client.Post(e.providerURL(), "application/json", strings.NewReader(body))
	if err != nil {
		return errors.New("can't do API call: " + err.Error())
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return errors.New("can't do API call. API responded with code: " +
			fmt.Sprint(httpResponse.StatusCode))
	}

	response, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return errors.New("can't read response data: " + err.Error())
	}

	err = json.Unmarshal(response, result)
	if err != nil {
		return errors.New("can't unmarshal response data: " + err.Error())
	}

	return nil
}

func (e *EthereumClient) providerURL() string {
	return "http://" + e.Host + ":" + fmt.Sprint(e.Port)
}
