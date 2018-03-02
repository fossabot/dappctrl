package contract

import (
	"errors"
	eth "pxctrl/eth/lib"
)

const (
	//
	// Contract events digests.
	// Please, see this article for the details:
	// https://codeburst.io/deep-dive-into-ethereum-logs-a8d2047c7371
	//

	EventDigestChannelCreated       = "a6153987181667023837aee39c3f1a702a16e5e146323ef10fb96844a526143c"
	EventDigestChannelToppedUp      = "392a992c1a7b756e553d8d97f43d59fafe79bc672808247debc077a6cdaba7b9"
	EthChannelCloseRequested        = "21ff66d79903f9d4ab6ab3c7c903af993e709be2ce2f4532d572925dea741cb1"
	EthOfferingCreated              = "49d573efb7cbb057727f6cadb4150ba6d5041c4fb55afe606508be636e158127"
	EthOfferingDeleted              = "21652905a07e2790c3a220d14394aee13681876bfbf38e658fa82ee5afe0c862"
	EthServiceOfferingEndpoint      = "00a7695de2bf4b4a523002334437d52e135b7a2a892d4471b5dd9005e5cd0681" // todo: This seems to be non-relevant for our tasks.
	EthServiceOfferingSupplyChanged = "1337b30376128e64c2ffd4e95d4c900b4ab42af11202b328722020216eeb46df"
	EthServiceOfferingPoppedUp      = "c8404827c21b5491a6c3dc0881307e47bfa40c3baf3d607c2d14f6bc808d4bfb"
	EthCooperativeChannelClose      = "56a4dfc7b9f93649d9142c7bef0a429decf8d3be895a3180c67a76a18d79f4ab"
	EthUncooperativeChannelClose    = "8a79bd24ee9bcfd977d6fc685befa8775c8a933f0abe82ab73b716cf419f968e"
)

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type Event interface {
	Digest() string
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventChannelCreated struct {
	Client            *eth.Address // indexed
	Agent             *eth.Address // indexed
	OfferingHash      *eth.Uint256 // indexed
	Deposit           *eth.Uint192
	AuthenticatedHash *eth.Uint256
}

func NewEventChannelCreated(topics [4]string, hexData string) (*EventChannelCreated, error) {
	var err error
	event := &EventChannelCreated{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Client, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.Agent, err = eth.NewAddress(topics[2])
	if err != nil {
		return nil, err
	}

	event.OfferingHash, err = eth.NewUint256(topics[3])
	if err != nil {
		return nil, err
	}

	// todo: implement data parsing
	return event, nil
}

func (e *EventChannelCreated) Digest() string {
	return EventDigestChannelCreated
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventChannelToppedUp struct {
	Client          *eth.Address // indexed
	Agent           *eth.Address // indexed
	OpenBlockNumber *eth.Uint256 // indexed
	OfferingHash    *eth.Uint256 // todo: seems, like this one must be indexed too
	AddedDeposit    *eth.Uint192
}

func NewEventChannelToppedUp(topics [4]string, hexData string) (*EventChannelToppedUp, error) {
	var err error
	event := &EventChannelToppedUp{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Client, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.Agent, err = eth.NewAddress(topics[2])
	if err != nil {
		return nil, err
	}

	event.OpenBlockNumber, err = eth.NewUint256(topics[3])
	if err != nil {
		return nil, err
	}

	// todo: implement data parsing
	return event, nil
}

func (e *EventChannelToppedUp) Digest() string {
	return EventDigestChannelToppedUp
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventChannelCloseRequested struct {
	Client          *eth.Address // indexed
	Agent           *eth.Address // indexed
	OpenBlockNumber *eth.Uint256 // indexed
	OfferingHash    *eth.Uint256 // todo: seems, like this one must be indexed too
	Balance         *eth.Uint192
}

func NewEventChannelCloseRequested(topics [4]string, hexData string) (*EventChannelCloseRequested, error) {
	var err error
	event := &EventChannelCloseRequested{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Client, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.Agent, err = eth.NewAddress(topics[2])
	if err != nil {
		return nil, err
	}

	event.OpenBlockNumber, err = eth.NewUint256(topics[3])
	if err != nil {
		return nil, err
	}

	// todo: implement data parsing
	return event, nil
}

func (e *EventChannelCloseRequested) Digest() string {
	return EthChannelCloseRequested
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventServiceOfferingCreated struct {
	Agent         *eth.Address // indexed
	OfferingHash  *eth.Uint256 // indexed
	MinDeposit    *eth.Uint192
	CurrentSupply *uint16
}

func NewEventServiceOfferingCreated(topics [3]string, hexData string) (*EventServiceOfferingCreated, error) {
	var err error
	event := &EventServiceOfferingCreated{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Agent, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.OfferingHash, err = eth.NewUint256(topics[2])
	if err != nil {
		return nil, err
	}

	// todo: implement data parsing
	return event, nil
}

func (e *EventServiceOfferingCreated) Digest() string {
	return EthOfferingCreated
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventServiceOfferingDeleted struct {
	OfferingHash *eth.Uint256 // indexed
}

func NewEventServiceOfferingDeleted(topics [2]string) (*EventServiceOfferingDeleted, error) {
	var err error
	event := &EventServiceOfferingDeleted{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.OfferingHash, err = eth.NewUint256(topics[1])
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventServiceOfferingDeleted) Digest() string {
	return EthOfferingDeleted
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventServiceOfferingEndpoint struct {
	Client          *eth.Address // indexed
	OfferingHash    *eth.Uint256 // indexed
	OpenBlockNumber *eth.Uint256 // indexed
	EndpointHash    *eth.Uint256
}

func NewEventServiceOfferingEndpoint(topics [4]string, hexData string) (*EventServiceOfferingEndpoint, error) {
	var err error
	event := &EventServiceOfferingEndpoint{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Client, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.OfferingHash, err = eth.NewUint256(topics[2])
	if err != nil {
		return nil, err
	}

	event.OpenBlockNumber, err = eth.NewUint256(topics[3])
	if err != nil {
		return nil, err
	}

	// todo: implement data parsing
	return event, nil
}

func (e *EventServiceOfferingEndpoint) Digest() string {
	return EthServiceOfferingEndpoint
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventServiceOfferingSupplyChanged struct {
	OfferingHash  *eth.Uint256 // indexed
	CurrentSupply *uint16
}

func NewEventServiceOfferingSupplyChanged(topics [2]string) (*EventServiceOfferingSupplyChanged, error) {
	var err error
	event := &EventServiceOfferingSupplyChanged{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.OfferingHash, err = eth.NewUint256(topics[1])
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventServiceOfferingSupplyChanged) Digest() string {
	return EthServiceOfferingSupplyChanged
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventServiceOfferingPoppedUp struct {
	OfferingHash *eth.Uint256 // indexed
}

func NewEventServiceOfferingPoppedUp(topics [2]string) (*EventServiceOfferingPoppedUp, error) {
	var err error
	event := &EventServiceOfferingPoppedUp{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.OfferingHash, err = eth.NewUint256(topics[1])
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventServiceOfferingPoppedUp) Digest() string {
	return EthServiceOfferingPoppedUp
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventCooperativeChannelClose struct {
	Client          *eth.Address // indexed
	Agent           *eth.Address // indexed
	OpenBlockNumber *eth.Uint256 // indexed
	OfferingHash    *eth.Uint256
	Balance         *eth.Uint192
}

func NewEventCooperativeChannelClose(topics [4]string, hexData string) (*EventCooperativeChannelClose, error) {
	var err error
	event := &EventCooperativeChannelClose{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Client, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.Agent, err = eth.NewAddress(topics[2])
	if err != nil {
		return nil, err
	}

	event.OpenBlockNumber, err = eth.NewUint256(topics[1])
	if err != nil {
		return nil, err
	}

	// todo: parse data
	return event, nil
}

func (e *EventCooperativeChannelClose) Digest() string {
	return EthCooperativeChannelClose
}

//---------------------------------------------------------------------------------------------------------------------

// todo: comment
type EventUncooperativeChannelClose struct {
	Client          *eth.Address // indexed
	Agent           *eth.Address // indexed
	OpenBlockNumber *eth.Uint256 // indexed
	OfferingHash    *eth.Uint256
	Balance         *eth.Uint192
}

func NewEventUnCooperativeChannelClose(topics [4]string, hexData string) (*EventUncooperativeChannelClose, error) {
	var err error
	event := &EventUncooperativeChannelClose{}

	receivedDigest := topics[0]
	if receivedDigest != event.Digest() {
		return nil, errorUnexpectedEventType(receivedDigest, event.Digest())
	}

	event.Client, err = eth.NewAddress(topics[1])
	if err != nil {
		return nil, err
	}

	event.Agent, err = eth.NewAddress(topics[2])
	if err != nil {
		return nil, err
	}

	event.OpenBlockNumber, err = eth.NewUint256(topics[1])
	if err != nil {
		return nil, err
	}

	// todo: parse data
	return event, nil
}

func (e *EventUncooperativeChannelClose) Digest() string {
	return EthUncooperativeChannelClose
}

//---------------------------------------------------------------------------------------------------------------------

func errorUnexpectedEventType(receivedDigest, expectedDigest string) error {
	return errors.New(
		"unexpected event type occurred: " + receivedDigest + ", but " + expectedDigest + " is expected")
}
