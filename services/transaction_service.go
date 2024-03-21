package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/genericmodels"
	"github.com/covalenthq/covalent-api-sdk-go/quotes"
	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

var clientKey string
var debugOutput bool
var workerCount int
var isKeyValid bool

type TransactionResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []Transaction `json:"items"`
}
type Transaction struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The height of the block.
	BlockHeight *int `json:"block_height,omitempty"`
	// The hash of the block. Use it to remove transactions from re-org-ed blocks.
	BlockHash *string `json:"block_hash,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	// The offset is the position of the tx in the block.
	TxOffset *int `json:"tx_offset,omitempty"`
	// Indicates whether a transaction failed or succeeded.
	Successful *bool `json:"successful,omitempty"`
	// The sender's wallet address.
	FromAddress *string `json:"from_address,omitempty"`
	// The address of the miner.
	MinerAddress *string `json:"miner_address,omitempty"`
	// The label of `from` address.
	FromAddressLabel *string `json:"from_address_label,omitempty"`
	// The receiver's wallet address.
	ToAddress *string `json:"to_address,omitempty"`
	// The label of `to` address.
	ToAddressLabel *string `json:"to_address_label,omitempty"`
	// The value attached to this tx.
	Value *utils.BigInt `json:"value,omitempty"`
	// The value attached in `quote-currency` to this tx.
	ValueQuote *float64 `json:"value_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyValueQuote *string `json:"pretty_value_quote,omitempty"`
	// The requested chain native gas token metadata.
	GasMetadata *genericmodels.ContractMetadata `json:"gas_metadata,omitempty"`
	GasOffered  *int64                          `json:"gas_offered,omitempty"`
	// The gas spent for this tx.
	GasSpent *int64 `json:"gas_spent,omitempty"`
	// The gas price at the time of this tx.
	GasPrice *int64 `json:"gas_price,omitempty"`
	// The total transaction fees (`gas_price` * `gas_spent`) paid for this tx, denoted in wei.
	FeesPaid *utils.BigInt `json:"fees_paid,omitempty"`
	// The gas spent in `quote-currency` denomination.
	GasQuote *float64 `json:"gas_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyGasQuote *string `json:"pretty_gas_quote,omitempty"`
	// The native gas exchange rate for the requested `quote-currency`.
	GasQuoteRate *float64 `json:"gas_quote_rate,omitempty"`
	// The explorer links for this transaction.
	Explorers *[]genericmodels.Explorer `json:"explorers,omitempty"`
	// The details for the dex transaction.
	DexDetails *[]DexReport `json:"dex_details,omitempty"`
	// The details for the NFT sale transaction.
	NftSaleDetails *[]NftSalesReport `json:"nft_sale_details,omitempty"`
	// The details for the lending protocol transaction.
	LendingDetails *[]LendingReport `json:"lending_details,omitempty"`
	// The log events.
	LogEvents *[]genericmodels.LogEvent `json:"log_events,omitempty"`
	// The details related to the safe transaction.
	SafeDetails *[]SafeDetails `json:"safe_details,omitempty"`
}
type DexReport struct {
	// The offset is the position of the log entry within an event log.
	LogOffset *int64 `json:"log_offset,omitempty"`
	// Stores the name of the protocol that facilitated the event.
	ProtocolName *string `json:"protocol_name,omitempty"`
	// Stores the contract address of the protocol that facilitated the event.
	ProtocolAddress *string `json:"protocol_address,omitempty"`
	// The protocol logo URL.
	ProtocolLogoUrl *string `json:"protocol_logo_url,omitempty"`
	// Stores the aggregator responsible for the event.
	AggregatorName *string `json:"aggregator_name,omitempty"`
	// Stores the contract address of the aggregator responsible for the event.
	AggregatorAddress *string `json:"aggregator_address,omitempty"`
	// DEXs often have multiple version - e.g Uniswap V1, V2 and V3. The `version` field allows you to look at a specific version of the DEX.
	Version *int64 `json:"version,omitempty"`
	// Similarly to the `version` field, `fork_version` gives you the version of the forked DEX. For example, most DEXs are a fork of Uniswap V2; therefore, `fork` = `aave` & `fork_version` = `2`
	ForkVersion *int64 `json:"fork_version,omitempty"`
	// Many DEXs are a fork of an already established DEX. The fork field allows you to see which DEX has been forked.
	Fork *string `json:"fork,omitempty"`
	// Stores the event taking place - e.g `swap`, `add_liquidity` and `remove_liquidity`.
	Event *string `json:"event,omitempty"`
	// Stores the address of the pair that the user interacts with.
	PairAddress        *string  `json:"pair_address,omitempty"`
	PairLpFeeBps       *float64 `json:"pair_lp_fee_bps,omitempty"`
	LpTokenAddress     *string  `json:"lp_token_address,omitempty"`
	LpTokenTicker      *string  `json:"lp_token_ticker,omitempty"`
	LpTokenNumDecimals *int64   `json:"lp_token_num_decimals,omitempty"`
	LpTokenName        *string  `json:"lp_token_name,omitempty"`
	LpTokenValue       *string  `json:"lp_token_value,omitempty"`
	ExchangeRateUsd    *float64 `json:"exchange_rate_usd,omitempty"`
	// Stores the address of token 0 in the specific pair.
	Token0Address *string `json:"token_0_address,omitempty"`
	// Stores the ticker symbol of token 0 in the specific pair.
	Token0Ticker *string `json:"token_0_ticker,omitempty"`
	// Stores the number of contract decimals of token 0 in the specific pair.
	Token0NumDecimals *int64 `json:"token_0_num_decimals,omitempty"`
	// Stores the contract name of token 0 in the specific pair.
	Token0Name *string `json:"token_0_name,omitempty"`
	// Stores the address of token 1 in the specific pair.
	Token1Address *string `json:"token_1_address,omitempty"`
	// Stores the ticker symbol of token 1 in the specific pair.
	Token1Ticker *string `json:"token_1_ticker,omitempty"`
	// Stores the number of contract decimals of token 1 in the specific pair.
	Token1NumDecimals *int64 `json:"token_1_num_decimals,omitempty"`
	// Stores the contract name of token 1 in the specific pair.
	Token1Name *string `json:"token_1_name,omitempty"`
	// Stores the amount of token 0 used in the transaction. For example, 1 ETH, 100 USDC, 30 UNI, etc.
	Token0Amount         *string  `json:"token_0_amount,omitempty"`
	Token0QuoteRate      *float64 `json:"token_0_quote_rate,omitempty"`
	Token0UsdQuote       *float64 `json:"token_0_usd_quote,omitempty"`
	PrettyToken0UsdQuote *string  `json:"pretty_token_0_usd_quote,omitempty"`
	Token0LogoUrl        *string  `json:"token_0_logo_url,omitempty"`
	// Stores the amount of token 1 used in the transaction. For example, 1 ETH, 100 USDC, 30 UNI, etc.
	Token1Amount         *string  `json:"token_1_amount,omitempty"`
	Token1QuoteRate      *float64 `json:"token_1_quote_rate,omitempty"`
	Token1UsdQuote       *float64 `json:"token_1_usd_quote,omitempty"`
	PrettyToken1UsdQuote *string  `json:"pretty_token_1_usd_quote,omitempty"`
	Token1LogoUrl        *string  `json:"token_1_logo_url,omitempty"`
	// Stores the wallet address that initiated the transaction (i.e the wallet paying the gas fee).
	Sender *string `json:"sender,omitempty"`
	// Stores the recipient of the transaction - recipients can be other wallets or smart contracts. For example, if you want to Swap tokens on Uniswap, the Uniswap router would typically be the recipient of the transaction.
	Recipient *string `json:"recipient,omitempty"`
}
type NftSalesReport struct {
	// The offset is the position of the log entry within an event log.
	LogOffset *int64 `json:"log_offset,omitempty"`
	// Stores the topic event hash. All events have a unique topic event hash.
	Topic0 *string `json:"topic0,omitempty"`
	// Stores the contract address of the protocol that facilitated the event.
	ProtocolContractAddress *string `json:"protocol_contract_address,omitempty"`
	// Stores the name of the protocol that facilitated the event.
	ProtocolName *string `json:"protocol_name,omitempty"`
	// The protocol logo URL.
	ProtocolLogoUrl *string `json:"protocol_logo_url,omitempty"`
	// Stores the address of the transaction recipient.
	To *string `json:"to,omitempty"`
	// Stores the address of the transaction sender.
	From *string `json:"from,omitempty"`
	// Stores the address selling the NFT.
	Maker *string `json:"maker,omitempty"`
	// Stores the address buying the NFT.
	Taker *string `json:"taker,omitempty"`
	// Stores the NFTs token ID. All NFTs have a token ID. Within a collection, these token IDs are unique. If the NFT is transferred to another owner, the token id remains the same, as this number is its identifier within a collection. For example, if a collection has 10K NFTs then an NFT in that collection can have a token ID from 1-10K.
	TokenId *string `json:"token_id,omitempty"`
	// Stores the address of the collection. For example, [Bored Ape Yacht Club](https://etherscan.io/token/0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d)
	CollectionAddress *string `json:"collection_address,omitempty"`
	// Stores the name of the collection.
	CollectionName *string `json:"collection_name,omitempty"`
	// Stores the address of the token used to purchase the NFT.
	TokenAddress *string `json:"token_address,omitempty"`
	// Stores the name of the token used to purchase the NFT.
	TokenName *string `json:"token_name,omitempty"`
	// Stores the ticker symbol of the token used to purchase the NFT.
	TickerSymbol *string `json:"ticker_symbol,omitempty"`
	// Stores the number decimal of the token used to purchase the NFT.
	NumDecimals       *int64   `json:"num_decimals,omitempty"`
	ContractQuoteRate *float64 `json:"contract_quote_rate,omitempty"`
	// The token amount used to purchase the NFT. For example, if the user purchased an NFT for 1 ETH. The `nft_token_price` field will hold `1`.
	NftTokenPrice *float64 `json:"nft_token_price,omitempty"`
	// The USD amount used to purchase the NFT.
	NftTokenPriceUsd       *float64 `json:"nft_token_price_usd,omitempty"`
	PrettyNftTokenPriceUsd *string  `json:"pretty_nft_token_price_usd,omitempty"`
	// The price of the NFT denominated in the chains native token. Even if a seller sells their NFT for DAI or MANA, this field denominates the price in the native token (e.g. ETH, AVAX, FTM, etc.)
	NftTokenPriceNative       *float64 `json:"nft_token_price_native,omitempty"`
	PrettyNftTokenPriceNative *string  `json:"pretty_nft_token_price_native,omitempty"`
	// Stores the number of NFTs involved in the sale. It's quick routine to see multiple NFTs involved in a single sale.
	TokenCount                *int64  `json:"token_count,omitempty"`
	NumTokenIdsSoldPerSale    *int64  `json:"num_token_ids_sold_per_sale,omitempty"`
	NumTokenIdsSoldPerTx      *int64  `json:"num_token_ids_sold_per_tx,omitempty"`
	NumCollectionsSoldPerSale *int64  `json:"num_collections_sold_per_sale,omitempty"`
	NumCollectionsSoldPerTx   *int64  `json:"num_collections_sold_per_tx,omitempty"`
	TradeType                 *string `json:"trade_type,omitempty"`
	TradeGroupType            *string `json:"trade_group_type,omitempty"`
}
type LendingReport struct {
	// The offset is the position of the log entry within an event log.
	LogOffset *int64 `json:"log_offset,omitempty"`
	// Stores the name of the lending protocol that facilitated the event.
	ProtocolName *string `json:"protocol_name,omitempty"`
	// Stores the contract address of the lending protocol that facilitated the event.
	ProtocolAddress *string `json:"protocol_address,omitempty"`
	// The protocol logo URL.
	ProtocolLogoUrl *string `json:"protocol_logo_url,omitempty"`
	// Lending protocols often have multiple version (e.g. Aave V1, V2 and V3). The `version` field allows you to look at a specific version of the Lending protocol.
	Version *string `json:"version,omitempty"`
	// Many lending protocols are a fork of an already established protocol. The fork column allows you to see which lending protocol has been forked.
	Fork *string `json:"fork,omitempty"`
	// Similarly to the `version` column, `fork_version` gives you the version of the forked lending protocol. For example, most lending protocols in the space are a fork of Aave V2; therefore, `fork` = `aave` & `fork_version` = `2`
	ForkVersion *string `json:"fork_version,omitempty"`
	// Stores the event taking place - e.g `borrow`, `deposit`, `liquidation`, 'repay' and 'withdraw'.
	Event *string `json:"event,omitempty"`
	// Stores the name of the LP token issued by the lending protocol. LP tokens can be debt or interest bearing tokens.
	LpTokenName *string `json:"lp_token_name,omitempty"`
	// Stores the number decimal of the LP token.
	LpDecimals *int `json:"lp_decimals,omitempty"`
	// Stores the ticker symbol of the LP token.
	LpTickerSymbol *string `json:"lp_ticker_symbol,omitempty"`
	// Stores the token address of the LP token.
	LpTokenAddress *string `json:"lp_token_address,omitempty"`
	// Stores the amount of LP token used in the event (e.g. 1 aETH, 100 cUSDC, etc).
	LpTokenAmount *float64 `json:"lp_token_amount,omitempty"`
	// Stores the total USD amount of all the LP Token used in the event.
	LpTokenPrice *float64 `json:"lp_token_price,omitempty"`
	// Stores the exchange rate between the LP and underlying token.
	ExchangeRate *float64 `json:"exchange_rate,omitempty"`
	// Stores the USD price of the LP Token used in the event.
	ExchangeRateUsd *float64 `json:"exchange_rate_usd,omitempty"`
	// Stores the name of the token going into the lending protocol (e.g the token getting deposited).
	TokenNameIn *string `json:"token_name_in,omitempty"`
	// Stores the number decimal of the token going into the lending protocol.
	TokenDecimalIn *int `json:"token_decimal_in,omitempty"`
	// Stores the address of the token going into the lending protocol.
	TokenAddressIn *string `json:"token_address_in,omitempty"`
	// Stores the ticker symbol of the token going into the lending protocol.
	TokenTickerIn *string `json:"token_ticker_in,omitempty"`
	// Stores the logo URL of the token going into the lending protocol.
	TokenLogoIn *string `json:"token_logo_in,omitempty"`
	// Stores the amount of tokens going into the lending protocol (e.g 1 ETH, 100 USDC, etc).
	TokenAmountIn *float64 `json:"token_amount_in,omitempty"`
	// Stores the total USD amount of all tokens going into the lending protocol.
	AmountInUsd       *float64 `json:"amount_in_usd,omitempty"`
	PrettyAmountInUsd *string  `json:"pretty_amount_in_usd,omitempty"`
	// Stores the name of the token going out of the lending protocol (e.g the token getting deposited).
	TokenNameOut *string `json:"token_name_out,omitempty"`
	// Stores the number decimal of the token going out of the lending protocol.
	TokenDecimalsOut *int `json:"token_decimals_out,omitempty"`
	// Stores the address of the token going out of the lending protocol.
	TokenAddressOut *string `json:"token_address_out,omitempty"`
	// Stores the ticker symbol of the token going out of the lending protocol.
	TokenTickerOut *string `json:"token_ticker_out,omitempty"`
	// Stores the logo URL of the token going out of the lending protocol.
	TokenLogoOut *string `json:"token_logo_out,omitempty"`
	// Stores the amount of tokens going out of the lending protocol (e.g 1 ETH, 100 USDC, etc).
	TokenAmountOut *float64 `json:"token_amount_out,omitempty"`
	// Stores the total USD amount of all tokens going out of the lending protocol.
	AmountOutUsd       *float64 `json:"amount_out_usd,omitempty"`
	PrettyAmountOutUsd *string  `json:"pretty_amount_out_usd,omitempty"`
	// Stores the type of loan the user is taking out. Lending protocols enable you to take out a stable or variable loan. Only relevant to borrow events.
	BorrowRateMode *float64 `json:"borrow_rate_mode,omitempty"`
	// Stores the interest rate of the loan. Only relevant to borrow events.
	BorrowRate *float64 `json:"borrow_rate,omitempty"`
	OnBehalfOf *string  `json:"on_behalf_of,omitempty"`
	// Stores the wallet address liquidating the loan. Only relevant to liquidation events.
	Liquidator *string `json:"liquidator,omitempty"`
	// Stores the wallet address of the user initiating the event.
	User *string `json:"user,omitempty"`
}
type SafeDetails struct {
	// The address that signed the safe transaction.
	OwnerAddress *string `json:"owner_address,omitempty"`
	// The signature of the owner for the safe transaction.
	Signature *string `json:"signature,omitempty"`
	// The type of safe signature used.
	SignatureType *string `json:"signature_type,omitempty"`
}
type TransactionsResponse struct {
	// The requested address.
	Address string `json:"address"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The current page of the response.
	CurrentPage int             `json:"current_page"`
	Links       PaginationLinks `json:"links"`
	// List of response items.
	Items []Transaction `json:"items"`
}
type PaginationLinks struct {
	// URL link to the next page.
	Prev *string `json:"prev,omitempty"`
	// URL link to the previous page.
	Next *string `json:"next,omitempty"`
}
type RecentTransactionsResponse struct {
	// The requested address.
	Address string `json:"address"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The current page of the response.
	CurrentPage int             `json:"current_page"`
	Links       PaginationLinks `json:"links"`
	// List of response items.
	Items []Transaction `json:"items"`
}
type TransactionsTimeBucketResponse struct {
	// The requested address.
	Address string `json:"address"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	Complete  bool   `json:"complete"`
	// The current bucket of the response.
	CurrentBucket int             `json:"current_bucket"`
	Links         PaginationLinks `json:"links"`
	// List of response items.
	Items []Transaction `json:"items"`
}
type TransactionsBlockPageResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string          `json:"chain_name"`
	Links     PaginationLinks `json:"links"`
	// List of response items.
	Items []Transaction `json:"items"`
}
type TransactionsBlockResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []Transaction `json:"items"`
}
type TransactionsSummaryResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []TransactionsSummary `json:"items"`
}
type TransactionsSummary struct {
	// The total number of transactions.
	TotalCount *int64 `json:"total_count,omitempty"`
	// The earliest transaction detected.
	EarliestTransaction *TransactionSummary `json:"earliest_transaction,omitempty"`
	// The latest transaction detected.
	LatestTransaction *TransactionSummary `json:"latest_transaction,omitempty"`
	// The gas summary for the transactions.
	GasSummary *GasSummary `json:"gas_summary,omitempty"`
}
type TransactionSummary struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	// The link to the transaction details using the Covalent API.
	TxDetailLink *string `json:"tx_detail_link,omitempty"`
}
type GasSummary struct {
	// The total number of transactions sent by the address.
	TotalSentCount *int64 `json:"total_sent_count,omitempty"`
	// The total transaction fees paid by the address, denoted in wei.
	TotalFeesPaid *utils.BigInt `json:"total_fees_paid,omitempty"`
	// The total transaction fees paid by the address, denoted in `quote-currency`.
	TotalGasQuote *float64 `json:"total_gas_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyTotalGasQuote *string `json:"pretty_total_gas_quote,omitempty"`
	// The average gas quote per transaction.
	AverageGasQuotePerTx *float64 `json:"average_gas_quote_per_tx,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyAverageGasQuotePerTx *string `json:"pretty_average_gas_quote_per_tx,omitempty"`
	// The requested chain native gas token metadata.
	GasMetadata *genericmodels.ContractMetadata `json:"gas_metadata,omitempty"`
}

