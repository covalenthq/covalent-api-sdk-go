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

type BlockResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []Block `json:"items"`
}
type Block struct {
	// The hash of the block.
	BlockHash *string `json:"block_hash,omitempty"`
	// The block signed timestamp in UTC.
	SignedAt *time.Time `json:"signed_at,omitempty"`
	// The block height.
	Height *int `json:"height,omitempty"`
	// The parent block hash.
	BlockParentHash *string `json:"block_parent_hash,omitempty"`
	// Extra data written to the block.
	ExtraData *string `json:"extra_data,omitempty"`
	// The address of the miner.
	MinerAddress *string `json:"miner_address,omitempty"`
	// The associated mining cost.
	MiningCost *int64 `json:"mining_cost,omitempty"`
	// The associated gas used.
	GasUsed *int64 `json:"gas_used,omitempty"`
	// The associated gas limit.
	GasLimit *int64 `json:"gas_limit,omitempty"`
	// The link to the related tx by block endpoint.
	TransactionsLink *string `json:"transactions_link,omitempty"`
}
type ResolvedAddress struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []ResolvedAddressItem `json:"items"`
}
type ResolvedAddressItem struct {
	// The requested address.
	Address *string `json:"address,omitempty"`
	Name    *string `json:"name,omitempty"`
}
type BlockHeightsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []BlockHeights `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type BlockHeights struct {
	// The block signed timestamp in UTC.
	SignedAt *time.Time `json:"signed_at,omitempty"`
	// The block height.
	Height *int `json:"height,omitempty"`
}
type GetLogsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []GetLogsEvent `json:"items"`
}
type GetLogsEvent struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
	// The hash of the block.
	BlockHash *string `json:"block_hash,omitempty"`
	// The offset is the position of the tx in the block.
	TxOffset *int `json:"tx_offset,omitempty"`
	// The offset is the position of the log entry within an event log.
	LogOffset *int `json:"log_offset,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	// The log topics in raw data.
	RawLogTopics *[]string `json:"raw_log_topics,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	SenderContractDecimals *int `json:"sender_contract_decimals,omitempty"`
	// The name of the sender.
	SenderName *string `json:"sender_name,omitempty"`
	// The ticker symbol for the sender. This field is set by a developer and non-unique across a network.
	SenderContractTickerSymbol *string `json:"sender_contract_ticker_symbol,omitempty"`
	// The address of the sender.
	SenderAddress *string `json:"sender_address,omitempty"`
	// The label of the sender address.
	SenderAddressLabel *string `json:"sender_address_label,omitempty"`
	// The contract logo URL.
	SenderLogoUrl *string `json:"sender_logo_url,omitempty"`
	// The address of the deployed UniswapV2 like factory contract for this DEX.
	SenderFactoryAddress *string `json:"sender_factory_address,omitempty"`
	// The log events in raw.
	RawLogData *string `json:"raw_log_data,omitempty"`
	// The decoded item.
	Decoded *genericmodels.DecodedItem `json:"decoded,omitempty"`
}
type LogEventsByAddressResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []genericmodels.LogEvent `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type LogEventsByTopicHashResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []genericmodels.LogEvent `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type AllChainsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []ChainItem `json:"items"`
}
type ChainItem struct {
	// The chain name eg: `eth-mainnet`.
	Name *string `json:"name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// True if the chain is a testnet.
	IsTestnet *bool `json:"is_testnet,omitempty"`
	// Schema name to use for direct SQL.
	DbSchemaName *string `json:"db_schema_name,omitempty"`
	// The chains label eg: `Ethereum Mainnet`.
	Label *string `json:"label,omitempty"`
	// The category label eg: `Ethereum`.
	CategoryLabel *string `json:"category_label,omitempty"`
	// A svg logo url for the chain.
	LogoUrl *string `json:"logo_url,omitempty"`
	// A black png logo url for the chain.
	BlackLogoUrl *string `json:"black_logo_url,omitempty"`
	// A white png logo url for the chain.
	WhiteLogoUrl *string `json:"white_logo_url,omitempty"`
	// The color theme for the chain.
	ColorTheme *ColorTheme `json:"color_theme,omitempty"`
	// True if the chain is an AppChain.
	IsAppchain *bool `json:"is_appchain,omitempty"`
	// The ChainItem the appchain is a part of.
	AppchainOf *ChainItem `json:"appchain_of,omitempty"`
}
type ColorTheme struct {
	// The red color code.
	Red *int `json:"red,omitempty"`
	// The green color code.
	Green *int `json:"green,omitempty"`
	// The blue color code.
	Blue *int `json:"blue,omitempty"`
	// The alpha color code.
	Alpha *int `json:"alpha,omitempty"`
	// The hexadecimal color code.
	Hex *string `json:"hex,omitempty"`
	// The color represented in css rgb() functional notation.
	CssRgb *string `json:"css_rgb,omitempty"`
}
type AllChainsStatusResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []ChainStatusItem `json:"items"`
}
type ChainStatusItem struct {
	// The chain name eg: `eth-mainnet`.
	Name *string `json:"name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// True if the chain is a testnet.
	IsTestnet *bool `json:"is_testnet,omitempty"`
	// A svg logo url for the chain.
	LogoUrl *string `json:"logo_url,omitempty"`
	// A black png logo url for the chain.
	BlackLogoUrl *string `json:"black_logo_url,omitempty"`
	// A white png logo url for the chain.
	WhiteLogoUrl *string `json:"white_logo_url,omitempty"`
	// True if the chain is an AppChain.
	IsAppchain *bool `json:"is_appchain,omitempty"`
	// The height of the lastest block available.
	SyncedBlockHeight *int `json:"synced_block_height,omitempty"`
	// The signed timestamp of lastest block available.
	SyncedBlockedSignedAt *time.Time `json:"synced_blocked_signed_at,omitempty"`
	// True if the chain has data and ready for querying.
	HasData *bool `json:"has_data,omitempty"`
}
type ChainActivityResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// List of response items.
	Items []ChainActivityEvent `json:"items"`
}
type ChainActivityEvent struct {
	ChainItem
	// The timestamp when the address was last seen on the chain.
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
}
type GasPricesResponse struct {
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested event type.
	EventType string `json:"event_type"`
	// The exchange rate for the requested quote currency.
	GasQuoteRate float64 `json:"gas_quote_rate"`
	// The lowest gas fee for the latest block height.
	BaseFee utils.BigInt `json:"base_fee"`
	// List of response items.
	Items []PriceItem `json:"items"`
}
type PriceItem struct {
	// The average gas price, in WEI, for the time interval.
	GasPrice *string `json:"gas_price,omitempty"`
	// The average gas spent for the time interval.
	GasSpent *string `json:"gas_spent,omitempty"`
	// The average gas spent in `quote-currency` denomination for the time interval.
	GasQuote *float64 `json:"gas_quote,omitempty"`
	// Other fees, when applicable. For example: OP chain L1 fees.
	OtherFees *OtherFees `json:"other_fees,omitempty"`
	// The sum of the L1 and L2 gas spent, in quote-currency, for the specified time interval.
	TotalGasQuote *float64 `json:"total_gas_quote,omitempty"`
	// A prettier version of the total average gas spent, in quote-currency, for the specified time interval, for rendering purposes.
	PrettyTotalGasQuote *string `json:"pretty_total_gas_quote,omitempty"`
	// The specified time interval.
	Interval *string `json:"interval,omitempty"`
}
type OtherFees struct {
	// The calculated L1 gas spent, when applicable, in quote-currency, for the specified time interval.
	L1GasQuote *float64 `json:"l1_gas_quote,omitempty"`
}
type BlockHeightsResult struct {
	BlockHeights BlockHeights
	Err          error
}

type LogEventResult struct {
	LogEvent genericmodels.LogEvent
	Err      error
}

type GetBlockHeightsQueryParamOpts struct {
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetLogsQueryParamOpts struct {
	// The first block to retrieve log events with. Accepts decimals, hexadecimals, or the strings `earliest` and `latest`.
	StartingBlock *int `json:"startingBlock,omitempty"`
	// The last block to retrieve log events with. Accepts decimals, hexadecimals, or the strings `earliest` and `latest`.
	EndingBlock *string `json:"endingBlock,omitempty"`
	// The address of the log events sender contract.
	Address *string `json:"address,omitempty"`
	// The topic hash(es) to retrieve logs with.
	Topics *string `json:"topics,omitempty"`
	// The block hash to retrieve logs for.
	BlockHash *string `json:"blockHash,omitempty"`
	// Omit decoded log events.
	SkipDecode *bool `json:"skipDecode,omitempty"`
}
type GetLogEventsByAddressQueryParamOpts struct {
	// The first block to retrieve log events with. Accepts decimals, hexadecimals, or the strings `earliest` and `latest`.
	StartingBlock *int `json:"startingBlock,omitempty"`
	// The last block to retrieve log events with. Accepts decimals, hexadecimals, or the strings `earliest` and `latest`.
	EndingBlock *string `json:"endingBlock,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetLogEventsByTopicHashQueryParamOpts struct {
	// The first block to retrieve log events with. Accepts decimals, hexadecimals, or the strings `earliest` and `latest`.
	StartingBlock *int `json:"startingBlock,omitempty"`
	// The last block to retrieve log events with. Accepts decimals, hexadecimals, or the strings `earliest` and `latest`.
	EndingBlock *string `json:"endingBlock,omitempty"`
	// Additional topic hash(es) to filter on - padded & unpadded address fields are supported. Separate multiple topics with a comma.
	SecondaryTopics *string `json:"secondaryTopics,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetAddressActivityQueryParamOpts struct {
	// Set to true to include testnets with activity in the response. By default, it's set to `false` and only returns mainnet activity.
	Testnets *bool `json:"testnets,omitempty"`
}
type GetGasPricesQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
}

func NewBaseServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) BaseService {

	return &baseServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type BaseService interface {

	// Commonly used to fetch and render a single block for a block explorer.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// blockHeight: The block height or `latest` for the latest block available.. Type: string
	GetBlock(chainName chains.Chain, blockHeight string) (*utils.Response[BlockResponse], error)

	// Commonly used to resolve ENS, RNS and Unstoppable Domains addresses.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetResolvedAddress(chainName chains.Chain, walletAddress string) (*utils.Response[ResolvedAddress], error)

	// Commonly used to get all the block heights within a particular date range. Useful for rendering a display where you sort blocks by day.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// startDate: The start date in YYYY-MM-DD format.. Type: string
	// endDate: The end date in YYYY-MM-DD format.. Type: string
	GetBlockHeights(chainName chains.Chain, startDate string, endDate string, queryParamOpts ...GetBlockHeightsQueryParamOpts) <-chan BlockHeightsResult

	// Commonly used to get all the block heights within a particular date range. Useful for rendering a display where you sort blocks by day.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// startDate: The start date in YYYY-MM-DD format.. Type: string
	// endDate: The end date in YYYY-MM-DD format.. Type: string
	GetBlockHeightsByPage(chainName chains.Chain, startDate string, endDate string, queryParamOpts ...GetBlockHeightsQueryParamOpts) (*utils.Response[BlockHeightsResponse], error)

	// Commonly used to get all the event logs of the latest block, or for a range of blocks. Includes sender contract metadata as well as decoded logs.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	GetLogs(chainName chains.Chain, queryParamOpts ...GetLogsQueryParamOpts) (*utils.Response[GetLogsResponse], error)

	// Commonly used to get all the event logs emitted from a particular contract address. Useful for building dashboards that examine on-chain interactions.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetLogEventsByAddress(chainName chains.Chain, contractAddress string, queryParamOpts ...GetLogEventsByAddressQueryParamOpts) <-chan LogEventResult

	// Commonly used to get all the event logs emitted from a particular contract address. Useful for building dashboards that examine on-chain interactions.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetLogEventsByAddressByPage(chainName chains.Chain, contractAddress string, queryParamOpts ...GetLogEventsByAddressQueryParamOpts) (*utils.Response[LogEventsByAddressResponse], error)

	// Commonly used to get all event logs of the same topic hash across all contracts within a particular chain. Useful for cross-sectional analysis of event logs that are emitted on-chain.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// topicHash: The endpoint will return event logs that contain this topic hash.. Type: string
	GetLogEventsByTopicHash(chainName chains.Chain, topicHash string, queryParamOpts ...GetLogEventsByTopicHashQueryParamOpts) <-chan LogEventResult

	// Commonly used to get all event logs of the same topic hash across all contracts within a particular chain. Useful for cross-sectional analysis of event logs that are emitted on-chain.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// topicHash: The endpoint will return event logs that contain this topic hash.. Type: string
	GetLogEventsByTopicHashByPage(chainName chains.Chain, topicHash string, queryParamOpts ...GetLogEventsByTopicHashQueryParamOpts) (*utils.Response[LogEventsByTopicHashResponse], error)

	// Commonly used to build internal dashboards for all supported chains on Covalent.
	//   Parameters:

	GetAllChains() (*utils.Response[AllChainsResponse], error)

	// Commonly used to build internal status dashboards of all supported chains.
	//   Parameters:

	GetAllChainStatus() (*utils.Response[AllChainsStatusResponse], error)

	// Commonly used to locate chains which an address is active on with a single API call.
	//   Parameters:
	// walletAddress: The requested wallet address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetAddressActivity(walletAddress string, queryParamOpts ...GetAddressActivityQueryParamOpts) (*utils.Response[ChainActivityResponse], error)

	// Get real-time gas estimates for different transaction speeds on a specific network, enabling users to optimize transaction costs and confirmation times.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// eventType: The desired event type to retrieve gas prices for. Supports `erc20` transfer events, `uniswapv3` swap events and `nativetokens` transfers.. Type: string
	GetGasPrices(chainName chains.Chain, eventType string, queryParamOpts ...GetGasPricesQueryParamOpts) (*utils.Response[GasPricesResponse], error)
}

type baseServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *baseServiceImpl) GetBlock(chainName chains.Chain, blockHeight string) (*utils.Response[BlockResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/block_v2/%s/", chainName, blockHeight)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[BlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[BlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[BlockResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[BlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[BlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[BlockResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[BlockResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[BlockResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetResolvedAddress(chainName chains.Chain, walletAddress string) (*utils.Response[ResolvedAddress], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/resolve_address/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[ResolvedAddress]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[ResolvedAddress]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[ResolvedAddress]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[ResolvedAddress]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ResolvedAddress]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ResolvedAddress]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[ResolvedAddress]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[ResolvedAddress]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetBlockHeights(chainName chains.Chain, startDate string, endDate string, queryParamOpts ...GetBlockHeightsQueryParamOpts) <-chan BlockHeightsResult {
	blockHeightsChannel := make(chan BlockHeightsResult)

	go func() {
		defer close(blockHeightsChannel)

		hasNext := true

		if !s.IskeyValid {
			blockHeightsChannel <- BlockHeightsResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/block_v2/%s/%s/", chainName, startDate, endDate)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			blockHeightsChannel <- BlockHeightsResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.PageSize != nil {
				params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
			}

			if opts.PageNumber != nil {
				params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
			}

		}

		// Add query parameters to the URL
		parsedURL.RawQuery = params.Encode()
		page := 0

		if !params.Has("page-number") {
			page = 0
		} else {
			page, err = strconv.Atoi(params.Get("page-number"))
			if err != nil {
				blockHeightsChannel <- BlockHeightsResult{Err: err}
				return
			}
		}

		var data utils.Response[BlockHeightsResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount)
			if err != nil {
				blockHeightsChannel <- BlockHeightsResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				blockHeightsChannel <- BlockHeightsResult{Err: err}
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
				blockHeightsChannel <- BlockHeightsResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				blockHeightsChannel <- BlockHeightsResult{BlockHeights: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return blockHeightsChannel
}

func (s *baseServiceImpl) GetBlockHeightsByPage(chainName chains.Chain, startDate string, endDate string, queryParamOpts ...GetBlockHeightsQueryParamOpts) (*utils.Response[BlockHeightsResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/block_v2/%s/%s/", chainName, startDate, endDate)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[BlockHeightsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.PageSize != nil {
			params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
		}

		if opts.PageNumber != nil {
			params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
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
		return &utils.Response[BlockHeightsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[BlockHeightsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[BlockHeightsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[BlockHeightsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[BlockHeightsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[BlockHeightsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[BlockHeightsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetLogs(chainName chains.Chain, queryParamOpts ...GetLogsQueryParamOpts) (*utils.Response[GetLogsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/events/", chainName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[GetLogsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.StartingBlock != nil {
			params.Add("starting-block", fmt.Sprintf("%v", *opts.StartingBlock))
		}

		if opts.EndingBlock != nil {
			params.Add("ending-block", fmt.Sprintf("%v", *opts.EndingBlock))
		}

		if opts.Address != nil {
			params.Add("address", fmt.Sprintf("%v", *opts.Address))
		}

		if opts.Topics != nil {
			params.Add("topics", fmt.Sprintf("%v", *opts.Topics))
		}

		if opts.BlockHash != nil {
			params.Add("block-hash", fmt.Sprintf("%v", *opts.BlockHash))
		}

		if opts.SkipDecode != nil {
			params.Add("skip-decode", fmt.Sprintf("%v", *opts.SkipDecode))
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
		return &utils.Response[GetLogsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[GetLogsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[GetLogsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[GetLogsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[GetLogsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[GetLogsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[GetLogsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetLogEventsByAddress(chainName chains.Chain, contractAddress string, queryParamOpts ...GetLogEventsByAddressQueryParamOpts) <-chan LogEventResult {
	logEventChannel := make(chan LogEventResult)

	go func() {
		defer close(logEventChannel)

		hasNext := true

		if !s.IskeyValid {
			logEventChannel <- LogEventResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/events/address/%s/", chainName, contractAddress)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			logEventChannel <- LogEventResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.StartingBlock != nil {
				params.Add("starting-block", fmt.Sprintf("%v", *opts.StartingBlock))
			}

			if opts.EndingBlock != nil {
				params.Add("ending-block", fmt.Sprintf("%v", *opts.EndingBlock))
			}

			if opts.PageSize != nil {
				params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
			}

			if opts.PageNumber != nil {
				params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
			}

		}

		// Add query parameters to the URL
		parsedURL.RawQuery = params.Encode()
		page := 0

		if !params.Has("page-number") {
			page = 0
		} else {
			page, err = strconv.Atoi(params.Get("page-number"))
			if err != nil {
				logEventChannel <- LogEventResult{Err: err}
				return
			}
		}

		var data utils.Response[LogEventsByAddressResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount)
			if err != nil {
				logEventChannel <- LogEventResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				logEventChannel <- LogEventResult{Err: err}
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
				logEventChannel <- LogEventResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				logEventChannel <- LogEventResult{LogEvent: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return logEventChannel
}

func (s *baseServiceImpl) GetLogEventsByAddressByPage(chainName chains.Chain, contractAddress string, queryParamOpts ...GetLogEventsByAddressQueryParamOpts) (*utils.Response[LogEventsByAddressResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/events/address/%s/", chainName, contractAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[LogEventsByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.StartingBlock != nil {
			params.Add("starting-block", fmt.Sprintf("%v", *opts.StartingBlock))
		}

		if opts.EndingBlock != nil {
			params.Add("ending-block", fmt.Sprintf("%v", *opts.EndingBlock))
		}

		if opts.PageSize != nil {
			params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
		}

		if opts.PageNumber != nil {
			params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
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
		return &utils.Response[LogEventsByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[LogEventsByAddressResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[LogEventsByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[LogEventsByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[LogEventsByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[LogEventsByAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[LogEventsByAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetLogEventsByTopicHash(chainName chains.Chain, topicHash string, queryParamOpts ...GetLogEventsByTopicHashQueryParamOpts) <-chan LogEventResult {
	logEventChannel := make(chan LogEventResult)

	go func() {
		defer close(logEventChannel)

		hasNext := true

		if !s.IskeyValid {
			logEventChannel <- LogEventResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/events/topics/%s/", chainName, topicHash)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			logEventChannel <- LogEventResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.StartingBlock != nil {
				params.Add("starting-block", fmt.Sprintf("%v", *opts.StartingBlock))
			}

			if opts.EndingBlock != nil {
				params.Add("ending-block", fmt.Sprintf("%v", *opts.EndingBlock))
			}

			if opts.SecondaryTopics != nil {
				params.Add("secondary-topics", fmt.Sprintf("%v", *opts.SecondaryTopics))
			}

			if opts.PageSize != nil {
				params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
			}

			if opts.PageNumber != nil {
				params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
			}

		}

		// Add query parameters to the URL
		parsedURL.RawQuery = params.Encode()
		page := 0

		if !params.Has("page-number") {
			page = 0
		} else {
			page, err = strconv.Atoi(params.Get("page-number"))
			if err != nil {
				logEventChannel <- LogEventResult{Err: err}
				return
			}
		}

		var data utils.Response[LogEventsByTopicHashResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount)
			if err != nil {
				logEventChannel <- LogEventResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				logEventChannel <- LogEventResult{Err: err}
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
				logEventChannel <- LogEventResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				logEventChannel <- LogEventResult{LogEvent: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return logEventChannel
}

func (s *baseServiceImpl) GetLogEventsByTopicHashByPage(chainName chains.Chain, topicHash string, queryParamOpts ...GetLogEventsByTopicHashQueryParamOpts) (*utils.Response[LogEventsByTopicHashResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/events/topics/%s/", chainName, topicHash)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[LogEventsByTopicHashResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.StartingBlock != nil {
			params.Add("starting-block", fmt.Sprintf("%v", *opts.StartingBlock))
		}

		if opts.EndingBlock != nil {
			params.Add("ending-block", fmt.Sprintf("%v", *opts.EndingBlock))
		}

		if opts.SecondaryTopics != nil {
			params.Add("secondary-topics", fmt.Sprintf("%v", *opts.SecondaryTopics))
		}

		if opts.PageSize != nil {
			params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
		}

		if opts.PageNumber != nil {
			params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
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
		return &utils.Response[LogEventsByTopicHashResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[LogEventsByTopicHashResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[LogEventsByTopicHashResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[LogEventsByTopicHashResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[LogEventsByTopicHashResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[LogEventsByTopicHashResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[LogEventsByTopicHashResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetAllChains() (*utils.Response[AllChainsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/chains/")

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[AllChainsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[AllChainsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[AllChainsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[AllChainsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[AllChainsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[AllChainsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[AllChainsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[AllChainsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetAllChainStatus() (*utils.Response[AllChainsStatusResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/chains/status/")

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[AllChainsStatusResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}

	// Add query parameters to the URL
	parsedURL.RawQuery = params.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		errorCode := 500
		errorMessage := "Unknown Error"
		return &utils.Response[AllChainsStatusResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[AllChainsStatusResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[AllChainsStatusResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[AllChainsStatusResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[AllChainsStatusResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[AllChainsStatusResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[AllChainsStatusResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetAddressActivity(walletAddress string, queryParamOpts ...GetAddressActivityQueryParamOpts) (*utils.Response[ChainActivityResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/address/%s/activity/", walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[ChainActivityResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.Testnets != nil {
			params.Add("testnets", fmt.Sprintf("%v", *opts.Testnets))
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
		return &utils.Response[ChainActivityResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[ChainActivityResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[ChainActivityResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ChainActivityResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ChainActivityResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[ChainActivityResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[ChainActivityResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *baseServiceImpl) GetGasPrices(chainName chains.Chain, eventType string, queryParamOpts ...GetGasPricesQueryParamOpts) (*utils.Response[GasPricesResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/event/%s/gas_prices/", chainName, eventType)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[GasPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[GasPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
	}

	req.Header.Set("Authorization", `Bearer `+s.APIKey)

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
	backoff := utils.NewExponentialBackoff(s.APIKey, s.Debug, 0)

	// // Read the response body
	var data utils.Response[GasPricesResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[GasPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[GasPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[GasPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[GasPricesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[GasPricesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
