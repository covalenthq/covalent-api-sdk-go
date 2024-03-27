# Covalent SDK for Golang

The Covalent SDK is the fastest way to integrate the Covalent Unified API for working with blockchain data. The SDK works with all [supported chains](https://www.covalenthq.com/docs/networks/) including Mainnets and Testnets. 

**Note - Require `go1.22.1` and above for best results.**

> **Sign up for an API Key**
>
> To get your own Covalent API key, **[sign up here](https://www.covalenthq.com/platform/auth/register/)** and create your key from the *API Keys* tab.

## Getting started

```
go get github.com/covalenthq/covalent-api-sdk-go
```

## How to use the Covalent SDK

After installing, you can import and use the SDK with:

```go
package main

import (
    "fmt"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
    "github.com/covalenthq/covalent-api-sdk-go/chains"
)

func main() {
    var Client = covalentclient.CovalentClient("API_KEY")
    resp, err := Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		if len(resp.Data.Items) != 0 {
            // get first Balance
            fmt.Println(*resp.Data.Items[0].Balance)
		}
	}
}
```
> **Name Resolution**
>
> The Covalent SDK natively supports ENS domains (e.g. `demo.eth`), Lens Handles (e.g. `@demo.lens`) and Unstoppable Domains (e.g. `demo.x`) which automatically resolve to the underlying user address (e.g. `0xfC43f5F9dd45258b3AFf31Bdbe6561D97e8B71de`)

### How to apply supported query parameters to endpoints
Query parameters are handled using variables that are `references and pointers`. Developers can reference the `godocs` for autocomplete suggestions for the supported parameters. These supported parameters can be passed in any order. 

For example, the following sets the `QuoteCurrency` query parameter to `CAD` and the parameter `Nft` to `true` for fetching all the token balances, including NFTs, for a wallet address:
```go
package main

import (
    "fmt"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
    "github.com/covalenthq/covalent-api-sdk-go/services"
    "github.com/covalenthq/covalent-api-sdk-go/quotes"
    "github.com/covalenthq/covalent-api-sdk-go/chains"
)

func main() {
    nft := true
	var quoteValue quotes.Quote = quotes.CAD
    var Client = covalentclient.CovalentClient("API_KEY")
    resp, err := Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth", services.GetTokenBalancesForWalletAddressQueryParamOpts{Nft: &nft, QuoteCurrency: &quoteValue})
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		if len(resp.Data.Items) != 0 {
            // get first Balance
            fmt.Println(*resp.Data.Items[0].Balance)
		}
	}
}
```

## Supported Endpoints

The Covalent SDK provides comprehensive support for all Class A, Class B, and Pricing endpoints that are grouped under the following *Service* namespaces:

- [`SecurityService`](#securityservice): Access to the token approvals endpoints
- [`BalanceService`](#balanceservice): Access to the balances endpoints
- [`BaseService`](#baseservice): Access to the address activity, log events, chain status, and block retrieval endpoints
- [`NftService`](#nftservice): Access to the  NFT endpoints
- [`PricingService`](#pricingservice): Access to the historical token prices endpoint
- [`TransactionService`](#transactionservice): Access to the transactions endpoints
- [`XykService`](#xykservice): Access to the XY=K suite of endpoints

### SecurityService

The `SecurityService` class refers to the [token approvals API endpoints](https://www.covalenthq.com/docs/api/security/get-token-approvals-for-address/):

- `GetApprovals()`: Get a list of approvals across all ERC20 token contracts categorized by spenders for a wallet’s assets.
- `GetNftApprovals()`: Get a list of approvals across all NFT contracts categorized by spenders for a wallet’s assets.

### BalanceService

The `BalanceService` class refers to the [balances API endpoints](https://www.covalenthq.com/docs/api/balances/get-token-balances-for-address/):

- `GetTokenBalancesForWalletAddress()`: Fetch the native, fungible (ERC20), and non-fungible (ERC721 & ERC1155) tokens held by an address. Response includes spot prices and other metadata.
- `GetHistoricalTokenBalancesForWalletAddress()`: Fetch the historical native, fungible (ERC20), and non-fungible (ERC721 & ERC1155) tokens held by an address at a given block height or date. Response includes daily prices and other metadata.
- `GetHistoricalPortfolioForWalletAddress()`: Render a daily portfolio balance for an address broken down by the token. The timeframe is user-configurable, defaults to 30 days.
- `GetErc20TransfersForWalletAddress()`: Render the transfer-in and transfer-out of a token along with historical prices from an address. (Paginated)
- `GetErc20TransfersForWalletAddressByPage()`: Render the transfer-in and transfer-out of a token along with historical prices from an address. (NonPaginated)
- `GetTokenHoldersV2ForTokenAddress()`: Get a list of all the token holders for a specified ERC20 or ERC721 token. Returns historic token holders when block-height is set (defaults to latest). Useful for building pie charts of token holders. (Paginated)
- `GetTokenHoldersV2ForTokenAddressByPage()`: Get a list of all the token holders for a specified ERC20 or ERC721 token. Returns historic token holders when block-height is set (defaults to latest). Useful for building pie charts of token holders. (Nonpaginated)
- `GetNativeTokenBalance()`: Get the native token balance for an address. This endpoint is required because native tokens are usually not ERC20 tokens and sometimes you want something lightweight.

### BaseService

The `BaseService` class refers to the [address activity, log events, chain status and block retrieval API endpoints](https://www.covalenthq.com/docs/api/base/get-address-activity/):

- `GetBlock()`: Fetch and render a single block for a block explorer.
- `GetLogs()`: Get all the event logs of the latest block, or for a range of blocks. Includes sender contract metadata as well as decoded logs.
- `GetResolvedAddress()`: Used to resolve ENS, RNS and Unstoppable Domains addresses.
- `GetBlockHeights()`: Get all the block heights within a particular date range. Useful for rendering a display where you sort blocks by day. (Paginated)
- `GetBlockHeightsByPage()`: Get all the block heights within a particular date range. Useful for rendering a display where you sort blocks by day. (Nonpaginated)
- `GetLogEventsByAddress()`: Get all the event logs emitted from a particular contract address. Useful for building dashboards that examine on-chain interactions. (Paginated)
- `GetLogEventsByAddressByPage()`: Get all the event logs emitted from a particular contract address. Useful for building dashboards that examine on-chain interactions. (Nonpaginated)
- `GetLogEventsByTopicHash()`: Get all event logs of the same topic hash across all contracts within a particular chain. Useful for cross-sectional analysis of event logs that are emitted on-chain. (Paginated)
- `GetLogEventsByTopicHashByPage()`: Get all event logs of the same topic hash across all contracts within a particular chain. Useful for cross-sectional analysis of event logs that are emitted on-chain. (Nonpaginated)
- `GetAllChains()`: Used to build internal dashboards for all supported chains on Covalent.
- `GetAllChainStatus()`: Used to build internal status dashboards of all supported chains.
- `GetAddressActivity()`: Locate chains where an address is active on with a single API call.
- `GetGasPrices()`: Get real-time gas estimates for different transaction speeds on a specific network, enabling users to optimize transaction costs and confirmation times.

### NftService

The `NftService` class refers to the [NFT API endpoints](https://www.covalenthq.com/docs/api/nft/get-nfts-for-address/):

- `GetChainCollections()`: Used to fetch the list of NFT collections with downloaded and cached off chain data like token metadata and asset files. (Paginated)
- `GetChainCollectionsByPage()`: Used to fetch the list of NFT collections with downloaded and cached off chain data like token metadata and asset files. (Nonpaginated)
- `GetNftsForAddress()`: Used to render the NFTs (including ERC721 and ERC1155) held by an address.
- `GetTokenIdsForContractWithMetadata()`: Get NFT token IDs with metadata from a collection. Useful for building NFT card displays. (Paginated)
- `GetTokenIdsForContractWithMetadataByPage()`: Get NFT token IDs with metadata from a collection. Useful for building NFT card displays. (Nonpaginated)
- `GetNftMetadataForGivenTokenIDForContract()`: Get a single NFT metadata by token ID from a collection. Useful for building NFT card displays.
- `GetNftTransactionsForContractTokenId()`: Get all transactions of an NFT token. Useful for building a transaction history table or price chart.
- `GetTraitsForCollection()`: Used to fetch and render the traits of a collection as seen in rarity calculators.
- `GetAttributesForTraitInCollection()`: Used to get the count of unique values for traits within an NFT collection.
- `GetCollectionTraitsSummary()`: Used to calculate rarity scores for a collection based on its traits.
- `CheckOwnershipInNft()`: Used to verify ownership of NFTs (including ERC-721 and ERC-1155) within a collection.
- `CheckOwnershipInNftForSpecificTokenId()`: Used to verify ownership of a specific token (ERC-721 or ERC-1155) within a collection.
- `GetNftMarketSaleCount()`: Used to build a time-series chart of the sales count of an NFT collection.
- `GetNftMarketVolume()`: Used to build a time-series chart of the transaction volume of an NFT collection.
- `GetNftMarketFloorPrice()`: Used to render a price floor chart for an NFT collection.

### PricingService

The `PricingService` class refers to the [historical token prices API endpoint](https://www.covalenthq.com/docs/api/pricing/get-historical-token-prices/):

- `GetTokenPrices()`: Get historic prices of a token between date ranges. Supports native tokens.

### TransactionService

The `TransactionService` class refers to the [transactions API endpoints](https://www.covalenthq.com/docs/api/transactions/get-a-transaction/):

- `GetAllTransactionsForAddress()`: Fetch and render the most recent transactions involving an address. Frequently seen in wallet applications. (Paginated)
- `GetAllTransactionsForAddressByPage()`: Fetch and render the most recent transactions involving an address. Frequently seen in wallet applications. (Nonpaginated)
- `GetTransactionsForAddressV3()`: Fetch and render the most recent transactions involving an address. Frequently seen in wallet applications.
- `GetTransaction()`: Fetch and render a single transaction including its decoded log events. Additionally return semantically decoded information for DEX trades, lending and NFT sales.
- `GetTransactionsForBlock()`: Fetch all transactions including their decoded log events in a block and further flag interesting wallets or transactions.
- `GetTransactionSummary()`: Fetch the earliest and latest transactions, and the transaction count for a wallet. Calculate the age of the wallet and the time it has been idle and quickly gain insights into their engagement with web3.
- `GetTimeBucketTransactionsForAddress()`: Fetch all transactions including their decoded log events in a 15-minute time bucket interval.
- `GetTransactionsForBlockHashByPage()`: Fetch all transactions including their decoded log events in a block and further flag interesting wallets or transactions.
- `GetTransactionsForBlockHash()`: Fetch all transactions including their decoded log events in a block and further flag interesting wallets or transactions.


The functions `GetAllTransactionsForAddressByPage()`, `GetTransactionsForAddressV3()`, and `GetTimeBucketTransactionsForAddress()` have been enhanced with the introduction of `Next()` and `Prev()` support functions. These functions facilitate a smoother transition for developers navigating through our links object, which includes `Prev` and `Next` fields. Instead of requiring developers to manually extract values from these fields and create Golang API calls for the URL values, the new `Next()` and `Prev()` functions provide a streamlined approach, allowing developers to simulate this behavior more efficiently.

```go
package main

import (
	"fmt"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
)

func main() {
	debug := true

	var Client = covalentclient.CovalentClient("API_KEY", covalentclient.CovalentClientSettings{Debug: &debug})
	resp, err := Client.TransactionService.GetAllTransactionsForAddressByPage(chains.EthMainnet, "demo.eth")
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		if len(resp.Data.Items) != 0 {
			// get first Balance
			fmt.Println(*resp.Data.Items[0].BlockHash)
			prev, er := resp.Data.Prev()
			if er != nil {
				fmt.Printf("error: %s", err)
			} else {
				fmt.Println(prev.Data.CurrentPage)
			}
		}
	}
}
```

### XykService

The `XykService` refers to the [Xy=k API endpoints](https://www.covalenthq.com/docs/api/xyk/get-xyk-pools/):

- `GetPools()`: Get all the pools of a particular DEX. Supports most common DEXs (Uniswap, SushiSwap, etc), and returns detailed trading data (volume, liquidity, swap counts, fees, LP token prices).
- `GetPoolByAddress()`: Get the 7 day and 30 day time-series data (volume, liquidity, price) of a particular liquidity pool in a DEX. Useful for building time-series charts on DEX trading activity.
- `GetPoolsForTokenAddress()`: Get all pools and the supported DEX for a token. Useful for building a table of top pairs across all supported DEXes that the token is trading on.
- `GetPoolsForWalletAddress()`: Get all pools and supported DEX for a wallet. Useful for building a personal DEX UI showcasing pairs and supported DEXes associated to the wallet.
- `GetAddressExchangeBalances()`: Return balance of a wallet/contract address on a specific DEX.
- `GetNetworkExchangeTokens()`: Retrieve all network exchange tokens for a specific DEX. Useful for building a top tokens table by total liquidity within a particular DEX.
- `GetLpTokenView()`: Get a detailed view for a single liquidity pool token. Includes time series data.
- `GetSupportedDEXes()`: Get all the supported DEXs available for the xy=k endpoints, along with the swap fees and factory addresses.
- `GetDexForPoolAddress()`: Get the supported DEX given a pool address, along with the swap fees, DEX's logo url, and factory addresses. Useful to identifying the specific DEX to which a pair address is associated.
- `GetSingleNetworkExchangeToken()`: Get historical daily swap count for a single network exchange token.
- `GetTransactionsForAccountAddress()`: Get all the DEX transactions of a wallet. Useful for building tables of DEX activity segmented by wallet.
- `GetTransactionsForTokenAddress()`: Get all the transactions of a token within a particular DEX. Useful for getting a per-token view of DEX activity.
- `GetTransactionsForExchange()`: Get all the transactions of a particular DEX liquidity pool. Useful for building a transactions history table for an individual pool.
- `GetTransactionsForDex()`: Get all the the transactions for a given DEX. Useful for building DEX activity views.
- `GetEcosystemChartData()`: Get a 7d and 30d time-series chart of DEX activity. Includes volume and swap count.
- `GetHealthData()`: Ping the health of xy=k endpoints to get the synced block height per chain.

## Additional Helper Functions
### CalculatePrettyBalance
The `CalculatePrettyBalance` function is designed to take up to 4 inputs: the `Balance` field obtained from the `TokenBalances` endpoint and the `ContractDecimals`. The function also includes two optional fields, `roundOff` and `precision`, to allow developers to round the unscaled balance to a certain decimal precision. The primary purpose of this function is to convert the scaled token balance (the balance parameter) into its unscaled, human-readable form. The scaled balance needs to be divided by 10^(contractDecimals) to remove the scaling factor.

```go
package main

import (
    "fmt"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
    "github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/valuefmt"
)

func main() {
    var Client = covalentclient.CovalentClient("API_KEY")
    resp, err := Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		if len(resp.Data.Items) != 0 {
            // get first Balance
			prettybalance := valuefmt.CalculatePrettyBalance(*resp.Data.Items[0].Balance, *resp.Data.Items[0].ContractDecimals, true, 0)
			fmt.Println(prettybalance)
		}
	}
}
```
### PrettifyCurrency
The `PrettifyCurrency` function refines the presentation of a monetary value, accepting a numerical amount and a fiat currency code as parameters (with USD as the default currency). It simplifies currency formatting for developers, ensuring visually polished representations of financial information in user interfaces for an enhanced user experience.

```go
package main

import (
	"fmt"
	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
	"github.com/covalenthq/covalent-api-sdk-go/quotes"
	"github.com/covalenthq/covalent-api-sdk-go/valuefmt"
)

func main() {
    var Client = covalentclient.CovalentClient("API_KEY")
    resp, err := Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth")
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		if len(resp.Data.Items) != 0 {
            // get first Balance
			opts := valuefmt.NewCurrencyOptions();
			opts.Currency = quotes.CAD
			opts.Decimals = 5			
			prettyCurrency := valuefmt.PrettifyCurrency(*resp.Data.Items[0].QuoteRate, opts)
			fmt.Println(prettyCurrency)
		}
	}
}
```

## Built-in SDK Features
### Explaining Pagination Mechanism Within the SDK

The following endpoints support pagination:

- `GetErc20TransfersForWalletAddress()`
- `GetTokenHoldersV2ForTokenAddress()`
- `GetBlockHeights()`
- `GetLogEventsByAddress()`
- `GetLogEventsByTopicHash()`
- `GetChainCollections()`
- `GetTokenIdsForContractWithMetadata()`
- `GetAllTransactionsForAddress()`

Using the Covalent API, paginated supported endpoints return only 100 items, such as transactions or log events, per page. However, the Covalent SDK leverages go channels to *seamlessly fetch all items without the user having to deal with pagination*. 

For example, the following fetches ALL transactions for `demo.eth` on Ethereum:
```go
package main

import (
	"fmt"
	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
	"github.com/covalenthq/covalent-api-sdk-go/services"
)

func main() {
	results := make(chan services.TransactionResult)
	go func() {
		var Client = covalentclient.CovalentClient("API_KEY")
		// Call the function and pass parameters
		resultChan := Client.TransactionService.GetAllTransactionsForAddress(chains.EthMainnet, "demo.eth")

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
			fmt.Printf("error: %s", result.Err)
		} else {
			fmt.Println(*result.Transaction.BlockHeight)
		}
	}
}
```

### Debugger Mode

Developers have the option to enable a debugger mode that provides response times, the URLs of called endpoints, and the HTTP statuses of those endpoints. This feature helps users identify which endpoints may have encountered failures. The default is `debug = false` if no input is provided.

```go
package main

import (
	"fmt"
	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
	"github.com/covalenthq/covalent-api-sdk-go/quotes"
	"github.com/covalenthq/covalent-api-sdk-go/services"
)

func main() {
	nft := true
	var quoteValue quotes.Quote = quotes.CAD
	debug := true

	var Client = covalentclient.CovalentClient("API_KEY", covalentclient.CovalentClientSettings{Debug: &debug})
	resp, err := Client.BalanceService.GetTokenBalancesForWalletAddress(chains.EthMainnet, "demo.eth",
		services.GetTokenBalancesForWalletAddressQueryParamOpts{Nft: &nft, QuoteCurrency: &quoteValue})
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		if len(resp.Data.Items) != 0 {
			// get first Balance
			fmt.Println(*resp.Data.Items[0].ContractName)
		}
	}
}
```

![example result image](https://www.datocms-assets.com/86369/1697154621-sdk-debugger-output.png)

### Retry Mechanism

Each endpoint is equipped with an exponential backoff algorithm that exponentially extends the wait time between retries, up to a `maximum of 5` retry attempts.


### Error Handling
The paginated endpoints throw an error message in this format: `An error occurred {ErrorCode}: {ErrorMessage}`. The developer will need to `catch` these errors. Note - these endpoints do not follow the default response format which is:
```go
❴ 
    "Data": ...,
    "Error": false,
    "ErrorMessage": nil,
    "ErrorCode": nil
❵
```

### Error codes
Covalent uses standard HTTP response codes to indicate the success or failure of an API request. In general: codes in the 2xx range indicate success. Codes in the 4xx range indicate an error that failed given the information provided (e.g., a required parameter was omitted, etc.). Codes in the 5xx range indicate an error with Covalent's servers (these are rare).

| Code      | Description |
| ----------- | ----------- |
| 200      | OK	Everything worked as expected.       |
| 400   | Bad Request	The request could not be accepted, usually due to a missing required parameter.        |
| 401   | Unauthorized	No valid API key was provided.        |
| 404   | Not Found	The request path does not exist.        |
| 429   | Too Many Requests	You are being rate-limited. Please see the rate limiting section for more information.        |
| 500, 502, 503   | Server Errors	Something went wrong on Covalent's servers. These are rare.        |

## Tests
Before running tests, go to the `testutil/test_util.go` file and update your `API_KEY`. Do not commit this into your branch. Now you can run your tests in the root directory.
```
go test ./tests/
```
OR

you can run individual tests for each service, for example
```
go test ./tests/security_service_test.go
```

## Documentation

The Covalent API SDK documentation is integrated within the source code through `godoc` comments. When utilizing an Integrated Development Environment (IDE), the SDK provides generated types and accompanying documentation for seamless reference and usage.
