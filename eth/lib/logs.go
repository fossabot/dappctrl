package lib

// This module provides low-level methods for accessing ethereum logs.
// For detailed API description, please refer to:
// https://ethereumbuilders.gitbooks.io/guide/content/en/ethereum_json_rpc.html

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LogsAPIRecord struct {
	Type                string   `json:"type"`
	TransactionIndexHex string   `json:"transactionIndex"`
	LogIndexHex         string   `json:"logIndex"`
	TransactionHash     string   `json:"transactionHash"`
	Address             string   `json:"address"`
	BlockHash           string   `json:"blockHash"`
	Data                string   `json:"data"`
	Topics              []string `json:"topics"`
	BlockNumberHex      string   `json:"blockNumber"`
}

type LogsAPIResponse struct {
	apiResponse
	Result  []LogsAPIRecord `json:"result"`
}


// Fetches logs form remote geth node.
//
// Tests: logs_test/TestNormalLogsFetching
func (e *EthereumClient) GetLogs(contractAddress string, topics []string) (*LogsAPIResponse, error){
	if contractAddress == "" {
		return nil, errors.New("contract address is required")
	}

	// note: topics are not checked for emptiness,
	// because empty topics are allowed by the geth-API:
	// in this case all events of the contract would be returned.
	for _, topic := range topics {
		const kTopicLength = 34
		if len(topic) != kTopicLength {
			return nil, errors.New("invalid topic occurred: " + topic)
		}

		if topic[:2] != "0x" {
			return nil, errors.New("invalid topic occurred: " + topic)
		}
	}

	topicsJson, err := json.Marshal(topics)
	if err != nil {
		return nil, errors.New("can't marshall topics: " + err.Error())
	}


	params := fmt.Sprintf(`{"topics":%s,"address":"%s"}`, topicsJson, contractAddress)

	response := &LogsAPIResponse{}
	err = e.fetch("eth_getLogs", params, response)
	if err != nil {
		return nil, errors.New("can't fetch response: " + err.Error())
	}

	return response, nil
}
