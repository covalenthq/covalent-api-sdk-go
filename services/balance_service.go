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

type BalancesResponse struct {
	// The requested address.
	Address string `json:"address"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []BalanceItem `json:"items"`
}
type BalanceItem struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A display-friendly name for the contract.
	ContractDisplayName *string `json:"contract_display_name,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *[]string `json:"supports_erc,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The contract logo URLs.
	LogoUrls *LogoUrls `json:"logo_urls,omitempty"`
	// The timestamp when the token was transferred.
	LastTransferredAt *time.Time `json:"last_transferred_at,omitempty"`
	// Indicates if a token is the chain's native gas token, eg: ETH on Ethereum.
	NativeToken *bool `json:"native_token,omitempty"`
	// One of `cryptocurrency`, `stablecoin`, `nft` or `dust`.
	Type *string `json:"type,omitempty"`
	// Denotes whether the token is suspected spam.
	IsSpam *bool `json:"is_spam,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// The 24h asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance24h *utils.BigInt `json:"balance_24h,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The 24h exchange rate for the requested quote currency.
	QuoteRate24h *float64 `json:"quote_rate_24h,omitempty"`
	// The current balance converted to fiat in `quote-currency`.
	Quote *float64 `json:"quote,omitempty"`
	// The 24h balance converted to fiat in `quote-currency`.
	Quote24h *float64 `json:"quote_24h,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyQuote *string `json:"pretty_quote,omitempty"`
	// A prettier version of the 24h quote for rendering purposes.
	PrettyQuote24h *string `json:"pretty_quote_24h,omitempty"`
	// The protocol metadata.
	ProtocolMetadata *ProtocolMetadata `json:"protocol_metadata,omitempty"`
	// NFT-specific data.
	NftData *[]BalanceNftData `json:"nft_data,omitempty"`
}
type ProtocolMetadata struct {
	// The name of the protocol.
	ProtocolName *string `json:"protocol_name,omitempty"`
}
type BalanceNftData struct {
	// The token's id.
	TokenId *utils.BigInt `json:"token_id,omitempty"`
	// The count of the number of NFTs with this ID.
	TokenBalance *utils.BigInt `json:"token_balance,omitempty"`
	// External URL for additional metadata.
	TokenUrl *string `json:"token_url,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *[]string `json:"supports_erc,omitempty"`
	// The latest price value on chain of the token ID.
	TokenPriceWei *utils.BigInt `json:"token_price_wei,omitempty"`
	// The latest quote_rate of the token ID denominated in unscaled ETH.
	TokenQuoteRateEth *string `json:"token_quote_rate_eth,omitempty"`
	// The address of the original owner of this NFT.
	OriginalOwner *string            `json:"original_owner,omitempty"`
	ExternalData  *NftExternalDataV1 `json:"external_data,omitempty"`
	// The current owner of this NFT.
	Owner *string `json:"owner,omitempty"`
	// The address of the current owner of this NFT.
	OwnerAddress *string `json:"owner_address,omitempty"`
	// When set to true, this NFT has been Burned.
	Burned *bool `json:"burned,omitempty"`
}
type NftExternalDataV1 struct {
	Name         *string                                 `json:"name,omitempty"`
	Description  *string                                 `json:"description,omitempty"`
	Image        *string                                 `json:"image,omitempty"`
	Image256     *string                                 `json:"image_256,omitempty"`
	Image512     *string                                 `json:"image_512,omitempty"`
	Image1024    *string                                 `json:"image_1024,omitempty"`
	AnimationUrl *string                                 `json:"animation_url,omitempty"`
	ExternalUrl  *string                                 `json:"external_url,omitempty"`
	Attributes   *[]genericmodels.NftCollectionAttribute `json:"attributes,omitempty"`
	Owner        *string                                 `json:"owner,omitempty"`
}
type PortfolioResponse struct {
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
	// List of response items.
	Items []PortfolioItem `json:"items"`
}
type PortfolioItem struct {
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// The contract logo URL.
	LogoUrl  *string        `json:"logo_url,omitempty"`
	Holdings *[]HoldingItem `json:"holdings,omitempty"`
}
type HoldingItem struct {
	// The exchange rate for the requested quote currency.
	QuoteRate *float64   `json:"quote_rate,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Close     *OhlcItem  `json:"close,omitempty"`
	High      *OhlcItem  `json:"high,omitempty"`
	Low       *OhlcItem  `json:"low,omitempty"`
	Open      *OhlcItem  `json:"open,omitempty"`
}
type OhlcItem struct {
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// The current balance converted to fiat in `quote-currency`.
	Quote *float64 `json:"quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyQuote *string `json:"pretty_quote,omitempty"`
}
type Erc20TransfersResponse struct {
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
	// List of response items.
	Items []BlockTransactionWithContractTransfers `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type BlockTransactionWithContractTransfers struct {
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
	// Whether or not transaction is successful.
	Successful *bool `json:"successful,omitempty"`
	// The address of the miner.
	MinerAddress *string `json:"miner_address,omitempty"`
	// The sender's wallet address.
	FromAddress *string `json:"from_address,omitempty"`
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
	// The transaction's gas_price * gas_spent, denoted in wei.
	FeesPaid *utils.BigInt `json:"fees_paid,omitempty"`
	// The gas spent in `quote-currency` denomination.
	GasQuote *float64 `json:"gas_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyGasQuote *string `json:"pretty_gas_quote,omitempty"`
	// The native gas exchange rate for the requested `quote-currency`.
	GasQuoteRate *float64             `json:"gas_quote_rate,omitempty"`
	Transfers    *[]TokenTransferItem `json:"transfers,omitempty"`
}
type TokenTransferItem struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	// The sender's wallet address.
	FromAddress *string `json:"from_address,omitempty"`
	// The label of `from` address.
	FromAddressLabel *string `json:"from_address_label,omitempty"`
	// The receiver's wallet address.
	ToAddress *string `json:"to_address,omitempty"`
	// The label of `to` address.
	ToAddressLabel *string `json:"to_address_label,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// Categorizes token transactions as either `transfer-in` or `transfer-out`, indicating whether tokens are being received or sent from an account.
	TransferType *string `json:"transfer_type,omitempty"`
	// The delta attached to this transfer.
	Delta *utils.BigInt `json:"delta,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The current delta converted to fiat in `quote-currency`.
	DeltaQuote *float64 `json:"delta_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyDeltaQuote *string `json:"pretty_delta_quote,omitempty"`
	// The current balance converted to fiat in `quote-currency`.
	BalanceQuote *float64 `json:"balance_quote,omitempty"`
	// Additional details on which transfer events were invoked. Defaults to `true`.
	MethodCalls *[]MethodCallsForTransfers `json:"method_calls,omitempty"`
	// The explorer links for this transaction.
	Explorers *[]genericmodels.Explorer `json:"explorers,omitempty"`
}
type MethodCallsForTransfers struct {
	// The address of the sender.
	SenderAddress *string `json:"sender_address,omitempty"`
	Method        *string `json:"method,omitempty"`
}
type TokenHoldersResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []TokenHolder `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type TokenHolder struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *[]string `json:"supports_erc,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The requested address.
	Address *string `json:"address,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// Total supply of this token.
	TotalSupply *utils.BigInt `json:"total_supply,omitempty"`
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
}
type TokenHoldersChangesResponse struct {
	// The token holder.
	TokenHolder string `json:"token_holder"`
	// The starting block balance.
	PrevBalance string `json:"prev_balance"`
	// The starting block height.
	PrevBlockHeight int64 `json:"prev_block_height"`
	// The ending block balance.
	NextBalance string `json:"next_balance"`
	// The ending block height.
	NextBlockHeight int64 `json:"next_block_height"`
	// The difference of the balance.
	Diff string `json:"diff"`
}
type HistoricalBalancesResponse struct {
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
	// List of response items.
	Items []HistoricalBalanceItem `json:"items"`
}
type HistoricalBalanceItem struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *[]string `json:"supports_erc,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
	// The block height when the token was last transferred.
	LastTransferredBlockHeight *int64  `json:"last_transferred_block_height,omitempty"`
	ContractDisplayName        *string `json:"contract_display_name,omitempty"`
	// The timestamp when the token was transferred.
	LastTransferredAt *time.Time `json:"last_transferred_at,omitempty"`
	// Indicates if a token is the chain's native gas token, eg: ETH on Ethereum.
	NativeToken *bool `json:"native_token,omitempty"`
	// One of `cryptocurrency`, `stablecoin`, `nft` or `dust`.
	Type *string `json:"type,omitempty"`
	// Denotes whether the token is suspected spam.
	IsSpam *bool `json:"is_spam,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The current balance converted to fiat in `quote-currency`.
	Quote *float64 `json:"quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyQuote *string `json:"pretty_quote,omitempty"`
	// The protocol metadata.
	ProtocolMetadata *ProtocolMetadata `json:"protocol_metadata,omitempty"`
	// NFT-specific data.
	NftData *[]BalanceNftData `json:"nft_data,omitempty"`
}
type TokenBalanceNativeResponse struct {
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
	// List of response items.
	Items []NativeBalanceItem `json:"items"`
}
type NativeBalanceItem struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *[]string `json:"supports_erc,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The current balance converted to fiat in `quote-currency`.
	Quote *float64 `json:"quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyQuote *string `json:"pretty_quote,omitempty"`
}
type BlockTransactionWithContractTransfersResult struct {
	BlockTransactionWithContractTransfers BlockTransactionWithContractTransfers
	Err                                   error
}

type TokenHolderResult struct {
	TokenHolder TokenHolder
	Err         error
}

type GetTokenBalancesForWalletAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// If `true`, NFTs will be included in the response.
	Nft *bool `json:"nft,omitempty"`
	// If `true`, only NFTs that have been cached will be included in the response. Helpful for faster response times.
	NoNftFetch *bool `json:"noNftFetch,omitempty"`
	// If `true`, the suspected spam tokens are removed. Supports `eth-mainnet` and `matic-mainnet`.
	NoSpam *bool `json:"noSpam,omitempty"`
	// If `true`, the response shape is limited to a list of collections and token ids, omitting metadata and asset information. Helpful for faster response times and wallets holding a large number of NFTs.
	NoNftAssetMetadata *bool `json:"noNftAssetMetadata,omitempty"`
}
type GetHistoricalPortfolioForWalletAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// The number of days to return data for. Defaults to 30 days.
	Days *int `json:"days,omitempty"`
}
type GetErc20TransfersForWalletAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.
	ContractAddress *string `json:"contractAddress,omitempty"`
	// The block height to start from, defaults to `0`.
	StartingBlock *int `json:"startingBlock,omitempty"`
	// The block height to end at, defaults to current block height.
	EndingBlock *int `json:"endingBlock,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetTokenHoldersV2ForTokenAddressQueryParamOpts struct {
	// Ending block to define a block range. Omitting this parameter defaults to the latest block height.
	BlockHeight *string `json:"blockHeight,omitempty"`
	// Ending date to define a block range (YYYY-MM-DD). Omitting this parameter defaults to the current date.
	Date *string `json:"date,omitempty"`
	// Number of items per page. Note: Currently, only values of `100` and `1000` are supported. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetHistoricalTokenBalancesForWalletAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// If `true`, NFTs will be included in the response.
	Nft *bool `json:"nft,omitempty"`
	// If `true`, only NFTs that have been cached will be included in the response. Helpful for faster response times.
	NoNftFetch *bool `json:"noNftFetch,omitempty"`
	// If `true`, the suspected spam tokens are removed. Supports `eth-mainnet` and `matic-mainnet`.
	NoSpam *bool `json:"noSpam,omitempty"`
	// If `true`, the response shape is limited to a list of collections and token ids, omitting metadata and asset information. Helpful for faster response times and wallets holding a large number of NFTs.
	NoNftAssetMetadata *bool `json:"noNftAssetMetadata,omitempty"`
	// Ending block to define a block range. Omitting this parameter defaults to the latest block height.
	BlockHeight *string `json:"blockHeight,omitempty"`
	// Ending date to define a block range (YYYY-MM-DD). Omitting this parameter defaults to the current date.
	Date *string `json:"date,omitempty"`
}
type GetNativeTokenBalanceQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Ending block to define a block range. Omitting this parameter defaults to the latest block height.
	BlockHeight *string `json:"blockHeight,omitempty"`
}

func NewBalanceServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) BalanceService {

	return &balanceServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type BalanceService interface {

	// Commonly used to fetch the native, fungible (ERC20), and non-fungible (ERC721 & ERC1155) tokens held by an address. Response includes spot prices and other metadata.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTokenBalancesForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetTokenBalancesForWalletAddressQueryParamOpts) (*utils.Response[BalancesResponse], error)

	// Commonly used to render a daily portfolio balance for an address broken down by the token. The timeframe is user-configurable, defaults to 30 days.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetHistoricalPortfolioForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetHistoricalPortfolioForWalletAddressQueryParamOpts) (*utils.Response[PortfolioResponse], error)

	// Commonly used to render the transfer-in and transfer-out of a token along with historical prices from an address.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetErc20TransfersForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetErc20TransfersForWalletAddressQueryParamOpts) <-chan BlockTransactionWithContractTransfersResult

	// Commonly used to render the transfer-in and transfer-out of a token along with historical prices from an address.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetErc20TransfersForWalletAddressByPage(chainName chains.Chain, walletAddress string, queryParamOpts ...GetErc20TransfersForWalletAddressQueryParamOpts) (*utils.Response[Erc20TransfersResponse], error)

	// Commonly used to get a list of all the token holders for a specified ERC20 or ERC721 token. Returns historic token holders when block-height is set (defaults to `latest`). Useful for building pie charts of token holders.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// tokenAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTokenHoldersV2ForTokenAddress(chainName chains.Chain, tokenAddress string, queryParamOpts ...GetTokenHoldersV2ForTokenAddressQueryParamOpts) <-chan TokenHolderResult

	// Commonly used to get a list of all the token holders for a specified ERC20 or ERC721 token. Returns historic token holders when block-height is set (defaults to `latest`). Useful for building pie charts of token holders.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// tokenAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTokenHoldersV2ForTokenAddressByPage(chainName chains.Chain, tokenAddress string, queryParamOpts ...GetTokenHoldersV2ForTokenAddressQueryParamOpts) (*utils.Response[TokenHoldersResponse], error)

	// Commonly used to fetch the historical native, fungible (ERC20), and non-fungible (ERC721 & ERC1155) tokens held by an address at a given block height or date. Response includes daily prices and other metadata.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetHistoricalTokenBalancesForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetHistoricalTokenBalancesForWalletAddressQueryParamOpts) (*utils.Response[HistoricalBalancesResponse], error)

	// Commonly used to get the native token balance for an address. This endpoint is required because native tokens are usually not ERC20 tokens and sometimes you want something lightweight.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetNativeTokenBalance(chainName chains.Chain, walletAddress string, queryParamOpts ...GetNativeTokenBalanceQueryParamOpts) (*utils.Response[TokenBalanceNativeResponse], error)
}

type balanceServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *balanceServiceImpl) GetTokenBalancesForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetTokenBalancesForWalletAddressQueryParamOpts) (*utils.Response[BalancesResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/balances_v2/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[BalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.Nft != nil {
			params.Add("nft", fmt.Sprintf("%v", *opts.Nft))
		}

		if opts.NoNftFetch != nil {
			params.Add("no-nft-fetch", fmt.Sprintf("%v", *opts.NoNftFetch))
		}

		if opts.NoSpam != nil {
			params.Add("no-spam", fmt.Sprintf("%v", *opts.NoSpam))
		}

		if opts.NoNftAssetMetadata != nil {
			params.Add("no-nft-asset-metadata", fmt.Sprintf("%v", *opts.NoNftAssetMetadata))
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
		return &utils.Response[BalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[BalancesResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[BalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[BalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[BalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[BalancesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[BalancesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *balanceServiceImpl) GetHistoricalPortfolioForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetHistoricalPortfolioForWalletAddressQueryParamOpts) (*utils.Response[PortfolioResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/portfolio_v2/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[PortfolioResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.Days != nil {
			params.Add("days", fmt.Sprintf("%v", *opts.Days))
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
		return &utils.Response[PortfolioResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[PortfolioResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[PortfolioResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PortfolioResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PortfolioResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[PortfolioResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[PortfolioResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *balanceServiceImpl) GetErc20TransfersForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetErc20TransfersForWalletAddressQueryParamOpts) <-chan BlockTransactionWithContractTransfersResult {
	blockTransactionWithContractTransfersChannel := make(chan BlockTransactionWithContractTransfersResult)

	go func() {
		defer close(blockTransactionWithContractTransfersChannel)

		hasNext := true

		if !s.IskeyValid {
			blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/transfers_v2/", chainName, walletAddress)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.QuoteCurrency != nil {
				params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
			}

			if opts.ContractAddress != nil {
				params.Add("contract-address", fmt.Sprintf("%v", *opts.ContractAddress))
			}

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
				blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{Err: err}
				return
			}
		}

		var data utils.Response[Erc20TransfersResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount, utils.UserAgent)
			if err != nil {
				blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{Err: err}
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
				blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				blockTransactionWithContractTransfersChannel <- BlockTransactionWithContractTransfersResult{BlockTransactionWithContractTransfers: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return blockTransactionWithContractTransfersChannel
}

func (s *balanceServiceImpl) GetErc20TransfersForWalletAddressByPage(chainName chains.Chain, walletAddress string, queryParamOpts ...GetErc20TransfersForWalletAddressQueryParamOpts) (*utils.Response[Erc20TransfersResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/transfers_v2/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[Erc20TransfersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.ContractAddress != nil {
			params.Add("contract-address", fmt.Sprintf("%v", *opts.ContractAddress))
		}

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
		return &utils.Response[Erc20TransfersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[Erc20TransfersResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[Erc20TransfersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[Erc20TransfersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[Erc20TransfersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[Erc20TransfersResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[Erc20TransfersResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *balanceServiceImpl) GetTokenHoldersV2ForTokenAddress(chainName chains.Chain, tokenAddress string, queryParamOpts ...GetTokenHoldersV2ForTokenAddressQueryParamOpts) <-chan TokenHolderResult {
	tokenHolderChannel := make(chan TokenHolderResult)

	go func() {
		defer close(tokenHolderChannel)

		hasNext := true

		if !s.IskeyValid {
			tokenHolderChannel <- TokenHolderResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/tokens/%s/token_holders_v2/", chainName, tokenAddress)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			tokenHolderChannel <- TokenHolderResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.BlockHeight != nil {
				params.Add("block-height", fmt.Sprintf("%v", *opts.BlockHeight))
			}

			if opts.Date != nil {
				params.Add("date", fmt.Sprintf("%v", *opts.Date))
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
				tokenHolderChannel <- TokenHolderResult{Err: err}
				return
			}
		}

		var data utils.Response[TokenHoldersResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount, utils.UserAgent)
			if err != nil {
				tokenHolderChannel <- TokenHolderResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				tokenHolderChannel <- TokenHolderResult{Err: err}
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
				tokenHolderChannel <- TokenHolderResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				tokenHolderChannel <- TokenHolderResult{TokenHolder: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return tokenHolderChannel
}

func (s *balanceServiceImpl) GetTokenHoldersV2ForTokenAddressByPage(chainName chains.Chain, tokenAddress string, queryParamOpts ...GetTokenHoldersV2ForTokenAddressQueryParamOpts) (*utils.Response[TokenHoldersResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/tokens/%s/token_holders_v2/", chainName, tokenAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TokenHoldersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.BlockHeight != nil {
			params.Add("block-height", fmt.Sprintf("%v", *opts.BlockHeight))
		}

		if opts.Date != nil {
			params.Add("date", fmt.Sprintf("%v", *opts.Date))
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
		return &utils.Response[TokenHoldersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[TokenHoldersResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TokenHoldersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TokenHoldersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TokenHoldersResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TokenHoldersResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TokenHoldersResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *balanceServiceImpl) GetHistoricalTokenBalancesForWalletAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetHistoricalTokenBalancesForWalletAddressQueryParamOpts) (*utils.Response[HistoricalBalancesResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/historical_balances/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[HistoricalBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.Nft != nil {
			params.Add("nft", fmt.Sprintf("%v", *opts.Nft))
		}

		if opts.NoNftFetch != nil {
			params.Add("no-nft-fetch", fmt.Sprintf("%v", *opts.NoNftFetch))
		}

		if opts.NoSpam != nil {
			params.Add("no-spam", fmt.Sprintf("%v", *opts.NoSpam))
		}

		if opts.NoNftAssetMetadata != nil {
			params.Add("no-nft-asset-metadata", fmt.Sprintf("%v", *opts.NoNftAssetMetadata))
		}

		if opts.BlockHeight != nil {
			params.Add("block-height", fmt.Sprintf("%v", *opts.BlockHeight))
		}

		if opts.Date != nil {
			params.Add("date", fmt.Sprintf("%v", *opts.Date))
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
		return &utils.Response[HistoricalBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[HistoricalBalancesResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[HistoricalBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[HistoricalBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[HistoricalBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[HistoricalBalancesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[HistoricalBalancesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *balanceServiceImpl) GetNativeTokenBalance(chainName chains.Chain, walletAddress string, queryParamOpts ...GetNativeTokenBalanceQueryParamOpts) (*utils.Response[TokenBalanceNativeResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/balances_native/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TokenBalanceNativeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.BlockHeight != nil {
			params.Add("block-height", fmt.Sprintf("%v", *opts.BlockHeight))
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
		return &utils.Response[TokenBalanceNativeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[TokenBalanceNativeResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TokenBalanceNativeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TokenBalanceNativeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TokenBalanceNativeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TokenBalanceNativeResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TokenBalanceNativeResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