type TransactionResult struct {
	Transaction Transaction
	Err         error
}

func (t *RecentTransactionsResponse) Prev() (*utils.Response[RecentTransactionsResponse], error) {
	// implementation here
	if t.Links.Prev == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Prev)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[RecentTransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[RecentTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[RecentTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *RecentTransactionsResponse) Next() (*utils.Response[RecentTransactionsResponse], error) {
	// implementation here
	if t.Links.Next == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Next)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[RecentTransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[RecentTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[RecentTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *TransactionsResponse) Prev() (*utils.Response[TransactionsResponse], error) {
	// implementation here
	if t.Links.Prev == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Prev)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *TransactionsResponse) Next() (*utils.Response[TransactionsResponse], error) {
	// implementation here
	if t.Links.Next == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Next)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *TransactionsTimeBucketResponse) Prev() (*utils.Response[TransactionsTimeBucketResponse], error) {
	// implementation here
	if t.Links.Prev == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Prev)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsTimeBucketResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsTimeBucketResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsTimeBucketResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *TransactionsTimeBucketResponse) Next() (*utils.Response[TransactionsTimeBucketResponse], error) {
	// implementation here
	if t.Links.Next == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Next)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsTimeBucketResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsTimeBucketResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsTimeBucketResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *TransactionsBlockPageResponse) Prev() (*utils.Response[TransactionsBlockPageResponse], error) {
	// implementation here
	if t.Links.Prev == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Prev)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsBlockPageResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsBlockPageResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsBlockPageResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

