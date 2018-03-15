package lib

import (
	"errors"
	"fmt"
)

const (
	//
	// Contract events digests.
	// Please, see this article for the details:
	// https://codeburst.io/deep-dive-into-ethereum-logs-a8d2047c7371
	//

	EthDigestChannelCreated         = "a6153987181667023837aee39c3f1a702a16e5e146323ef10fb96844a526143c"
	EthDigestChannelToppedUp        = "392a992c1a7b756e553d8d97f43d59fafe79bc672808247debc077a6cdaba7b9"
	EthChannelCloseRequested        = "21ff66d79903f9d4ab6ab3c7c903af993e709be2ce2f4532d572925dea741cb1"
	EthOfferingCreated              = "49d573efb7cbb057727f6cadb4150ba6d5041c4fb55afe606508be636e158127"
	EthOfferingDeleted              = "21652905a07e2790c3a220d14394aee13681876bfbf38e658fa82ee5afe0c862"
	EthServiceOfferingEndpoint      = "00a7695de2bf4b4a523002334437d52e135b7a2a892d4471b5dd9005e5cd0681"
	EthServiceOfferingSupplyChanged = "1337b30376128e64c2ffd4e95d4c900b4ab42af11202b328722020216eeb46df"
	EthServiceOfferingPoppedUp      = "c8404827c21b5491a6c3dc0881307e47bfa40c3baf3d607c2d14f6bc808d4bfb"
	EthCooperativeChannelClose      = "56a4dfc7b9f93649d9142c7bef0a429decf8d3be895a3180c67a76a18d79f4ab"
	EthUncooperativeChannelClose    = "8a79bd24ee9bcfd977d6fc685befa8775c8a933f0abe82ab73b716cf419f968e"
)

//---------------------------------------------------------------------------------------------------------------------

// Base interface for all events, that are expected to be received from the ethereum block-chain.
type Event interface {
	Digest() string
}

//---------------------------------------------------------------------------------------------------------------------

// todo: here and elsewhere: To clearly deliver the meaning, I would suggest to change naming style here and in other place from NewEventServiceOfferingCreated to NewServiceOfferingCreatedEvent
type EventChannelCreated struct {
	Client            *Address // Indexed.
	Agent             *Address // Indexed.
	OfferingHash      *Uint256 // Indexed.
	Deposit           *Uint192
	AuthenticatedHash *Uint256
}

