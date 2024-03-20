package tests

import (
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/testutil"
)

func TestGetRecentTransactionsForAddress(t *testing.T) {
	results := make(chan services.TransactionResult)
	go func() {
		// Call the function and pass parameters
		resultChan := testutil.Client.TransactionService.GetAllTransactionsForAddress(chains.EthMainnet, "demo.eth")

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

func TestGetRecentTransactionsForAddressByPage(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetAllTransactionsForAddressByPage(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTimeBucketTransactionsForAddress(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTimeBucketTransactionsForAddress(chains.EthMainnet, "demo.eth", 1799272)
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransaction(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTransaction(chains.EthMainnet, "0xb27a3a3d660b7d679ebbd7065635c8c3613e32eb0ebae24863a6375d73d1a128")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionSummary(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTransactionSummary(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForAddressV3(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTransactionsForAddressV3(chains.EthMainnet, "demo.eth", 0)
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForBlock(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTransactionsForBlock(chains.EthMainnet, "17685920")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForBlockHash(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTransactionsForBlockHash(chains.EthMainnet, "0x4ee50495ce7fbc4bfe412c38052eb8ca1bc470c0c07d756757f2fced9ad9d60b")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestGetTransactionsForBlockHashByPage(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.TransactionService.GetTransactionsForBlockHashByPage(chains.EthMainnet, "0x4ee50495ce7fbc4bfe412c38052eb8ca1bc470c0c07d756757f2fced9ad9d60b", 0)
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}
