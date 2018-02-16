package eth

type Conn struct {
}

func NewConn(conf *Config) (*Conn, error) {
	return nil, nil
}

func (c *Conn) Close() {

}

type Address [20]byte
type Bytes32 [32]byte
type Uint192 [24]byte
type Uint256 [32]byte

type CreatedServiceOffering struct {
	MinDeposit    Uint256
	CurrentSupply uint16
}
