package eth

type CreatedChannel struct {
	Deposit Uint192
	Auth    Bytes32
}

func (c *Conn) QueryCreatedChannels(
	client, agent Address, offering Bytes32) ([]CreatedChannel, error) {
	return nil, nil
}

type ToppedUpChannel struct {
	Offering     Bytes32
	AddedDeposit Uint192
}

func (c *Conn) QueryToppedUpChannels(
	client, agent Address, openBlock uint32) ([]ToppedUpChannel, error) {
	return nil, nil
}

type CloseRequestedChannel struct {
	Offering Bytes32
	Ballance Uint192
}

func (c *Conn) QueryCloseRequestedChannels(client, agent Address,
	openBlock uint32) ([]CloseRequestedChannel, error) {
	return nil, nil
}
