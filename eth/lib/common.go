package lib

// This module provides low-level methods for accessing common ethereum info.
// For detailed API description, please refer to:
// https://ethereumbuilders.gitbooks.io/guide/content/en/ethereum_json_rpc.html

type GasPriceAPIResponse struct {
	apiResponse
	Result string `json:"result"`
}

// Returns the current price per gas in wei.
//
// For the details, please, refer to:
// https://ethereumbuilders.gitbooks.io/guide/content/en/ethereum_json_rpc.html#eth_gasprice
//
// Tests: common_test/TestGasPriceFetching
func (e *EthereumClient) GetGasPrice() (*GasPriceAPIResponse, error) {
	response := &GasPriceAPIResponse{}
	return response, e.fetch("eth_gasPrice", "", response)
}

//---------------------------------------------------------------------------------------------------------------------

type BlockNumberAPIResponse GasPriceAPIResponse

// Returns the number of most recent block in blockchain.
//
// For the details, please, refer to:
// https://ethereumbuilders.gitbooks.io/guide/content/en/ethereum_json_rpc.html#eth_blocknumber
//
// Tests: common_test/TestBlockNumberFetching
func (e *EthereumClient) GetBlockNumber() (*BlockNumberAPIResponse, error) {
	response := &BlockNumberAPIResponse{}
	return response, e.fetch("eth_blockNumber", "", response)
}