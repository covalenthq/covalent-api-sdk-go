package tests

import (
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/testutil"
)

func TestCheckOwnershipInNft(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.CheckOwnershipInNft(chains.EthMainnet, "0xA926597159c76314fBb8aAD077e394F2C77cfFBf", "0x39ee2c7b3cb80254225884ca001F57118C8f21B6")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestCheckOwnershipInNftForSpecificTokenId(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.CheckOwnershipInNftForSpecificTokenId(chains.EthMainnet, "0xA926597159c76314fBb8aAD077e394F2C77cfFBf", "0x39ee2c7b3cb80254225884ca001F57118C8f21B6", "9465")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetAttributesForTraitInCollection(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetAttributesForTraitInCollection(chains.EthMainnet, "0x39ee2c7b3cb80254225884ca001f57118c8f21b6", "Type")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetChainCollections(t *testing.T) {
	results := make(chan services.ChainCollectionItemResult)
	go func() {
		// Call the function and pass parameters
		resultChan := testutil.Client.NftService.GetChainCollections(chains.EthMainnet)

		// Receive values from the result channel
		for result := range resultChan {
			// Process each result as it becomes available
			results <- result
		}

		// Close the results channel when done
		close(results)
	}()
	// Now you can read values from the results channel as they arrive
	for result := range results {
		// Process each result
		if result.Err != nil {
			t.Errorf("error: %s", result.Err)
		}
	}
}

func TestGetChainCollectionsByPage(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetChainCollectionsByPage(chains.EthMainnet)
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetCollectionTraitsSummary(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetCollectionTraitsSummary(chains.EthMainnet, "0x39ee2c7b3cb80254225884ca001f57118c8f21b6")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetNftMarketFloorPrice(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetNftMarketFloorPrice(chains.EthMainnet, "0x3e511FE60D5FE09503C5F2a6477a75d0b905b335")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetNftMarketSaleCount(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetNftMarketSaleCount(chains.EthMainnet, "0x3e511FE60D5FE09503C5F2a6477a75d0b905b335")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetNftMarketVolume(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetNftMarketVolume(chains.EthMainnet, "0x3e511FE60D5FE09503C5F2a6477a75d0b905b335")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetNftMetadataForGivenTokenIdForContract(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetNftMetadataForGivenTokenIdForContract(chains.EthMainnet, "0x39ee2c7b3cb80254225884ca001f57118c8f21b6", "7142")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedAddress := "0x39ee2c7b3cb80254225884ca001f57118c8f21b6"
		if *resp.Data.Items[0].ContractAddress != expectedAddress {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].ContractAddress, expectedAddress)
		}
	}
}

func TestGetNftTransactionsForContractTokenId(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetNftTransactionsForContractTokenId(chains.EthMainnet, "0x39ee2c7b3cb80254225884ca001f57118c8f21b6", "7142")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetNftsForAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetNftsForAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTokenIdsForContractWithMetadata(t *testing.T) {
	results := make(chan services.NftTokenContractResult)
	go func() {
		// Call the function and pass parameters
		resultChan := testutil.Client.NftService.GetTokenIdsForContractWithMetadata(chains.EthMainnet, "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")

		// Receive values from the result channel
		for result := range resultChan {
			// Process each result as it becomes available
			results <- result
		}

		// Close the results channel when done
		close(results)
	}()
	// Now you can read values from the results channel as they arrive
	for result := range results {
		// Process each result
		if result.Err != nil {
			t.Errorf("WHAT HAPPENED WITH THIS ERROR error: %s", result.Err)
		}
	}
}

func TestGetTokenIdsForContractWithMetadataByPage(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetTokenIdsForContractWithMetadataByPage(chains.EthMainnet, "0x39ee2c7b3cb80254225884ca001f57118c8f21b6")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTraitsForCollection(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.NftService.GetTraitsForCollection(chains.EthMainnet, "0x39ee2c7b3cb80254225884ca001f57118c8f21b6")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}
