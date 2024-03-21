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

type ChainCollectionResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []ChainCollectionItem `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type ChainCollectionItem struct {
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// Denotes whether the token is suspected spam. Supports `eth-mainnet` and `matic-mainnet`.
	IsSpam              *bool      `json:"is_spam,omitempty"`
	TokenTotalSupply    *int64     `json:"token_total_supply,omitempty"`
	CachedMetadataCount *int64     `json:"cached_metadata_count,omitempty"`
	CachedAssetCount    *int64     `json:"cached_asset_count,omitempty"`
	LastScrapedAt       *time.Time `json:"last_scraped_at,omitempty"`
}
type NftAddressBalanceNftResponse struct {
	// The requested address.
	Address string `json:"address"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []NftTokenContractBalanceItem `json:"items"`
}
type NftTokenContractBalanceItem struct {
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *[]string `json:"supports_erc,omitempty"`
	// Denotes whether the token is suspected spam. Supports `eth-mainnet` and `matic-mainnet`.
	IsSpam           *bool      `json:"is_spam,omitempty"`
	LastTransferedAt *time.Time `json:"last_transfered_at,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance    *utils.BigInt `json:"balance,omitempty"`
	Balance24h *string       `json:"balance_24h,omitempty"`
	Type       *string       `json:"type,omitempty"`
	// The current floor price converted to fiat in `quote-currency`. The floor price is determined by the last minimum sale price within the last 30 days across all the supported markets where the collection is sold on.
	FloorPriceQuote *float64 `json:"floor_price_quote,omitempty"`
	// A prettier version of the floor price quote for rendering purposes.
	PrettyFloorPriceQuote *string `json:"pretty_floor_price_quote,omitempty"`
	// The current floor price in native currency. The floor price is determined by the last minimum sale price within the last 30 days across all the supported markets where the collection is sold on.
	FloorPriceNativeQuote *float64   `json:"floor_price_native_quote,omitempty"`
	NftData               *[]NftData `json:"nft_data,omitempty"`
}
type NftData struct {
	// The token's id.
	TokenId  *utils.BigInt `json:"token_id,omitempty"`
	TokenUrl *string       `json:"token_url,omitempty"`
	// The original minter.
	OriginalOwner *string `json:"original_owner,omitempty"`
	// The current holder of this NFT.
	CurrentOwner *string          `json:"current_owner,omitempty"`
	ExternalData *NftExternalData `json:"external_data,omitempty"`
	// If `true`, the asset data is available from the Covalent CDN.
	AssetCached *bool `json:"asset_cached,omitempty"`
	// If `true`, the image data is available from the Covalent CDN.
	ImageCached *bool `json:"image_cached,omitempty"`
}
type NftExternalData struct {
	Name               *string                                 `json:"name,omitempty"`
	Description        *string                                 `json:"description,omitempty"`
	AssetUrl           *string                                 `json:"asset_url,omitempty"`
	AssetFileExtension *string                                 `json:"asset_file_extension,omitempty"`
	AssetMimeType      *string                                 `json:"asset_mime_type,omitempty"`
	AssetSizeBytes     *string                                 `json:"asset_size_bytes,omitempty"`
	Image              *string                                 `json:"image,omitempty"`
	Image256           *string                                 `json:"image_256,omitempty"`
	Image512           *string                                 `json:"image_512,omitempty"`
	Image1024          *string                                 `json:"image_1024,omitempty"`
	AnimationUrl       *string                                 `json:"animation_url,omitempty"`
	ExternalUrl        *string                                 `json:"external_url,omitempty"`
	Attributes         *[]genericmodels.NftCollectionAttribute `json:"attributes,omitempty"`
}
type NftMetadataResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []NftTokenContract `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type NftTokenContract struct {
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// Denotes whether the token is suspected spam. Supports `eth-mainnet` and `matic-mainnet`.
	IsSpam  *bool    `json:"is_spam,omitempty"`
	Type    *string  `json:"type,omitempty"`
	NftData *NftData `json:"nft_data,omitempty"`
}
type NftTransactionsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []NftTransaction `json:"items"`
}
type NftTransaction struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc     *[]string             `json:"supports_erc,omitempty"`
	NftTransactions *[]NftTransactionItem `json:"nft_transactions,omitempty"`
	// Denotes whether the token is suspected spam. Supports `eth-mainnet` and `matic-mainnet`.
	IsSpam *bool `json:"is_spam,omitempty"`
}
type NftTransactionItem struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The height of the block.
	BlockHeight *int `json:"block_height,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	// The offset is the position of the tx in the block.
	TxOffset *int `json:"tx_offset,omitempty"`
	// Whether or not transaction is successful.
	Successful *bool `json:"successful,omitempty"`
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
	GasOffered       *int64  `json:"gas_offered,omitempty"`
	// The gas spent for this tx.
	GasSpent *int64 `json:"gas_spent,omitempty"`
	// The gas price at the time of this tx.
	GasPrice *int64 `json:"gas_price,omitempty"`
	// The total transaction fees (gas_price * gas_spent) paid for this tx, denoted in wei.
	FeesPaid *utils.BigInt `json:"fees_paid,omitempty"`
	// The gas spent in `quote-currency` denomination.
	GasQuote *float64 `json:"gas_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyGasQuote *string  `json:"pretty_gas_quote,omitempty"`
	GasQuoteRate   *float64 `json:"gas_quote_rate,omitempty"`
	// The log events.
	LogEvents *[]genericmodels.LogEvent `json:"log_events,omitempty"`
}
type NftCollectionTraitsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []NftTrait `json:"items"`
}
type NftTrait struct {
	Name *string `json:"name,omitempty"`
}
type NftCollectionAttributesForTraitResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []NftSummaryAttribute `json:"items"`
}
type NftSummaryAttribute struct {
	TraitType    *string         `json:"trait_type,omitempty"`
	Values       *[]NftAttribute `json:"values,omitempty"`
	UniqueValues *int            `json:"unique_values,omitempty"`
}
type NftAttribute struct {
	Value *string `json:"value,omitempty"`
	Count *int    `json:"count,omitempty"`
}
type NftCollectionTraitSummaryResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []NftTraitSummary `json:"items"`
}
type NftTraitSummary struct {
	// Trait name
	Name *string `json:"name,omitempty"`
	// Type of the value of the trait.
	ValueType *string `json:"value_type,omitempty"`
	// Populated for `numeric` traits.
	ValueNumeric *NftTraitNumeric `json:"value_numeric,omitempty"`
	// Populated for `string` traits.
	ValueString *NftTraitString        `json:"value_string,omitempty"`
	Attributes  *[]NftSummaryAttribute `json:"attributes,omitempty"`
}
type NftTraitNumeric struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}
type NftTraitString struct {
	// String value
	Value *string `json:"value,omitempty"`
	// Number of distinct tokens that have this trait value.
	TokenCount *int `json:"token_count,omitempty"`
	// Percentage of tokens in the collection that have this trait.
	TraitPercentage *float64 `json:"trait_percentage,omitempty"`
}
type NftOwnershipForCollectionResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested collection.
	Collection string `json:"collection"`
	// Denotes whether the token is suspected spam. Supports `eth-mainnet` and `matic-mainnet`.
	IsSpam bool `json:"is_spam"`
	// List of response items.
	Items []NftOwnershipForCollectionItem `json:"items"`
}
type NftOwnershipForCollectionItem struct {
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The token's id.
	TokenId *utils.BigInt `json:"token_id,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc      *[]string  `json:"supports_erc,omitempty"`
	LastTransferedAt *time.Time `json:"last_transfered_at,omitempty"`
	// Nft balance.
	Balance    *utils.BigInt `json:"balance,omitempty"`
	Balance24h *string       `json:"balance_24h,omitempty"`
	Type       *string       `json:"type,omitempty"`
	NftData    *NftData      `json:"nft_data,omitempty"`
}
type NftMarketSaleCountResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// List of response items.
	Items []MarketSaleCountItem `json:"items"`
}
type MarketSaleCountItem struct {
	// The timestamp of the date of sale.
	Date *time.Time `json:"date,omitempty"`
	// The total amount of sales for the current day.
	SaleCount *int `json:"sale_count,omitempty"`
}
type NftMarketVolumeResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// List of response items.
	Items []MarketVolumeItem `json:"items"`
}
type MarketVolumeItem struct {
	// The timestamp of the date of sale.
	Date *time.Time `json:"date,omitempty"`
	// The ticker symbol for the native currency.
	NativeTickerSymbol *string `json:"native_ticker_symbol,omitempty"`
	// The contract name of the native currency.
	NativeName *string `json:"native_name,omitempty"`
	// The current volume converted to fiat in `quote-currency`.
	VolumeQuote *float64 `json:"volume_quote,omitempty"`
	// The current volume in native currency.
	VolumeNativeQuote *float64 `json:"volume_native_quote,omitempty"`
	// A prettier version of the volume quote for rendering purposes.
	PrettyVolumeQuote *string `json:"pretty_volume_quote,omitempty"`
}
type NftMarketFloorPriceResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// List of response items.
	Items []MarketFloorPriceItem `json:"items"`
}
type MarketFloorPriceItem struct {
	// The timestamp of the date of sale.
	Date *time.Time `json:"date,omitempty"`
	// The ticker symbol for the native currency.
	NativeTickerSymbol *string `json:"native_ticker_symbol,omitempty"`
	// The contract name of the native currency.
	NativeName *string `json:"native_name,omitempty"`
	// The current floor price in native currency.
	FloorPriceNativeQuote *float64 `json:"floor_price_native_quote,omitempty"`
	// The current floor price converted to fiat in `quote-currency`.
	FloorPriceQuote *float64 `json:"floor_price_quote,omitempty"`
	// A prettier version of the floor price quote for rendering purposes.
	PrettyFloorPriceQuote *string `json:"pretty_floor_price_quote,omitempty"`
}
type ChainCollectionItemResult struct {
	ChainCollectionItem ChainCollectionItem
	Err                 error
}

type NftTokenContractResult struct {
	NftTokenContract NftTokenContract
	Err              error
}

type GetChainCollectionsQueryParamOpts struct {
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
	// If `true`, the suspected spam tokens are removed. Supports `eth-mainnet` and `matic-mainnet`.
	NoSpam *bool `json:"noSpam,omitempty"`
}
type GetNftsForAddressQueryParamOpts struct {
	// If `true`, the suspected spam tokens are removed. Supports `eth-mainnet` and `matic-mainnet`.
	NoSpam *bool `json:"noSpam,omitempty"`
	// If `true`, the response shape is limited to a list of collections and token ids, omitting metadata and asset information. Helpful for faster response times and wallets holding a large number of NFTs.
	NoNftAssetMetadata *bool `json:"noNftAssetMetadata,omitempty"`
	// By default, this endpoint only works on chains where we've cached the assets and the metadata. When set to `true`, the API will fetch metadata from upstream servers even if it's not cached - the downside being that the upstream server can block or rate limit the call and therefore resulting in time outs or slow response times on the Covalent side.
	WithUncached *bool `json:"withUncached,omitempty"`
}
type GetTokenIdsForContractWithMetadataQueryParamOpts struct {
	// Omit metadata.
	NoMetadata *bool `json:"noMetadata,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
	// Filters NFTs based on a specific trait. If this filter is used, the API will return all NFTs with the specified trait. Accepts comma-separated values, is case-sensitive, and requires proper URL encoding.
	TraitsFilter *string `json:"traitsFilter,omitempty"`
	// Filters NFTs based on a specific trait value. If this filter is used, the API will return all NFTs with the specified trait value. If used with "traits-filter", only NFTs matching both filters will be returned. Accepts comma-separated values, is case-sensitive, and requires proper URL encoding.
	ValuesFilter *string `json:"valuesFilter,omitempty"`
	// By default, this endpoint only works on chains where we've cached the assets and the metadata. When set to `true`, the API will fetch metadata from upstream servers even if it's not cached - the downside being that the upstream server can block or rate limit the call and therefore resulting in time outs or slow response times on the Covalent side.
	WithUncached *bool `json:"withUncached,omitempty"`
}
type GetNftMetadataForGivenTokenIdForContractQueryParamOpts struct {
	// Omit metadata.
	NoMetadata *bool `json:"noMetadata,omitempty"`
	// By default, this endpoint only works on chains where we've cached the assets and the metadata. When set to `true`, the API will fetch metadata from upstream servers even if it's not cached - the downside being that the upstream server can block or rate limit the call and therefore resulting in time outs or slow response times on the Covalent side.
	WithUncached *bool `json:"withUncached,omitempty"`
}
type GetNftTransactionsForContractTokenIdQueryParamOpts struct {
	// If `true`, the suspected spam tokens are removed. Supports `eth-mainnet` and `matic-mainnet`.
	NoSpam *bool `json:"noSpam,omitempty"`
}
type CheckOwnershipInNftQueryParamOpts struct {
	// Filters NFTs based on a specific trait. If this filter is used, the API will return all NFTs with the specified trait. Must be used with "values-filter", is case-sensitive, and requires proper URL encoding.
	TraitsFilter *string `json:"traitsFilter,omitempty"`
	// Filters NFTs based on a specific trait value. If this filter is used, the API will return all NFTs with the specified trait value. Must be used with "traits-filter", is case-sensitive, and requires proper URL encoding.
	ValuesFilter *string `json:"valuesFilter,omitempty"`
}
type GetNftMarketSaleCountQueryParamOpts struct {
	// The number of days to return data for. Request up 365 days. Defaults to 30 days.
	Days *int `json:"days,omitempty"`
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
}
type GetNftMarketVolumeQueryParamOpts struct {
	// The number of days to return data for. Request up 365 days. Defaults to 30 days.
	Days *int `json:"days,omitempty"`
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
}
type GetNftMarketFloorPriceQueryParamOpts struct {
	// The number of days to return data for. Request up 365 days. Defaults to 30 days.
	Days *int `json:"days,omitempty"`
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
}

func NewNftServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) NftService {

	return &nftServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type NftService interface {

	// Commonly used to fetch the list of NFT collections with downloaded and cached off chain data like token metadata and asset files.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	GetChainCollections(chainName chains.Chain, queryParamOpts ...GetChainCollectionsQueryParamOpts) <-chan ChainCollectionItemResult

	// Commonly used to fetch the list of NFT collections with downloaded and cached off chain data like token metadata and asset files.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	GetChainCollectionsByPage(chainName chains.Chain, queryParamOpts ...GetChainCollectionsQueryParamOpts) (*utils.Response[ChainCollectionResponse], error)

	// Commonly used to render the NFTs (including ERC721 and ERC1155) held by an address.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetNftsForAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetNftsForAddressQueryParamOpts) (*utils.Response[NftAddressBalanceNftResponse], error)

	// Commonly used to get NFT token IDs with metadata from a collection. Useful for building NFT card displays.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTokenIdsForContractWithMetadata(chainName chains.Chain, contractAddress string, queryParamOpts ...GetTokenIdsForContractWithMetadataQueryParamOpts) <-chan NftTokenContractResult

	// Commonly used to get NFT token IDs with metadata from a collection. Useful for building NFT card displays.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTokenIdsForContractWithMetadataByPage(chainName chains.Chain, contractAddress string, queryParamOpts ...GetTokenIdsForContractWithMetadataQueryParamOpts) (*utils.Response[NftMetadataResponse], error)

	// Commonly used to get a single NFT metadata by token ID from a collection. Useful for building NFT card displays.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// tokenId: The requested token ID.. Type: string
	GetNftMetadataForGivenTokenIdForContract(chainName chains.Chain, contractAddress string, tokenId string, queryParamOpts ...GetNftMetadataForGivenTokenIdForContractQueryParamOpts) (*utils.Response[NftMetadataResponse], error)

	// Commonly used to get all transactions of an NFT token. Useful for building a transaction history table or price chart.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// tokenId: The requested token ID.. Type: string
	GetNftTransactionsForContractTokenId(chainName chains.Chain, contractAddress string, tokenId string, queryParamOpts ...GetNftTransactionsForContractTokenIdQueryParamOpts) (*utils.Response[NftTransactionsResponse], error)

	// Commonly used to fetch and render the traits of a collection as seen in rarity calculators.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// collectionContract: The requested collection address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTraitsForCollection(chainName chains.Chain, collectionContract string) (*utils.Response[NftCollectionTraitsResponse], error)

	// Commonly used to get the count of unique values for traits within an NFT collection.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// collectionContract: The requested collection address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// trait: The requested trait.. Type: string
	GetAttributesForTraitInCollection(chainName chains.Chain, collectionContract string, trait string) (*utils.Response[NftCollectionAttributesForTraitResponse], error)

	// Commonly used to calculate rarity scores for a collection based on its traits.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// collectionContract: The requested collection address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetCollectionTraitsSummary(chainName chains.Chain, collectionContract string) (*utils.Response[NftCollectionTraitSummaryResponse], error)

	// Commonly used to verify ownership of NFTs (including ERC-721 and ERC-1155) within a collection.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// collectionContract: The requested collection address.. Type: string
	CheckOwnershipInNft(chainName chains.Chain, walletAddress string, collectionContract string, queryParamOpts ...CheckOwnershipInNftQueryParamOpts) (*utils.Response[NftOwnershipForCollectionResponse], error)

	// Commonly used to verify ownership of a specific token (ERC-721 or ERC-1155) within a collection.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// collectionContract: The requested collection address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// tokenId: The requested token ID.. Type: string
	CheckOwnershipInNftForSpecificTokenId(chainName chains.Chain, walletAddress string, collectionContract string, tokenId string) (*utils.Response[NftOwnershipForCollectionResponse], error)

	// Commonly used to build a time-series chart of the sales count of an NFT collection.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetNftMarketSaleCount(chainName chains.Chain, contractAddress string, queryParamOpts ...GetNftMarketSaleCountQueryParamOpts) (*utils.Response[NftMarketSaleCountResponse], error)

	// Commonly used to build a time-series chart of the transaction volume of an NFT collection.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetNftMarketVolume(chainName chains.Chain, contractAddress string, queryParamOpts ...GetNftMarketVolumeQueryParamOpts) (*utils.Response[NftMarketVolumeResponse], error)

	// Commonly used to render a price floor chart for an NFT collection.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// contractAddress: The requested contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetNftMarketFloorPrice(chainName chains.Chain, contractAddress string, queryParamOpts ...GetNftMarketFloorPriceQueryParamOpts) (*utils.Response[NftMarketFloorPriceResponse], error)
}

type nftServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *nftServiceImpl) GetChainCollections(chainName chains.Chain, queryParamOpts ...GetChainCollectionsQueryParamOpts) <-chan ChainCollectionItemResult {
	chainCollectionItemChannel := make(chan ChainCollectionItemResult)

	go func() {
		defer close(chainCollectionItemChannel)

		hasNext := true

		if !s.IskeyValid {
			chainCollectionItemChannel <- ChainCollectionItemResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/collections/", chainName)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			chainCollectionItemChannel <- ChainCollectionItemResult{Err: err}
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

			if opts.NoSpam != nil {
				params.Add("no-spam", fmt.Sprintf("%v", *opts.NoSpam))
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
				chainCollectionItemChannel <- ChainCollectionItemResult{Err: err}
				return
			}
		}

		var data utils.Response[ChainCollectionResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount, utils.UserAgent)
			if err != nil {
				chainCollectionItemChannel <- ChainCollectionItemResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				chainCollectionItemChannel <- ChainCollectionItemResult{Err: err}
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
				chainCollectionItemChannel <- ChainCollectionItemResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				chainCollectionItemChannel <- ChainCollectionItemResult{ChainCollectionItem: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return chainCollectionItemChannel
}

func (s *nftServiceImpl) GetChainCollectionsByPage(chainName chains.Chain, queryParamOpts ...GetChainCollectionsQueryParamOpts) (*utils.Response[ChainCollectionResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/collections/", chainName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[ChainCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.NoSpam != nil {
			params.Add("no-spam", fmt.Sprintf("%v", *opts.NoSpam))
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
		return &utils.Response[ChainCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[ChainCollectionResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[ChainCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ChainCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ChainCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[ChainCollectionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[ChainCollectionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetNftsForAddress(chainName chains.Chain, walletAddress string, queryParamOpts ...GetNftsForAddressQueryParamOpts) (*utils.Response[NftAddressBalanceNftResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/balances_nft/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftAddressBalanceNftResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.NoSpam != nil {
			params.Add("no-spam", fmt.Sprintf("%v", *opts.NoSpam))
		}

		if opts.NoNftAssetMetadata != nil {
			params.Add("no-nft-asset-metadata", fmt.Sprintf("%v", *opts.NoNftAssetMetadata))
		}

		if opts.WithUncached != nil {
			params.Add("with-uncached", fmt.Sprintf("%v", *opts.WithUncached))
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
		return &utils.Response[NftAddressBalanceNftResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftAddressBalanceNftResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftAddressBalanceNftResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftAddressBalanceNftResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftAddressBalanceNftResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftAddressBalanceNftResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftAddressBalanceNftResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetTokenIdsForContractWithMetadata(chainName chains.Chain, contractAddress string, queryParamOpts ...GetTokenIdsForContractWithMetadataQueryParamOpts) <-chan NftTokenContractResult {
	nftTokenContractChannel := make(chan NftTokenContractResult)

	go func() {
		defer close(nftTokenContractChannel)

		hasNext := true

		if !s.IskeyValid {
			nftTokenContractChannel <- NftTokenContractResult{Err: fmt.Errorf(`An error occurred 401: ` + utils.InvalidAPIKeyMessage)}
			return
		}

		apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/%s/metadata/", chainName, contractAddress)

		// Parse the formatted URL
		parsedURL, err := url.Parse(apiURL)
		if err != nil {
			nftTokenContractChannel <- NftTokenContractResult{Err: err}
			return
		}

		params := url.Values{}
		if len(queryParamOpts) > 0 {
			opts := queryParamOpts[0]

			if opts.NoMetadata != nil {
				params.Add("no-metadata", fmt.Sprintf("%v", *opts.NoMetadata))
			}

			if opts.PageSize != nil {
				params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
			}

			if opts.PageNumber != nil {
				params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
			}

			if opts.TraitsFilter != nil {
				params.Add("traits-filter", fmt.Sprintf("%v", *opts.TraitsFilter))
			}

			if opts.ValuesFilter != nil {
				params.Add("values-filter", fmt.Sprintf("%v", *opts.ValuesFilter))
			}

			if opts.WithUncached != nil {
				params.Add("with-uncached", fmt.Sprintf("%v", *opts.WithUncached))
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
				nftTokenContractChannel <- NftTokenContractResult{Err: err}
				return
			}
		}

		var data utils.Response[NftMetadataResponse]
		for hasNext {

			res, err := utils.PaginateEndpoint(apiURL, s.APIKey, params, page, s.Debug, s.ThreadCount, utils.UserAgent)
			if err != nil {
				nftTokenContractChannel <- NftTokenContractResult{Err: err}
				hasNext = false
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				nftTokenContractChannel <- NftTokenContractResult{Err: err}
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
				nftTokenContractChannel <- NftTokenContractResult{Err: errors.New("An error occurred " + strconv.Itoa(*data.ErrorCode) + ": " + errorMessage)}
				return
			}

			for _, item := range data.Data.Items {
				nftTokenContractChannel <- NftTokenContractResult{NftTokenContract: item, Err: err}
			}

			hasNext = *data.Data.Pagination.HasMore
			page++
		}

	}()
	return nftTokenContractChannel
}

func (s *nftServiceImpl) GetTokenIdsForContractWithMetadataByPage(chainName chains.Chain, contractAddress string, queryParamOpts ...GetTokenIdsForContractWithMetadataQueryParamOpts) (*utils.Response[NftMetadataResponse], error) {
	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/%s/metadata/", chainName, contractAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.NoMetadata != nil {
			params.Add("no-metadata", fmt.Sprintf("%v", *opts.NoMetadata))
		}

		if opts.PageSize != nil {
			params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
		}

		if opts.PageNumber != nil {
			params.Add("page-number", fmt.Sprintf("%v", *opts.PageNumber))
		}

		if opts.TraitsFilter != nil {
			params.Add("traits-filter", fmt.Sprintf("%v", *opts.TraitsFilter))
		}

		if opts.ValuesFilter != nil {
			params.Add("values-filter", fmt.Sprintf("%v", *opts.ValuesFilter))
		}

		if opts.WithUncached != nil {
			params.Add("with-uncached", fmt.Sprintf("%v", *opts.WithUncached))
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
		return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftMetadataResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftMetadataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftMetadataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetNftMetadataForGivenTokenIdForContract(chainName chains.Chain, contractAddress string, tokenId string, queryParamOpts ...GetNftMetadataForGivenTokenIdForContractQueryParamOpts) (*utils.Response[NftMetadataResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/%s/metadata/%s/", chainName, contractAddress, tokenId)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.NoMetadata != nil {
			params.Add("no-metadata", fmt.Sprintf("%v", *opts.NoMetadata))
		}

		if opts.WithUncached != nil {
			params.Add("with-uncached", fmt.Sprintf("%v", *opts.WithUncached))
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
		return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftMetadataResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMetadataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftMetadataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftMetadataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetNftTransactionsForContractTokenId(chainName chains.Chain, contractAddress string, tokenId string, queryParamOpts ...GetNftTransactionsForContractTokenIdQueryParamOpts) (*utils.Response[NftTransactionsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/tokens/%s/nft_transactions/%s/", chainName, contractAddress, tokenId)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.NoSpam != nil {
			params.Add("no-spam", fmt.Sprintf("%v", *opts.NoSpam))
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
		return &utils.Response[NftTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftTransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetTraitsForCollection(chainName chains.Chain, collectionContract string) (*utils.Response[NftCollectionTraitsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/%s/traits/", chainName, collectionContract)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftCollectionTraitsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NftCollectionTraitsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftCollectionTraitsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftCollectionTraitsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftCollectionTraitsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftCollectionTraitsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftCollectionTraitsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftCollectionTraitsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetAttributesForTraitInCollection(chainName chains.Chain, collectionContract string, trait string) (*utils.Response[NftCollectionAttributesForTraitResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/%s/traits/%s/attributes/", chainName, collectionContract, trait)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftCollectionAttributesForTraitResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftCollectionAttributesForTraitResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetCollectionTraitsSummary(chainName chains.Chain, collectionContract string) (*utils.Response[NftCollectionTraitSummaryResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/%s/traits_summary/", chainName, collectionContract)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftCollectionTraitSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NftCollectionTraitSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftCollectionTraitSummaryResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftCollectionTraitSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftCollectionTraitSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftCollectionTraitSummaryResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftCollectionTraitSummaryResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftCollectionTraitSummaryResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) CheckOwnershipInNft(chainName chains.Chain, walletAddress string, collectionContract string, queryParamOpts ...CheckOwnershipInNftQueryParamOpts) (*utils.Response[NftOwnershipForCollectionResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/collection/%s/", chainName, walletAddress, collectionContract)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.TraitsFilter != nil {
			params.Add("traits-filter", fmt.Sprintf("%v", *opts.TraitsFilter))
		}

		if opts.ValuesFilter != nil {
			params.Add("values-filter", fmt.Sprintf("%v", *opts.ValuesFilter))
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
		return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftOwnershipForCollectionResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftOwnershipForCollectionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftOwnershipForCollectionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) CheckOwnershipInNftForSpecificTokenId(chainName chains.Chain, walletAddress string, collectionContract string, tokenId string) (*utils.Response[NftOwnershipForCollectionResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/address/%s/collection/%s/token/%s/", chainName, walletAddress, collectionContract, tokenId)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftOwnershipForCollectionResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftOwnershipForCollectionResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftOwnershipForCollectionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftOwnershipForCollectionResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetNftMarketSaleCount(chainName chains.Chain, contractAddress string, queryParamOpts ...GetNftMarketSaleCountQueryParamOpts) (*utils.Response[NftMarketSaleCountResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft_market/%s/sale_count/", chainName, contractAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftMarketSaleCountResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.Days != nil {
			params.Add("days", fmt.Sprintf("%v", *opts.Days))
		}

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
		return &utils.Response[NftMarketSaleCountResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftMarketSaleCountResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftMarketSaleCountResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMarketSaleCountResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMarketSaleCountResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftMarketSaleCountResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftMarketSaleCountResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetNftMarketVolume(chainName chains.Chain, contractAddress string, queryParamOpts ...GetNftMarketVolumeQueryParamOpts) (*utils.Response[NftMarketVolumeResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft_market/%s/volume/", chainName, contractAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftMarketVolumeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.Days != nil {
			params.Add("days", fmt.Sprintf("%v", *opts.Days))
		}

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
		return &utils.Response[NftMarketVolumeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftMarketVolumeResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftMarketVolumeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMarketVolumeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMarketVolumeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftMarketVolumeResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftMarketVolumeResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *nftServiceImpl) GetNftMarketFloorPrice(chainName chains.Chain, contractAddress string, queryParamOpts ...GetNftMarketFloorPriceQueryParamOpts) (*utils.Response[NftMarketFloorPriceResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft_market/%s/floor_price/", chainName, contractAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftMarketFloorPriceResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.Days != nil {
			params.Add("days", fmt.Sprintf("%v", *opts.Days))
		}

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
		return &utils.Response[NftMarketFloorPriceResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftMarketFloorPriceResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftMarketFloorPriceResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMarketFloorPriceResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftMarketFloorPriceResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftMarketFloorPriceResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftMarketFloorPriceResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
