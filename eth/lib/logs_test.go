package lib

import (
	"testing"
)

const (
	// Contract events digests.
	// Copied here to avoid cyclic reference with contract.

	EventDigestChannelCreated       = "a6153987181667023837aee39c3f1a702a16e5e146323ef10fb96844a526143c"
	EventDigestChannelToppedUp      = "392a992c1a7b756e553d8d97f43d59fafe79bc672808247debc077a6cdaba7b9"
	EthChannelCloseRequested        = "21ff66d79903f9d4ab6ab3c7c903af993e709be2ce2f4532d572925dea741cb1"
	EthOfferingCreated              = "49d573efb7cbb057727f6cadb4150ba6d5041c4fb55afe606508be636e158127"
	EthOfferingDeleted              = "21652905a07e2790c3a220d14394aee13681876bfbf38e658fa82ee5afe0c862"
	EthServiceOfferingEndpoint      = "00a7695de2bf4b4a523002334437d52e135b7a2a892d4471b5dd9005e5cd0681"
	EthServiceOfferingSupplyChanged = "1337b30376128e64c2ffd4e95d4c900b4ab42af11202b328722020216eeb46df"
	EthServiceOfferingPoppedUp      = "c8404827c21b5491a6c3dc0881307e47bfa40c3baf3d607c2d14f6bc808d4bfb"
	EthCooperativeChannelClose      = "56a4dfc7b9f93649d9142c7bef0a429decf8d3be895a3180c67a76a18d79f4ab"
	EthUncooperativeChannelClose    = "8a79bd24ee9bcfd977d6fc685befa8775c8a933f0abe82ab73b716cf419f968e"
)

func TestNormalLogsFetching(t *testing.T) {
	// todo: call truffle tests and form the blockchain
	// todo: parse info from the truffle output
	const kContractAddress = "0xd1eddf50e1678dd1ea2e946fc21c8b50478126d6"

	client := NewEthereumClient(host, port)

	checkEvent := func(eventDigest string) {
		response, err := client.GetLogs(
			kContractAddress,
			[]string{"0x" + eventDigest})

		if err != nil {
			t.Fatal("Can't call API: ", err, " Event digest: ", eventDigest)
		}

		if len(response.Result) == 0 {
			t.Fatal("Can't fetch result. Event digest: ", eventDigest)
		}
	}

	{
		// Test purpose:
		// Check if all events types are accessible via raw API call.
		checkEvent(EventDigestChannelCreated)

		// todo: enable back when truffle tests would be available
		//checkEvent(EthChannelCloseRequested)
		//checkEvent(EthOfferingCreated)
		//checkEvent(EthOfferingDeleted)
		//checkEvent(EthServiceOfferingEndpoint)
		//checkEvent(EthServiceOfferingSupplyChanged)
		//checkEvent(EthServiceOfferingPoppedUp)
		//checkEvent(EthCooperativeChannelClose)
		//checkEvent(EthUncooperativeChannelClose)
	}
}

func TestNegativeLogsFetching(t *testing.T) {
	// todo: call truffle tests and form the blockchain

	client := NewEthereumClient(host, port)

	{
		// Test purpose:
		// To check that no logs fetching is done, if no contract address is specified.

		_, err := client.GetLogs("", []string{"0x0"})
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}

	{
		// Test purpose:
		// To check that no logs fetching is done, if invalid topics are transferred.

		// todo: parse info from the truffle output
		const kContractAddress = "0xd1eddf50e1678dd1ea2e946fc21c8b50478126d6"

		_, err := client.GetLogs(kContractAddress, []string{"0x0"})
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}

	{
		// Test purpose:
		// To check that no logs fetching is done, if invalid topics are transferred.

		// todo: parse info from the truffle output
		const kContractAddress = "0xd1eddf50e1678dd1ea2e946fc21c8b50478126d6"

		_, err := client.GetLogs(kContractAddress, []string{"", ""})
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}
}

func TestLogsFetchingWithBrokenNetworkFetching(t *testing.T) {
	// todo: kill truffle node

	client := NewEthereumClient(host, port)

	{
		// Test purpose:
		// To check that error is emitted in case if network is not operable.

		// todo: parse info from the truffle output
		const kContractAddress = "0xd1eddf50e1678dd1ea2e946fc21c8b50478126d6"

		_, err := client.GetLogs(
			kContractAddress,
			[]string{"0x8a79bd24ee9bcfd977d6fc685befa8775c8a933f0abe82ab73b716cf419f968e"})
		if err == nil {
			t.Fatal("Error must be returned")
		}
	}
}
