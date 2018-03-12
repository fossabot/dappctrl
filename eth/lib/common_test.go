package lib

import (
	"testing"
	"pxctrl/eth/lib/tests"
)


func TestGasPriceFetching(t *testing.T) {
	node := tests.GethNodeConfig().Geth
	client := NewEthereumClient(node.Host, node.Port)
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
	node := tests.GethNodeConfig().Geth
	client := NewEthereumClient(node.Host, node.Port)
	response, err := client.GetBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	// todo: check with predicted result [truffle tests]
	if response.Result == "" {
		t.Fatal("Unexpected response received")
	}
}
