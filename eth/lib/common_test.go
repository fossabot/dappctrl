package lib

import "testing"

func TestGasPriceFetching(t *testing.T) {
	client := NewEthereumClient(host, port)
	response, err := client.GetGasPrice()
	if err != nil {
		t.Fatal(err)
	}

	// todo: check with predicted result [truffle tests]
	if response.Result == "" {
		t.Fatal("Unexpected response received")
	}
}

func TestBlockNumberFetching(t *testing.T) {
	client := NewEthereumClient(host, port)
	response, err := client.GetBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	// todo: check with predicted result [truffle tests]
	if response.Result == "" {
		t.Fatal("Unexpected response received")
	}
}