func (t *TransactionsBlockPageResponse) Next() (*utils.Response[TransactionsBlockPageResponse], error) {
	// implementation here
	if t.Links.Next == nil {
		errorCode := 400
		errorMessage := "Invalid URL: URL link cannot be null"
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf("Invalid URL: URL link cannot be null")
	}

	if !isKeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	// Create an HTTP client
	client := &http.Client{}

	parsedURL, err := url.Parse(*t.Links.Next)
	if err != nil {
		return nil, err
	}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+clientKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debugOutput {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(clientKey, debugOutput, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsBlockPageResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsBlockPageResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsBlockPageResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil

}

type GetTransactionQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Decoded DEX details including protocol (e.g. Uniswap), event (e.g 'add_liquidity') and tokens involved with historical prices. Additional 0.05 credits charged if data available.
	WithDex *bool `json:"withDex,omitempty"`
	// Decoded NFT sales details including marketplace (e.g. Opensea) and cached media links. Additional 0.05 credits charged if data available.
	WithNftSales *bool `json:"withNftSales,omitempty"`
	// Decoded lending details including protocol (e.g. Aave), event (e.g. 'deposit') and tokens involved with prices. Additional 0.05 credits charged if data available.
	WithLending *bool `json:"withLending,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionsForAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Sort the transactions in ascending chronological order. By default, it's set to `false` and returns transactions in descending chronological order.
	BlockSignedAtAsc *bool `json:"blockSignedAtAsc,omitempty"`
}
type GetAllTransactionsForAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Sort the transactions in ascending chronological order. By default, it's set to `false` and returns transactions in descending chronological order.
	BlockSignedAtAsc *bool `json:"blockSignedAtAsc,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionsForAddressV3QueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Sort the transactions in ascending chronological order. By default, it's set to `false` and returns transactions in descending chronological order.
	BlockSignedAtAsc *bool `json:"blockSignedAtAsc,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTimeBucketTransactionsForAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetEarliestTimeBucketTransactionsForAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionsForBlockByPageQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionsForBlockQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionsForBlockHashByPageQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionsForBlockHashQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Omit log events.
	NoLogs *bool `json:"noLogs,omitempty"`
	// Include safe details.
	WithSafe *bool `json:"withSafe,omitempty"`
}
type GetTransactionSummaryQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Include gas summary details. Additional charge of 1 credit when true. Response times may be impacted for wallets with millions of transactions.
	WithGas *bool `json:"withGas,omitempty"`
}

func NewTransactionServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) TransactionService {

	clientKey = apiKey
	debugOutput = debug
	workerCount = threadCount
	isKeyValid = isValidKey

	return &transactionServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type TransactionService interface {

	// Commonly used to fetch and render a single transaction including its decoded log events. Additionally return semantically decoded information for DEX trades, lending and NFT sales.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// txHash: The transaction hash.. Type: string
	GetTransaction(chainName chains.Chain, txHash string, queryParamOpts ...GetTransactionQueryParamOpts) (*utils.Response[TransactionResponse], error)

	// Commonly used to fetch and render the most recent transactions involving an address. Frequently seen in wallet applications.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetAllTransactionsForAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetAllTransactionsForAddressQueryParamOpts) <-chan TransactionResult

	// Commonly used to fetch and render the most recent transactions involving an address. Frequently seen in wallet applications.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetAllTransactionsForAddressByPage(chainName chains.Chain, walletAddress string, queryParamOpts ...GetAllTransactionsForAddressQueryParamOpts) (*utils.Response[RecentTransactionsResponse], error)

	// Commonly used to fetch the transactions involving an address including the decoded log events in a paginated fashion.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// page: The requested page, 0-indexed.. Type: int
	GetTransactionsForAddressV3(chainName chains.Chain, walletAddress string, page int, queryParamOpts ...GetTransactionsForAddressV3QueryParamOpts) (*utils.Response[TransactionsResponse], error)

	// Commonly used to fetch all transactions including their decoded log events in a 15-minute time bucket interval.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// timeBucket: The 0-indexed 15-minute time bucket. E.g. 27 Feb 2023 05:23 GMT = 1677475383 (Unix time). 1677475383/900=1863861 timeBucket.. Type: int
	GetTimeBucketTransactionsForAddress(chainName chains.Chain, walletAddress string, timeBucket int, queryParamOpts ...GetTimeBucketTransactionsForAddressQueryParamOpts) (*utils.Response[TransactionsTimeBucketResponse], error)

	// Commonly used to fetch all transactions including their decoded log events in a block and further flag interesting wallets or transactions.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// blockHeight: The requested block height.. Type: string
	GetTransactionsForBlock(chainName chains.Chain, blockHeight string, queryParamOpts ...GetTransactionsForBlockQueryParamOpts) (*utils.Response[TransactionsBlockResponse], error)

	// undefined
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// blockHash: The requested block hash.. Type: string
	// page: The requested 0-indexed page number.. Type: int
	GetTransactionsForBlockHashByPage(chainName chains.Chain, blockHash string, page int, queryParamOpts ...GetTransactionsForBlockHashByPageQueryParamOpts) (*utils.Response[TransactionsBlockPageResponse], error)

	// Commonly used to fetch all transactions including their decoded log events in a block and further flag interesting wallets or transactions.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// blockHash: The requested block hash.. Type: string
	GetTransactionsForBlockHash(chainName chains.Chain, blockHash string, queryParamOpts ...GetTransactionsForBlockHashQueryParamOpts) (*utils.Response[TransactionsBlockResponse], error)

	// Commonly used to fetch the earliest and latest transactions, and the transaction count for a wallet. Calculate the age of the wallet and the time it has been idle and quickly gain insights into their engagement with web3.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTransactionSummary(chainName chains.Chain, walletAddress string, queryParamOpts ...GetTransactionSummaryQueryParamOpts) (*utils.Response[TransactionsSummaryResponse], error)
}

type transactionServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *transactionServiceImpl) GetTransaction(chainName chains.Chain, txHash string, queryParamOpts ...GetTransactionQueryParamOpts) (*utils.Response[TransactionResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/transaction_v2/%s/", chainName, txHash)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.WithDex != nil {
			params.Add("with-dex", fmt.Sprintf("%v", *opts.WithDex))
		}

		if opts.WithNftSales != nil {
			params.Add("with-nft-sales", fmt.Sprintf("%v", *opts.WithNftSales))
		}

		if opts.WithLending != nil {
			params.Add("with-lending", fmt.Sprintf("%v", *opts.WithLending))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetAllTransactionsForAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetAllTransactionsForAddressQueryParamOpts) <-chan TransactionResult {
	transactionChannel := make(chan TransactionResult)

	go func() {
		defer close(transactionChannel)

		hasNext := true

		if !s.IskeyValid {
			transactionChannel <- TransactionResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/transactions_v3/", chainName, walletAddress)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			transactionChannel <- TransactionResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.QuoteCurrency != nil {
				params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
			}

			if opts.NoLogs != nil {
				params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
			}

			if opts.BlockSignedAtAsc != nil {
				params.Add("block-signed-at-asc", fmt.Sprintf("%v", *opts.BlockSignedAtAsc))
			}

			if opts.WithSafe != nil {
				params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
			}

		}

		// Add query parameters to the URL
		parsedURL.RawQuery = params.Encode()

		var data utils.Response[RecentTransactionsResponse]
		for hasNext {

			res, err := utils.PaginateEndpointUsingLinks(apiURL, s.APIKey, params, s.Debug, s.ThreadCount, utils.UserAgent)
			if err != nil {
				transactionChannel <- TransactionResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				transactionChannel <- TransactionResult{Err: err}
				res.Body.Close()
				hasNext = false
				return
			}
			res.Body.Close() // Ensure the body is closed after processing it

			if data.Error {
				var errorMessage string
				if data.ErrorMessage != nil {
					errorMessage = *data.ErrorMessage
				} else {
					errorMessage = "default error message" // Provide a default or handle differently
				}
				transactionChannel <- TransactionResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				transactionChannel <- TransactionResult{Transaction: item, Err: err}
			}

			if data.Data.Links.Prev == nil {
				hasNext = false
			} else {
				// Dereference Prev pointer to get its string value
				apiURL = *data.Data.Links.Prev
			}
		}

	}()
	return transactionChannel
}

func (s *transactionServiceImpl) GetAllTransactionsForAddressByPage(chainName chains.Chain, walletAddress string, queryParamOpts ...GetAllTransactionsForAddressQueryParamOpts) (*utils.Response[RecentTransactionsResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/transactions_v3/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.BlockSignedAtAsc != nil {
			params.Add("block-signed-at-asc", fmt.Sprintf("%v", *opts.BlockSignedAtAsc))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[RecentTransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[RecentTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[RecentTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[RecentTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetTransactionsForAddressV3(chainName chains.Chain, walletAddress string, page int, queryParamOpts ...GetTransactionsForAddressV3QueryParamOpts) (*utils.Response[TransactionsResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/transactions_v3/page/%d/", chainName, walletAddress, page)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.BlockSignedAtAsc != nil {
			params.Add("block-signed-at-asc", fmt.Sprintf("%v", *opts.BlockSignedAtAsc))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetTimeBucketTransactionsForAddress(chainName chains.Chain, walletAddress string, timeBucket int, queryParamOpts ...GetTimeBucketTransactionsForAddressQueryParamOpts) (*utils.Response[TransactionsTimeBucketResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/bulk/transactions/%s/%d/", chainName, walletAddress, timeBucket)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsTimeBucketResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsTimeBucketResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsTimeBucketResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsTimeBucketResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetTransactionsForBlock(chainName chains.Chain, blockHeight string, queryParamOpts ...GetTransactionsForBlockQueryParamOpts) (*utils.Response[TransactionsBlockResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/block/%s/transactions_v3/", chainName, blockHeight)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsBlockResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsBlockResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsBlockResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetTransactionsForBlockHashByPage(chainName chains.Chain, blockHash string, page int, queryParamOpts ...GetTransactionsForBlockHashByPageQueryParamOpts) (*utils.Response[TransactionsBlockPageResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/block_hash/%s/transactions_v3/page/%d/", chainName, blockHash, page)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsBlockPageResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockPageResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsBlockPageResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsBlockPageResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetTransactionsForBlockHash(chainName chains.Chain, blockHash string, queryParamOpts ...GetTransactionsForBlockHashQueryParamOpts) (*utils.Response[TransactionsBlockResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/block_hash/%s/transactions_v3/", chainName, blockHash)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.NoLogs != nil {
			params.Add("no-logs", fmt.Sprintf("%v", *opts.NoLogs))
		}

		if opts.WithSafe != nil {
			params.Add("with-safe", fmt.Sprintf("%v", *opts.WithSafe))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsBlockResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsBlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsBlockResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsBlockResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *transactionServiceImpl) GetTransactionSummary(chainName chains.Chain, walletAddress string, queryParamOpts ...GetTransactionSummaryQueryParamOpts) (*utils.Response[TransactionsSummaryResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/transactions_summary/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.WithGas != nil {
			params.Add("with-gas", fmt.Sprintf("%v", *opts.WithGas))
		}

	}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[TransactionsSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)
	req.Header.Set("X-Requested-With", utils.UserAgent)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if s.Debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// create a &Response that returns data: nil, error: true, errorCode: "Unknown Error Code", errorMessage: "Unknown Error"
		return nil, err
	}

	defer resp.Body.Close()

	utils.DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0, utils.UserAgent)

	// // Read the response body
	var data utils.Response[TransactionsSummaryResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsSummaryResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsSummaryResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
