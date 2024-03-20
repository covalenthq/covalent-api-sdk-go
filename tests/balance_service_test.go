package tests

import (
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/testutil"
)

func TestGetTokenBalancesForWalletAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTokenBalancesForWalletAddressNFT(t *testing.T) {
	// Test case 1
	nft := true
	noFetchNft := true
	resp, err := testutil.Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth", services.GetTokenBalancesForWalletAddressQueryParamOpts{Nft: &nft, NoNftFetch: &noFetchNft})
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetHistoricalPortfolioForWalletAddress(t *testing.T) {
	resp, err := testutil.Client.BalanceService.GetHistoricalPortfolioForWalletAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTokenHoldersV2ForTokenAddress(t *testing.T) {
	results := make(chan services.TokenHolderResult)
	go func() {
		// Call the function and pass parameters
		resultChan := testutil.Client.BalanceService.GetTokenHoldersV2ForTokenAddress(chains.EthMainnet, "0x987d7cc04652710b74fff380403f5c02f82e290a")

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

func TestGetErc20TransfersForWalletAddress(t *testing.T) {
	results := make(chan services.BlockTransactionWithContractTransfersResult)
	go func() {
		// Call the function and pass parameters
		contractAddress := "0xdac17f958d2ee523a2206206994597c13d831ec7"
		resultChan := testutil.Client.BalanceService.GetErc20TransfersForWalletAddress(chains.EthMainnet, "demo.eth", services.GetErc20TransfersForWalletAddressQueryParamOpts{ContractAddress: &contractAddress})

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

func TestGetHistoricalTokenBalancesForWalletAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BalanceService.GetHistoricalTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetHistoricalTokenBalancesForWalletAddressNFT(t *testing.T) {
	// Test case 1
	nft := true
	noFetchNft := true
	resp, err := testutil.Client.BalanceService.GetHistoricalTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth", services.GetHistoricalTokenBalancesForWalletAddressQueryParamOpts{Nft: &nft, NoNftFetch: &noFetchNft})
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetNativeTokenBalance(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BalanceService.GetNativeTokenBalance(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTokenHoldersV2ForTokenAddressByPage(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BalanceService.GetTokenHoldersV2ForTokenAddressByPage(chains.EthMainnet, "0x987d7cc04652710b74fff380403f5c02f82e290a")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}
