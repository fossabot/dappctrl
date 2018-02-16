package eth

type CreatedOffering struct {
	MinDeposit    Uint256
	CurrentSupply uint16
}

func (c *Conn) QueryCreatedOfferings(
	agent Address, offering Bytes32) ([]CreatedChannel, error) {
	return nil, nil
}

func (c *Conn) IsOfferingDeleted(offering Bytes32) (bool, error) {
	return false, nil
}

func (c *Conn) QueryOfferingEndpoint(
	client Address, offering Bytes32, openBlock uint32) (Bytes32, error) {
	return Bytes32{}, nil
}

func (c *Conn) QueryOfferingSupply(offering Bytes32) (uint16, error) {
	return 0, nil
}

func (c *Conn) IsOfferingPoppedUp(offering Bytes32) (bool, error) {
	return false, nil
}

func (c *Conn) SettleOffering(
	agent Address, openBlock uint32, offering Bytes32) error {
	return nil
}

type ChannelInfo struct{}

func (c *Conn) QueryChannelInfo(client, agent Address,
	openBlock uint32, offering Bytes32) (*ChannelInfo, error) {
	return nil, nil
}

func (c *Conn) PublishOfferingEndpoint(client Address,
	offering Bytes32, openBlock uint32, endpoint Bytes32) error {
	return nil
}

func (c *Conn) RegisterOffering(
	offering Bytes32, minDeposit Uint256, maxSupply uint16) error {
	return nil
}

func (c *Conn) RemoveOffering(offering Bytes32) error {
	return nil
}

func (c *Conn) PopupOffering(offering Bytes32) error {
	return nil
}
