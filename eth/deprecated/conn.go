package deprecated

type Address [20]byte
type Bytes32 [32]byte
type Uint192 [24]byte
type Uint256 [32]byte

type Conn struct {
}

func NewConn(conf *Config) (*Conn, error) {
	return nil, nil
}

func (c *Conn) Close() {

}
