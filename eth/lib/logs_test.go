package lib

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"pxctrl/eth/contract"
	"testing"
	"bytes"
	"pxctrl/eth/lib/tests"
)

var (
	PrivateKey = ""
	PSCAddress = ""

	// Test sets of dummy data.
	// Used as placeholders for parameters in contract methods calls.
	testEthAddress1    = [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	testEthAddress2    = [20]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	testByte32SetZero  = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	testByte32SetFull  = [32]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	testUint256Zero, _ = NewUint256("0")
	testUint256Full, _ = NewUint256("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	testUint192Zero, _ = NewUint192("0")
)

// Internal methods
//---------------------------------------------------------------------------------------------------------------------

// fetchPSCAddress returns address of PSC in the currently active test chain.
// In case of successfully retrieved address  - caches retrieved address and returns it on the next calls,
// instead of doing redundant requests.
func fetchPSCAddress() string {
	if PSCAddress != "" {
		return PSCAddress
	}

	truffleAPI := tests.GethEthereumConfig().TruffleAPI
	response, err := http.Get(truffleAPI.Interface() + "/getPSC")
	if err != nil || response.StatusCode != 200 {
		log.Fatal("Can't fetch PSC address. It seems that test environment is broken.")
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Can't read response body. It seems that test environment is broken.")
	}

	data := make(map[string]interface{})
	json.Unmarshal(body, &data)

	PSCAddress = data["contract"].(map[string]interface{})["address"].(string)
	return PSCAddress
}

func fetchTestPrivateKey() string {
	if PrivateKey != "" {
		return PrivateKey
	}

	truffleAPI := tests.GethEthereumConfig().TruffleAPI
	response, err := http.Get(truffleAPI.Interface() + "/getKeys")
	if err != nil || response.StatusCode != 200 {
		log.Fatal("Can't fetch private key. It seems that test environment is broken.")
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Can't read response body. It seems that test environment is broken.")
	}

	data := make([]interface{}, 0, 0)
	json.Unmarshal(body, &data)

	PrivateKey = data[0].(map[string]interface{})["privateKey"].(string)
	return PrivateKey
}

func populateEvents() {
	geth := tests.GethEthereumConfig().Geth
	conn, err := ethclient.Dial(geth.Interface())
	if err != nil {
		log.Fatalf("Failed to connect to the EthereumConf client: %v", err)
	}

	contractAddress, err := NewAddress(fetchPSCAddress())
	if err != nil {
		log.Fatal("Failed to parse received contract address: ", err)
	}

	psc, err := contract.NewPrivatixServiceContract(contractAddress.Bytes(), conn)
	if err != nil {
		log.Fatal("Failed to connect to the EthereumConf client: ", err)
	}

	pKeyBytes, err := hex.DecodeString(fetchTestPrivateKey())
	if err != nil {
		log.Fatal("Failed to fetch test private key from the API: ", err)
	}

	key, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		log.Fatal("Failed parse received test private key: ", err)
	}

	auth := bind.NewKeyedTransactor(key)

	// Events populating
	//
	// WARN: note events arguments.
	// Them would be used for the further events deserialization tests.
	{
		_, err := psc.ThrowEventLogChannelCreated(
			auth, testEthAddress1, testEthAddress2, testByte32SetZero, big.NewInt(0), testByte32SetFull)

		if err != nil {
			log.Fatal("Failed to call ThrowEventLogChannelCreated: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogChannelToppedUp(
			auth, testEthAddress1, testEthAddress2, 0, testByte32SetFull, big.NewInt(0))

		if err != nil {
			log.Fatal("Failed to call ThrowEventLogChannelToppedUp: ", err)
		}
	}

	// todo: uncomment when API would be fixed
	//{
	//	_, err := psc.ThrowEventLogChannelCloseRequested(
	//		auth, testEthAddress1, testEthAddress2, testByte32SetZero, big.NewInt(0), testByte32SetFull)
	//
	//	if err != nil {
	//		log.Fatal("Failed to call ThrowEventLogChannelCloseRequested: ", err)
	//	}
	//}

	{
		_, err := psc.ThrowEventLogServiceOfferingCreated(
			auth, testEthAddress1, testByte32SetZero, big.NewInt(0), 0)

		if err != nil {
			log.Fatal("Failed to call ThrowEventLogServiceOfferingCreated: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogServiceOfferingDeleted(auth, testByte32SetZero)
		if err != nil {
			log.Fatal("Failed to call ThrowEventLogServiceOfferingDeleted: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogServiceOfferingEndpoint(
			auth, testEthAddress1, testByte32SetZero, 0, testByte32SetFull)

		if err != nil {
			log.Fatal("Failed to call ThrowEventLogServiceOfferingEndpoint: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogServiceOfferingSupplyChanged(auth, testByte32SetZero, 0)
		if err != nil {
			log.Fatal("Failed to call ThrowEventLogServiceOfferingSupplyChanged: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogServiceOfferingPopedUp(auth, testByte32SetZero)
		if err != nil {
			log.Fatal("Failed to call ThrowEventLogServiceOfferingPopedUp: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogCooperativeChannelClose(
			auth, testEthAddress1, testEthAddress2, 0, testByte32SetFull, big.NewInt(0))
		if err != nil {
			log.Fatal("Failed to call ThrowEventLogCooperativeChannelClose: ", err)
		}
	}

	{
		_, err := psc.ThrowEventLogUnCooperativeChannelClose(
			auth, testEthAddress1, testEthAddress2, 0, testByte32SetFull, big.NewInt(0))
		if err != nil {
			log.Fatal("Failed to call ThrowEventLogUnCooperativeChannelClose: ", err)
		}
	}
}

func TestNormalLogsFetching(t *testing.T) {
	populateEvents()

	node := tests.GethEthereumConfig().Geth
	client := NewEthereumClient(node.Host, node.Port)

	getEvent := func(eventDigest string) ([]string, string) {
		response, err := client.GetLogs(
			fetchPSCAddress(),
			[]string{"0x" + eventDigest}, "", "")

		if err != nil {
			t.Fatal("Can't call API: ", err, " Event digest: ", eventDigest)
		}

		if len(response.Result) == 0 {
			t.Fatal("Can't fetch result. Event digest: ", eventDigest)
		}

		return response.Result[0].Topics, response.Result[0].Data
	}

	{
		// Test purpose:
		// Check if all events types are accessible via raw API call.

		{
			topics, data := getEvent(EthDigestChannelCreated)
			event, err := NewEventChannelCreated(
				[4]string{topics[0], topics[1], topics[2], topics[3]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventChannelCreated: ", err)
			}

			// Comparing received parameters with originally sent values.
			clientAddress := event.Client.Bytes()
			if bytes.Compare(clientAddress[:], testEthAddress1[:]) != 0 {
				t.Fatal()
			}

			agentAddress := event.Agent.Bytes()
			if bytes.Compare(agentAddress[:], testEthAddress2[:]) != 0 {
				t.Fatal()
			}

			if event.OfferingHash.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.Deposit.String() != testUint192Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, data := getEvent(EthDigestChannelToppedUp)
			event, err := NewEventChannelToppedUp(
				[4]string{topics[0], topics[1], topics[2], topics[3]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventChannelToppedUp: ", err)
			}

			// Comparing received parameters with originally sent values.
			clientAddress := event.Client.Bytes()
			if bytes.Compare(clientAddress[:], testEthAddress1[:]) != 0 {
				t.Fatal()
			}

			agentAddress := event.Agent.Bytes()
			if bytes.Compare(agentAddress[:], testEthAddress2[:]) != 0 {
				t.Fatal()
			}

			if event.OpenBlockNumber.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.OfferingHash.String() != testUint256Full.String() {
				t.Fatal()
			}

			if event.AddedDeposit.String() != testUint192Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, data := getEvent(EthOfferingCreated)
			event, err := NewEventServiceOfferingCreated(
				[3]string{topics[0], topics[1], topics[2]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventChannelToppedUp: ", err)
			}

			agentAddress := event.Agent.Bytes()
			if bytes.Compare(agentAddress[:], testEthAddress1[:]) != 0 {
				t.Fatal()
			}

			if event.OfferingHash.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.MinDeposit.String() != testUint192Zero.String() {
				t.Fatal()
			}

			if event.CurrentSupply.String() != testUint192Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, _ := getEvent(EthOfferingDeleted)
			event, err := NewEventServiceOfferingDeleted([2]string{topics[0], topics[1]})

			if err != nil {
				t.Fatal("Can't create event NewEventServiceOfferingDeleted: ", err)
			}

			if event.OfferingHash.String() != testUint256Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, data := getEvent(EthServiceOfferingEndpoint)
			event, err := NewEventServiceOfferingEndpoint(
				[4]string{topics[0], topics[1], topics[2], topics[3]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventServiceOfferingEndpoint: ", err)
			}

			clientAddress := event.Client.Bytes()
			if bytes.Compare(clientAddress[:], testEthAddress1[:]) != 0 {
				t.Fatal()
			}

			if event.OfferingHash.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.OpenBlockNumber.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.EndpointHash.String() != testUint256Full.String() {
				t.Fatal()
			}
		}

		{
			topics, data := getEvent(EthServiceOfferingSupplyChanged)
			event, err := NewEventServiceOfferingSupplyChanged(
				[2]string{topics[0], topics[1]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventServiceOfferingSupplyChanged: ", err)
			}

			if event.OfferingHash.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.CurrentSupply.String() != testUint192Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, _ := getEvent(EthServiceOfferingPoppedUp)
			event, err := NewEventServiceOfferingPoppedUp(
				[2]string{topics[0], topics[1]})

			if err != nil {
				t.Fatal("Can't create event NewEventServiceOfferingPoppedUp: ", err)
			}

			if event.OfferingHash.String() != testUint256Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, data := getEvent(EthCooperativeChannelClose)
			event, err := NewEventCooperativeChannelClose(
				[4]string{topics[0], topics[1], topics[2], topics[3]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventCooperativeChannelClose: ", err)
			}

			// Comparing received parameters with originally sent values.
			clientAddress := event.Client.Bytes()
			if bytes.Compare(clientAddress[:], testEthAddress1[:]) != 0 {
				t.Fatal()
			}

			agentAddress := event.Agent.Bytes()
			if bytes.Compare(agentAddress[:], testEthAddress2[:]) != 0 {
				t.Fatal()
			}

			if event.OpenBlockNumber.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.OfferingHash.String() != testUint256Full.String() {
				t.Fatal()
			}

			if event.Balance.String() != testUint192Zero.String() {
				t.Fatal()
			}
		}

		{
			topics, data := getEvent(EthUncooperativeChannelClose)
			event, err := NewEventUnCooperativeChannelClose(
				[4]string{topics[0], topics[1], topics[2], topics[3]}, data)

			if err != nil {
				t.Fatal("Can't create event NewEventUnCooperativeChannelClose: ", err)
			}

			// Comparing received parameters with originally sent values.
			clientAddress := event.Client.Bytes()
			if bytes.Compare(clientAddress[:], testEthAddress1[:]) != 0 {
				t.Fatal()
			}

			agentAddress := event.Agent.Bytes()
			if bytes.Compare(agentAddress[:], testEthAddress2[:]) != 0 {
				t.Fatal()
			}

			if event.OpenBlockNumber.String() != testUint256Zero.String() {
				t.Fatal()
			}

			if event.OfferingHash.String() != testUint256Full.String() {
				t.Fatal()
			}

			if event.Balance.String() != testUint192Zero.String() {
				t.Fatal()
			}
		}
	}
}

func TestNegativeLogsFetching(t *testing.T) {
	node := tests.GethEthereumConfig().Geth
	client := NewEthereumClient(node.Host, node.Port)

	{
		// Test purpose:
		// To check that no logs fetching is done, if no contract address is specified.

		_, err := client.GetLogs("", []string{"0x0"}, "", "")
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}

	{
		// Test purpose:
		// To check that no logs fetching is done, if invalid topics are transferred.

		_, err := client.GetLogs(fetchPSCAddress(), []string{"0x0"}, "", "")
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}

	{
		// Test purpose:
		// To check that no logs fetching is done, if invalid topics are transferred.

		_, err := client.GetLogs(fetchPSCAddress(), []string{"", ""}, "", "")
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}
}

func TestLogsFetchingWithBrokenNetwork(t *testing.T) {
	node := tests.GethEthereumConfig().Geth
	client := NewEthereumClient(node.Host, node.Port+1) // note: invalid port is used

	{
		// Test purpose:
		// To check that error is emitted in case if network is not operable.

		_, err := client.GetLogs(
			fetchPSCAddress(),
			[]string{EthOfferingCreated}, "", "")
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}
}
