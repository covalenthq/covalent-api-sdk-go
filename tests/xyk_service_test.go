package tests

import (
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/testutil"
)

func TestGetAddressExchangeBalances(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetAddressExchangeBalances(chains.EthMainnet, "uniswap_v2", "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetDexForPoolAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetDexForPoolAddress(chains.EthMainnet, "0x4161fa43eaa1ac3882aeed12c5fc05249e533e67")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedDexName := "uniswap_v2"
		if *resp.Data.Items[0].DexName != expectedDexName {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].DexName, expectedDexName)
		}
	}
}

func TestGetEcosystemChartData(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetEcosystemChartData(chains.EthMainnet, "uniswap_v2")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetHealthData(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetHealthData(chains.EthMainnet, "uniswap_v2")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetLpTokenView(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetLpTokenView(chains.EthMainnet, "uniswap_v2", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedTicker := "WETH"
		expectedName := "Wrapped Ether"
		if *resp.Data.Items[0].ContractTickerSymbol != expectedTicker {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].ContractTickerSymbol, expectedTicker)
		}
		if *resp.Data.Items[0].ContractName != expectedName {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].ContractName, expectedName)
		}
	}
}

func TestGetNetworkExchangeTokens(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetNetworkExchangeTokens(chains.EthMainnet, "uniswap_v2")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedTicker := "WETH"
		expectedName := "Wrapped Ether"
		if *resp.Data.Items[0].ContractTickerSymbol != expectedTicker {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].ContractTickerSymbol, expectedTicker)
		}
		if *resp.Data.Items[0].ContractName != expectedName {
			t.Errorf("Error - Got this:  %s, should be %s", *resp.Data.Items[0].ContractName, expectedName)
		}
	}
}

func TestGetPoolByAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetPoolByAddress(chains.FantomMainnet, "spiritswap", "0xdbc490b47508d31c9ec44afb6e132ad01c61a02c")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedExchange := "0xdbc490b47508d31c9ec44afb6e132ad01c61a02c"
		if *resp.Data.Items[0].Exchange != expectedExchange {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].Exchange, expectedExchange)
		}
	}
}

func TestGetPools(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetPools(chains.EthMainnet, "uniswap_v2")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedDexName := "uniswap_v2"
		if *resp.Data.Items[0].DexName != expectedDexName {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].DexName, expectedDexName)
		}
	}
}

func TestGetPoolsForTokenAddress(t *testing.T) {
	// Test case 1
	requestedDexName := "uniswap_v2"
	resp, err := testutil.Client.XykService.GetPoolsForTokenAddress(chains.EthMainnet, "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", 0, services.GetPoolsForTokenAddressQueryParamOpts{DexName: &requestedDexName})
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedDexName := "uniswap_v2"
		if *resp.Data.Items[0].DexName != expectedDexName {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].DexName, expectedDexName)
		}
	}
}

func TestGetPoolsForWalletAddress(t *testing.T) {
	// Test case 1
	requestedDexName := "uniswap_v2"
	resp, err := testutil.Client.XykService.GetPoolsForWalletAddress(chains.EthMainnet, "ganeshswami.eth", 0, services.GetPoolsForWalletAddressQueryParamOpts{DexName: &requestedDexName})
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedDexName := "uniswap_v2"
		if *resp.Data.Items[0].DexName != expectedDexName {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].DexName, expectedDexName)
		}
	}
}

func TestGetSingleNetworkExchangeToken(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetSingleNetworkExchangeToken(chains.EthMainnet, "uniswap_v2", "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedDexName := "uniswap_v2"
		if *resp.Data.Items[0].DexName != expectedDexName {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].DexName, expectedDexName)
		}
	}
}

func TestGetSupportedDEXes(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetSupportedDEXes()
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForAccountAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetTransactionsForAccountAddress(chains.EthMainnet, "uniswap_v2", "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForDex(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetTransactionsForDex(chains.EthMainnet, "uniswap_v2")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForExchange(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetTransactionsForExchange(chains.FantomMainnet, "spiritswap", "0xdbc490b47508d31c9ec44afb6e132ad01c61a02c")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForTokenAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.XykService.GetTransactionsForTokenAddress(chains.EthMainnet, "uniswap_v2", "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}