func NewEventChannelCreated(topics [4]string, hexData string) (*EventChannelCreated, error) {
	var err error
	event := &EventChannelCreated{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Client address parsing.
		clientAddressHex := topics[1][2:] // Skipping "0x".
		offsetFrom := 24                  // Skipping first 24 bytes from 64.
		event.Client, err = NewAddress(clientAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Agent address parsing.
		agentAddressHex := topics[2][2:] // Skipping "0x".
		offsetFrom := 24                 // Skipping first 24 bytes from 64.
		event.Agent, err = NewAddress(agentAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Offering hash parsing.
		offeringHashHex := topics[3][2:]
		event.OfferingHash, err = NewUint256(offeringHashHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Deposit parsing (data field).
		offsetFrom :=
			2 + // "0x..."
				16 - // topic length is 256 bits, but we need only 192 (first 16 bytes are omitted).
				1 // indexing is starting from zero.

		offsetTo := offsetFrom +
			48 // we need 48 bytes.

		event.Deposit, err = NewUint192(hexData[offsetFrom:offsetTo])
		if err != nil {
			return nil, err
		}
	}

	{
		// Authenticated hash parsing (data field).
		offsetFrom :=
			2 + // "0x...".
				64 // deposit amount encoded into 256 bits.

		event.AuthenticatedHash, err = NewUint256(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventChannelCreated) Digest() string {
	return EthDigestChannelCreated
}

//---------------------------------------------------------------------------------------------------------------------

type EventChannelToppedUp struct {
	Client          *Address // Indexed.
	Agent           *Address // Indexed.
	OpenBlockNumber *Uint256 // Indexed.
	OfferingHash    *Uint256
	AddedDeposit    *Uint192
}

func NewEventChannelToppedUp(topics [4]string, hexData string) (*EventChannelToppedUp, error) {
	var err error
	event := &EventChannelToppedUp{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Client address parsing.
		clientAddressHex := topics[1][2:] // Skipping "0x".
		offsetFrom := 24                  // Skipping first 24 bytes from 64.
		event.Client, err = NewAddress(clientAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Agent address parsing.
		agentAddressHex := topics[2][2:] // skipping "0x".
		offsetFrom := 24                 // skipping first 24 bytes from 64.
		event.Agent, err = NewAddress(agentAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Open block number parsing.
		openBlockNumberHex := topics[3][2:]
		event.OpenBlockNumber, err = NewUint256(openBlockNumberHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Offering hash parsing (data field).
		offsetFrom := 2             // "0x...".
		offsetTo := offsetFrom + 64 // we need 48 bytes.
		event.OfferingHash, err = NewUint256(hexData[offsetFrom:offsetTo])
		if err != nil {
			return nil, err
		}
	}

	{
		// Added deposit parsing (data field).
		offsetFrom :=
			2 + // "0x...".
				64 + // offering hash field.
				16 // only 192 bites are used (from the end).

		event.AddedDeposit, err = NewUint192(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventChannelToppedUp) Digest() string {
	return EthDigestChannelToppedUp
}

//---------------------------------------------------------------------------------------------------------------------

type EventChannelCloseRequested struct {
	Client          *Address // Indexed.
	Agent           *Address // Indexed.
	OpenBlockNumber *Uint256 // Indexed.
	OfferingHash    *Uint256
	Balance         *Uint192
}

func NewEventChannelCloseRequested(topics [4]string, hexData string) (*EventChannelCloseRequested, error) {
	return nil, errors.New("not implemented") // todo: fix this when API would be fixed
}

func (e *EventChannelCloseRequested) Digest() string {
	return EthChannelCloseRequested
}

//---------------------------------------------------------------------------------------------------------------------

type EventOfferingCreated struct {
	Agent         *Address // Indexed.
	OfferingHash  *Uint256 // Indexed.
	MinDeposit    *Uint192
	CurrentSupply *Uint192
}

func NewEventServiceOfferingCreated(topics [3]string, hexData string) (*EventOfferingCreated, error) {
	var err error
	event := &EventOfferingCreated{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Agent address parsing.
		agentAddressHex := topics[1][2:] // Skipping "0x".
		offsetFrom := 24                 // Skipping first 24 bytes from 64.
		event.Agent, err = NewAddress(agentAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Offering hash parsing.
		offeringHashHex := topics[2][2:] // Skipping "0x".
		event.OfferingHash, err = NewUint256(offeringHashHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Min deposit parsing (data field).
		offsetFrom :=
			2 + // "0x...".
				16 // Only 192 bites are used (from the end).

		offsetTo := offsetFrom + 48
		event.MinDeposit, err = NewUint192(hexData[offsetFrom:offsetTo])
		if err != nil {
			return nil, err
		}
	}

	{
		// Current supply (data field).
		offsetFrom :=
			2 + // "0x...".
				64 + // min deposit field.
				16 // only 192 bites are used (from the end).

		event.CurrentSupply, err = NewUint192(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventOfferingCreated) Digest() string {
	return EthOfferingCreated
}

//---------------------------------------------------------------------------------------------------------------------

type EventOfferingDeleted struct {
	OfferingHash *Uint256 // Indexed.
}

func NewEventServiceOfferingDeleted(topics [2]string) (*EventOfferingDeleted, error) {
	var err error
	event := &EventOfferingDeleted{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Offering hash parsing.
		offeringHashHex := topics[1][2:] // Skipping "0x".
		event.OfferingHash, err = NewUint256(offeringHashHex)
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventOfferingDeleted) Digest() string {
	return EthOfferingDeleted
}

//---------------------------------------------------------------------------------------------------------------------

type EventOfferingEndpoint struct {
	Client          *Address // Indexed.
	OfferingHash    *Uint256 // Indexed.
	OpenBlockNumber *Uint256 // Indexed.
	EndpointHash    *Uint256
}

func NewEventServiceOfferingEndpoint(topics [4]string, hexData string) (*EventOfferingEndpoint, error) {
	var err error
	event := &EventOfferingEndpoint{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Agent address parsing.
		clientAddressHex := topics[1][2:] // Skipping "0x".
		offsetFrom := 24                  // Skipping first 24 bytes from 64.
		event.Client, err = NewAddress(clientAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Offering hash parsing.
		offeringHashHex := topics[2][2:] // Skipping "0x".
		event.OfferingHash, err = NewUint256(offeringHashHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Open block number parsing.
		openBlockNumberHex := topics[3][2:] // Skipping "0x".
		event.OpenBlockNumber, err = NewUint256(openBlockNumberHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// endpoint hash (data field).
		offsetFrom := 2 // "0x...".
		event.EndpointHash, err = NewUint256(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventOfferingEndpoint) Digest() string {
	return EthServiceOfferingEndpoint
}

//---------------------------------------------------------------------------------------------------------------------

type EventOfferingSupplyChanged struct {
	OfferingHash  *Uint256 // Indexed.
	CurrentSupply *Uint192
}

func NewEventServiceOfferingSupplyChanged(topics [2]string, hexData string) (*EventOfferingSupplyChanged, error) {
	var err error
	event := &EventOfferingSupplyChanged{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Offering hash parsing.
		offeringHashHex := topics[1][2:] // Skipping "0x".
		event.OfferingHash, err = NewUint256(offeringHashHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Current supply (data field).
		offsetFrom :=
			2 + // "0x...".
				16 // Only 192 bites are used (from the end).

		event.CurrentSupply, err = NewUint192(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventOfferingSupplyChanged) Digest() string {
	return EthServiceOfferingSupplyChanged
}

//---------------------------------------------------------------------------------------------------------------------

type EventOfferingPoppedUp struct {
	OfferingHash *Uint256 // Indexed.
}

func NewEventServiceOfferingPoppedUp(topics [2]string) (*EventOfferingPoppedUp, error) {
	var err error
	event := &EventOfferingPoppedUp{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Offering hash parsing.
		offeringHashHex := topics[1][2:] // Skipping "0x".
		event.OfferingHash, err = NewUint256(offeringHashHex)
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventOfferingPoppedUp) Digest() string {
	return EthServiceOfferingPoppedUp
}

//---------------------------------------------------------------------------------------------------------------------

type EventCooperativeChannelClose struct {
	Client          *Address // Indexed.
	Agent           *Address // Indexed.
	OpenBlockNumber *Uint256 // Indexed.
	OfferingHash    *Uint256
	Balance         *Uint192
}

func NewEventCooperativeChannelClose(topics [4]string, hexData string) (*EventCooperativeChannelClose, error) {
	var err error
	event := &EventCooperativeChannelClose{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Client address parsing.
		clientAddressHex := topics[1][2:] // Skipping "0x".
		offsetFrom := 24                  // Skipping first 24 bytes from 64.
		event.Client, err = NewAddress(clientAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Agent address parsing.
		agentAddressHex := topics[2][2:] // Skipping "0x".
		offsetFrom := 24                 // Skipping first 24 bytes from 64.
		event.Agent, err = NewAddress(agentAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Open block number parsing.
		openBlockNumberHex := topics[3][2:] // Skipping "0x".
		event.OpenBlockNumber, err = NewUint256(openBlockNumberHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Offering hash parsing (data field).
		offsetFrom := 2  // "0x...".
		offsetTo := offsetFrom + 64
		event.OfferingHash, err = NewUint256(hexData[offsetFrom:offsetTo])
		if err != nil {
			return nil, err
		}
	}

	{
		// Balance parsing (data field).
		offsetFrom :=
			2 + // "0x...".
				64 + // Min deposit field.
				16 // Only 192 bites are used (from the end).

		event.Balance, err = NewUint192(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventCooperativeChannelClose) Digest() string {
	return EthCooperativeChannelClose
}

//---------------------------------------------------------------------------------------------------------------------

type EventUncooperativeChannelClose struct {
	Client          *Address // Indexed.
	Agent           *Address // Indexed.
	OpenBlockNumber *Uint256 // Indexed.
	OfferingHash    *Uint256
	Balance         *Uint192
}

func NewEventUnCooperativeChannelClose(topics [4]string, hexData string) (*EventUncooperativeChannelClose, error) {
	var err error
	event := &EventUncooperativeChannelClose{}

	err = validateTopics(topics[:])
	if err != nil {
		return nil, err
	}

	{
		// Event digest parsing.
		digestHex := topics[0][2:] // Skipping "0x".
		if digestHex != event.Digest() {
			return nil, errorUnexpectedEventType(digestHex, event.Digest())
		}
	}

	{
		// Client address parsing.
		clientAddressHex := topics[1][2:] // Skipping "0x".
		offsetFrom := 24                  // Skipping first 24 bytes from 64.
		event.Client, err = NewAddress(clientAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Agent address parsing.
		agentAddressHex := topics[2][2:] // Skipping "0x".
		offsetFrom := 24                 // Skipping first 24 bytes from 64.
		event.Agent, err = NewAddress(agentAddressHex[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	{
		// Open block number parsing.
		openBlockNumberHex := topics[3][2:] // Skipping "0x".
		event.OpenBlockNumber, err = NewUint256(openBlockNumberHex)
		if err != nil {
			return nil, err
		}
	}

	{
		// Offering hash parsing (data field).
		offsetFrom := 2  // "0x...".
		offsetTo := offsetFrom + 64
		event.OfferingHash, err = NewUint256(hexData[offsetFrom:offsetTo])
		if err != nil {
			return nil, err
		}
	}

	{
		// Balance parsing (data field).
		offsetFrom :=
			2 + // "0x...".
				64 + // Min deposit field.
				16 // Only 192 bites are used (from the end).

		event.Balance, err = NewUint192(hexData[offsetFrom:])
		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (e *EventUncooperativeChannelClose) Digest() string {
	return EthUncooperativeChannelClose
}

//---------------------------------------------------------------------------------------------------------------------

func errorUnexpectedEventType(receivedDigest, expectedDigest string) error {
	return fmt.Errorf("unexpected event type occurred: %s, but %s is expected", receivedDigest, expectedDigest)
}

func validateTopics(topics []string) error {
	for _, topic := range topics {
		if len(topic) != 66 { // "0x" + 64 bytes.
			return errors.New("Invalid topic occurred: " + topic)
		}

		if topic[:2] != "0x" {
			return errors.New("Invalid topic occurred: " + topic)
		}
	}

	return nil
}
