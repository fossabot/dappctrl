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

type ChannelClose struct {
	Offering Bytes32
	Balance  Uint192
}

func (c *Conn) QueryCooperativeChannelCloses(
	client, agent Address, block uint32) ([]ChannelClose, error) {
	return nil, nil
}

func (c *Conn) QueryUncooperativeChannelCloses(
	client, agent Address, block uint32) ([]ChannelClose, error) {
	return nil, nil
}

func (c *Conn) CreateChannel(
	agent Address, offering Bytes32, deposit Uint192, auth Bytes32) error {
	return nil
}

func (c *Conn) TopUpChannel(agent Address,
	openBlock uint32, offering Bytes32, addedDeposit Uint192) error {
	return nil
}

func CloseChannelCooperatively(agent Address,
	openBlock uint32, offering Bytes32, balance Uint192,
	balanceSignature []byte, closingSignature []byte) error {
	return nil
}

func CloseChannelUncooperatively(agent Address,
	openBlock uint32, offering Bytes32, balance Uint192) error {
	return nil
}
