package tests

import (
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/testutil"
)

func TestGetAddressActivity(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetAddressActivity("demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetAllChainStatus(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetAllChainStatus()
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetAllChains(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetAllChains()
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetBlock(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetBlock(chains.EthMainnet, "latest")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetBlockHeights(t *testing.T) {
	results := make(chan services.BlockHeightsResult)
	go func() {
		// Call the function and pass parameters
		resultChan := testutil.Client.BaseService.GetBlockHeights(chains.EthMainnet, "2023-01-01", "2023-01-02")

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

func TestGetBlockHeightsByPage(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetBlockHeightsByPage(chains.EthMainnet, "2023-01-01", "2023-01-02")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetGasPrices(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetGasPrices(chains.EthMainnet, "erc20")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetLogEventsByAddress(t *testing.T) {
	results := make(chan services.LogEventResult)
	go func() {
		// Call the function and pass parameters
		startingBlock := 17679143
		endingBlock := "17679148"
		resultChan := testutil.Client.BaseService.GetLogEventsByAddress(chains.EthMainnet, "0xdac17f958d2ee523a2206206994597c13d831ec7", services.GetLogEventsByAddressQueryParamOpts{StartingBlock: &startingBlock, EndingBlock: &endingBlock})

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

func TestGetLogEventsByAddressByPage(t *testing.T) {
	// Test case 1
	startingBlock := 17679143
	endingBlock := "17679148"
	resp, err := testutil.Client.BaseService.GetLogEventsByAddressByPage(chains.EthMainnet, "0xdac17f958d2ee523a2206206994597c13d831ec7", services.GetLogEventsByAddressQueryParamOpts{StartingBlock: &startingBlock, EndingBlock: &endingBlock})
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetLogEventsByTopicHash(t *testing.T) {
	results := make(chan services.LogEventResult)
	go func() {
		// Call the function and pass parameters
		startingBlock := 17666774
		endingBlock := "17679143"
		resultChan := testutil.Client.BaseService.GetLogEventsByTopicHash(chains.EthMainnet, "0x27f12abfe35860a9a927b465bb3d4a9c23c8428174b83f278fe45ed7b4da2662", services.GetLogEventsByTopicHashQueryParamOpts{StartingBlock: &startingBlock, EndingBlock: &endingBlock})

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

func TestGetLogEventsByTopicHashByPage(t *testing.T) {
	// Test case 1
	startingBlock := 17666774
	endingBlock := "17679143"
	resp, err := testutil.Client.BaseService.GetLogEventsByTopicHashByPage(chains.EthMainnet, "0x27f12abfe35860a9a927b465bb3d4a9c23c8428174b83f278fe45ed7b4da2662", services.GetLogEventsByTopicHashQueryParamOpts{StartingBlock: &startingBlock, EndingBlock: &endingBlock})
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetLogs(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetLogs(chains.EthMainnet)
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetResolvedAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.BaseService.GetResolvedAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
		expectedAddress := "0xfc43f5f9dd45258b3aff31bdbe6561d97e8b71de"
		if *resp.Data.Items[0].Address != expectedAddress {
			t.Errorf("Error - Got this: %s, should be %s", *resp.Data.Items[0].Address, expectedAddress)
		}
	}
}
