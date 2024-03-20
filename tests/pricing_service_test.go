package tests

import (
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/testutil"
)

func TestGetTokenPrices(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.PricingService.GetTokenPrices(chains.EthMainnet, "USD", "0xb8c77482e45f1f44de1745f52c74426c631bdd52")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(*resp.Data) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}
