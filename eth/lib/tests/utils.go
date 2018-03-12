package tests

import (
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)

type GethNode struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type EthereumConf struct {
	Geth GethNode `json:"geth"`
}


var (
	conf *EthereumConf = nil
)

func GethNodeConfig() *EthereumConf {
	if conf == nil {
		loadTestConfig()
	}

	return conf
}

func loadTestConfig() {
	confPath, err := filepath.Abs(filepath.FromSlash("../../dappctrl-test.config.json"))
	if err != nil {
		panic(err)
	}

	confData, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	data := make(map[string]json.RawMessage)
	err = json.Unmarshal(confData, &data)
	if err != nil {
		panic(err)
	}

	conf = &EthereumConf{}
	err = json.Unmarshal(data["EthereumLib"], conf)
	if err != nil {
		panic(err)
	}
}