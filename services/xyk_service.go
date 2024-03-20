package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/genericmodels"
	"github.com/covalenthq/covalent-api-sdk-go/quotes"
	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

type PoolResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []Pool `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type Pool struct {
	// The pair address.
	Exchange     *string `json:"exchange,omitempty"`
	SwapCount24h *int64  `json:"swap_count_24h,omitempty"`
	// The total liquidity converted to fiat in `quote-currency`.
	TotalLiquidityQuote *float64 `json:"total_liquidity_quote,omitempty"`
	Volume24hQuote      *float64 `json:"volume_24h_quote,omitempty"`
	Fee24hQuote         *float64 `json:"fee_24h_quote,omitempty"`
	// Total supply of this pool token.
	TotalSupply *utils.BigInt `json:"total_supply,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// A prettier version of the total liquidity quote for rendering purposes.
	PrettyTotalLiquidityQuote *string `json:"pretty_total_liquidity_quote,omitempty"`
	// A prettier version of the volume 24h quote for rendering purposes.
	PrettyVolume24hQuote *string `json:"pretty_volume_24h_quote,omitempty"`
	// A prettier version of the fee 24h quote for rendering purposes.
	PrettyFee24hQuote *string `json:"pretty_fee_24h_quote,omitempty"`
	// A prettier version of the volume 7d quote for rendering purposes.
	PrettyVolume7dQuote *string `json:"pretty_volume_7d_quote,omitempty"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName *string `json:"chain_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The name of the DEX, eg: `uniswap_v2`.
	DexName       *string  `json:"dex_name,omitempty"`
	Volume7dQuote *float64 `json:"volume_7d_quote,omitempty"`
	AnnualizedFee *float64 `json:"annualized_fee,omitempty"`
	Token0        *Token   `json:"token_0,omitempty"`
	Token1        *Token   `json:"token_1,omitempty"`
}
type Token struct {
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	VolumeIn24h  *string `json:"volume_in_24h,omitempty"`
	VolumeOut24h *string `json:"volume_out_24h,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	Reserve   *string  `json:"reserve,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int    `json:"contract_decimals,omitempty"`
	VolumeIn7d       *string `json:"volume_in_7d,omitempty"`
	VolumeOut7d      *string `json:"volume_out_7d,omitempty"`
}
type PoolToDexResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []PoolToDexItem `json:"items"`
}
type PoolToDexItem struct {
	SupportedDex
	// The dex logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
}
type SupportedDex struct {
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName *string `json:"chain_name,omitempty"`
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// A display-friendly name for the dex.
	DisplayName *string `json:"display_name,omitempty"`
	// The dex logo URL.
	LogoUrl                 *string   `json:"logo_url,omitempty"`
	FactoryContractAddress  *string   `json:"factory_contract_address,omitempty"`
	RouterContractAddresses *[]string `json:"router_contract_addresses,omitempty"`
	SwapFee                 *float64  `json:"swap_fee,omitempty"`
}
type PoolByAddressResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []PoolWithTimeseries `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type PoolWithTimeseries struct {
	// The pair address.
	Exchange *string `json:"exchange,omitempty"`
	// A list of explorers for this address.
	Explorers    *[]genericmodels.Explorer `json:"explorers,omitempty"`
	SwapCount24h *int64                    `json:"swap_count_24h,omitempty"`
	// The total liquidity converted to fiat in `quote-currency`.
	TotalLiquidityQuote *float64 `json:"total_liquidity_quote,omitempty"`
	Volume24hQuote      *float64 `json:"volume_24h_quote,omitempty"`
	Fee24hQuote         *float64 `json:"fee_24h_quote,omitempty"`
	// Total supply of this pool token.
	TotalSupply *utils.BigInt `json:"total_supply,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The name of the DEX, eg: `uniswap_v2`.
	DexName       *string  `json:"dex_name,omitempty"`
	Volume7dQuote *float64 `json:"volume_7d_quote,omitempty"`
	AnnualizedFee *float64 `json:"annualized_fee,omitempty"`
	// A prettier version of the total liquidity quote for rendering purposes.
	PrettyTotalLiquidityQuote *string `json:"pretty_total_liquidity_quote,omitempty"`
	// A prettier version of the volume 24h quote for rendering purposes.
	PrettyVolume24hQuote *string `json:"pretty_volume_24h_quote,omitempty"`
	// A prettier version of the fee 24h quote for rendering purposes.
	PrettyFee24hQuote *string `json:"pretty_fee_24h_quote,omitempty"`
	// A prettier version of the volume 7d quote for rendering purposes.
	PrettyVolume7dQuote    *string                `json:"pretty_volume_7d_quote,omitempty"`
	Token0                 *Token                 `json:"token_0,omitempty"`
	Token1                 *Token                 `json:"token_1,omitempty"`
	Token0ReserveQuote     *float64               `json:"token_0_reserve_quote,omitempty"`
	Token1ReserveQuote     *float64               `json:"token_1_reserve_quote,omitempty"`
	VolumeTimeseries7d     *[]VolumeTimeseries    `json:"volume_timeseries_7d,omitempty"`
	VolumeTimeseries30d    *[]VolumeTimeseries    `json:"volume_timeseries_30d,omitempty"`
	LiquidityTimeseries7d  *[]LiquidityTimeseries `json:"liquidity_timeseries_7d,omitempty"`
	LiquidityTimeseries30d *[]LiquidityTimeseries `json:"liquidity_timeseries_30d,omitempty"`
	PriceTimeseries7d      *[]PriceTimeseries     `json:"price_timeseries_7d,omitempty"`
	PriceTimeseries30d     *[]PriceTimeseries     `json:"price_timeseries_30d,omitempty"`
}
type VolumeTimeseries struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string    `json:"chain_id,omitempty"`
	Dt      *time.Time `json:"dt,omitempty"`
	// The pair address.
	Exchange      *string  `json:"exchange,omitempty"`
	SumAmount0in  *string  `json:"sum_amount0in,omitempty"`
	SumAmount0out *string  `json:"sum_amount0out,omitempty"`
	SumAmount1in  *string  `json:"sum_amount1in,omitempty"`
	SumAmount1out *string  `json:"sum_amount1out,omitempty"`
	VolumeQuote   *float64 `json:"volume_quote,omitempty"`
	// A prettier version of the volume quote for rendering purposes.
	PrettyVolumeQuote *string  `json:"pretty_volume_quote,omitempty"`
	Token0QuoteRate   *float64 `json:"token_0_quote_rate,omitempty"`
	Token1QuoteRate   *float64 `json:"token_1_quote_rate,omitempty"`
	SwapCount24       *int64   `json:"swap_count_24,omitempty"`
}
type LiquidityTimeseries struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string    `json:"chain_id,omitempty"`
	Dt      *time.Time `json:"dt,omitempty"`
	// The pair address.
	Exchange       *string  `json:"exchange,omitempty"`
	R0C            *string  `json:"r0_c,omitempty"`
	R1C            *string  `json:"r1_c,omitempty"`
	LiquidityQuote *float64 `json:"liquidity_quote,omitempty"`
	// A prettier version of the liquidity quote for rendering purposes.
	PrettyLiquidityQuote *string  `json:"pretty_liquidity_quote,omitempty"`
	Token0QuoteRate      *float64 `json:"token_0_quote_rate,omitempty"`
	Token1QuoteRate      *float64 `json:"token_1_quote_rate,omitempty"`
}
type PriceTimeseries struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string    `json:"chain_id,omitempty"`
	Dt      *time.Time `json:"dt,omitempty"`
	// The pair address.
	Exchange              *string  `json:"exchange,omitempty"`
	PriceOfToken0InToken1 *float64 `json:"price_of_token0_in_token1,omitempty"`
	// A prettier version of the price token0 for rendering purposes.
	PrettyPriceOfToken0InToken1      *string  `json:"pretty_price_of_token0_in_token1,omitempty"`
	PriceOfToken0InToken1Description *string  `json:"price_of_token0_in_token1_description,omitempty"`
	PriceOfToken1InToken0            *float64 `json:"price_of_token1_in_token0,omitempty"`
	// A prettier version of the price token1 for rendering purposes.
	PrettyPriceOfToken1InToken0      *string `json:"pretty_price_of_token1_in_token0,omitempty"`
	PriceOfToken1InToken0Description *string `json:"price_of_token1_in_token0_description,omitempty"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency                *string  `json:"quote_currency,omitempty"`
	PriceOfToken0InQuoteCurrency *float64 `json:"price_of_token0_in_quote_currency,omitempty"`
	PriceOfToken1InQuoteCurrency *float64 `json:"price_of_token1_in_quote_currency,omitempty"`
}
type PoolsDexDataResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested address.
	Address string `json:"address"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// List of response items.
	Items []PoolsDexDataItem `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type PoolsDexDataItem struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The pair address.
	Exchange *string `json:"exchange,omitempty"`
	// The combined ticker symbol of token0 and token1 separated with a hypen.
	ExchangeTickerSymbol *string `json:"exchange_ticker_symbol,omitempty"`
	// The dex logo URL for the pair address.
	ExchangeLogoUrl *string `json:"exchange_logo_url,omitempty"`
	// The list of explorers for the token address.
	Explorers *[]genericmodels.Explorer `json:"explorers,omitempty"`
	// The total liquidity converted to fiat in `quote-currency`.
	TotalLiquidityQuote *float64 `json:"total_liquidity_quote,omitempty"`
	// A prettier version of the total liquidity quote for rendering purposes.
	PrettyTotalLiquidityQuote *string `json:"pretty_total_liquidity_quote,omitempty"`
	// The volume 24h converted to fiat in `quote-currency`.
	Volume24hQuote *float64 `json:"volume_24h_quote,omitempty"`
	// The volume 7d converted to fiat in `quote-currency`.
	Volume7dQuote *float64 `json:"volume_7d_quote,omitempty"`
	// The fee 24h converted to fiat in `quote-currency`.
	Fee24hQuote *float64 `json:"fee_24h_quote,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// A prettier version of the quote rate for rendering purposes.
	PrettyQuoteRate *string `json:"pretty_quote_rate,omitempty"`
	// The annual fee percentage.
	AnnualizedFee *float64 `json:"annualized_fee,omitempty"`
	// A prettier version of the volume 24h quote for rendering purposes.
	PrettyVolume24hQuote *string `json:"pretty_volume_24h_quote,omitempty"`
	// A prettier version of the volume 7d quote for rendering purposes.
	PrettyVolume7dQuote *string `json:"pretty_volume_7d_quote,omitempty"`
	// A prettier version of the fee 24h quote for rendering purposes.
	PrettyFee24hQuote *string `json:"pretty_fee_24h_quote,omitempty"`
	// Token0's contract metadata and reserve data.
	Token0 *PoolsDexToken `json:"token_0,omitempty"`
	// Token1's contract metadata and reserve data.
	Token1 *PoolsDexToken `json:"token_1,omitempty"`
}
type PoolsDexToken struct {
	// The reserves for the token.
	Reserve *string `json:"reserve,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
}
type AddressExchangeBalancesResponse struct {
	// The requested address.
	Address string `json:"address"`
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []UniswapLikeBalanceItem `json:"items"`
}
type UniswapLikeBalanceItem struct {
	Token0    *UniswapLikeToken           `json:"token_0,omitempty"`
	Token1    *UniswapLikeToken           `json:"token_1,omitempty"`
	PoolToken *UniswapLikeTokenWithSupply `json:"pool_token,omitempty"`
}
type UniswapLikeToken struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	Quote   *float64      `json:"quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyQuote *string `json:"pretty_quote,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
}
type UniswapLikeTokenWithSupply struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The asset balance. Use `contract_decimals` to scale this balance for display purposes.
	Balance *utils.BigInt `json:"balance,omitempty"`
	Quote   *float64      `json:"quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyQuote *string `json:"pretty_quote,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// Total supply of this pool token.
	TotalSupply *utils.BigInt `json:"total_supply,omitempty"`
}
type NetworkExchangeTokensResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []TokenV2Volume `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type TokenV2Volume struct {
	// The requested chain name eg: `eth-mainnet`.
	ChainName *string `json:"chain_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The string returned by the `name()` method.
	ContractName   *string `json:"contract_name,omitempty"`
	TotalLiquidity *string `json:"total_liquidity,omitempty"`
	TotalVolume24h *string `json:"total_volume_24h,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int   `json:"contract_decimals,omitempty"`
	SwapCount24h     *int64 `json:"swap_count_24h,omitempty"`
	// The list of explorers for the token address.
	Explorers *[]genericmodels.Explorer `json:"explorers,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The 24h exchange rate for the requested quote currency.
	QuoteRate24h *float64 `json:"quote_rate_24h,omitempty"`
	// A prettier version of the exchange rate for rendering purposes.
	PrettyQuoteRate *string `json:"pretty_quote_rate,omitempty"`
	// A prettier version of the 24h exchange rate for rendering purposes.
	PrettyQuoteRate24h *string `json:"pretty_quote_rate_24h,omitempty"`
	// A prettier version of the total liquidity quote for rendering purposes.
	PrettyTotalLiquidityQuote *string `json:"pretty_total_liquidity_quote,omitempty"`
	// A prettier version of the 24h volume quote for rendering purposes.
	PrettyTotalVolume24hQuote *string `json:"pretty_total_volume_24h_quote,omitempty"`
	// The total liquidity converted to fiat in `quote-currency`.
	TotalLiquidityQuote *float64 `json:"total_liquidity_quote,omitempty"`
	// The total volume 24h converted to fiat in `quote-currency`.
	TotalVolume24hQuote *float64 `json:"total_volume_24h_quote,omitempty"`
}
type NetworkExchangeTokenViewResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []TokenV2VolumeWithChartData `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type TokenV2VolumeWithChartData struct {
	// The requested chain name eg: `eth-mainnet`.
	ChainName *string `json:"chain_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// A list of explorers for this address.
	Explorers *[]genericmodels.Explorer `json:"explorers,omitempty"`
	// The total liquidity unscaled value.
	TotalLiquidity *string `json:"total_liquidity,omitempty"`
	// The total volume 24h unscaled value.
	TotalVolume24h *string `json:"total_volume_24h,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The total amount of swaps in the last 24h.
	SwapCount24h *int64 `json:"swap_count_24h,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// The 24h exchange rate for the requested quote currency.
	QuoteRate24h *float64 `json:"quote_rate_24h,omitempty"`
	// A prettier version of the exchange rate for rendering purposes.
	PrettyQuoteRate *string `json:"pretty_quote_rate,omitempty"`
	// A prettier version of the 24h exchange rate for rendering purposes.
	PrettyQuoteRate24h *string `json:"pretty_quote_rate_24h,omitempty"`
	// A prettier version of the total liquidity quote for rendering purposes.
	PrettyTotalLiquidityQuote *string `json:"pretty_total_liquidity_quote,omitempty"`
	// A prettier version of the 24h volume quote for rendering purposes.
	PrettyTotalVolume24hQuote *string `json:"pretty_total_volume_24h_quote,omitempty"`
	// The total liquidity converted to fiat in `quote-currency`.
	TotalLiquidityQuote *float64 `json:"total_liquidity_quote,omitempty"`
	// The total volume 24h converted to fiat in `quote-currency`.
	TotalVolume24hQuote *float64 `json:"total_volume_24h_quote,omitempty"`
	// The number of transactions in the last 24h.
	Transactions24h        *int64                      `json:"transactions_24h,omitempty"`
	VolumeTimeseries7d     *[]VolumeTokenTimeseries    `json:"volume_timeseries_7d,omitempty"`
	VolumeTimeseries30d    *[]VolumeTokenTimeseries    `json:"volume_timeseries_30d,omitempty"`
	LiquidityTimeseries7d  *[]LiquidityTokenTimeseries `json:"liquidity_timeseries_7d,omitempty"`
	LiquidityTimeseries30d *[]LiquidityTokenTimeseries `json:"liquidity_timeseries_30d,omitempty"`
	PriceTimeseries7d      *[]PriceTokenTimeseries     `json:"price_timeseries_7d,omitempty"`
	PriceTimeseries30d     *[]PriceTokenTimeseries     `json:"price_timeseries_30d,omitempty"`
}
type VolumeTokenTimeseries struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The current date.
	Dt *time.Time `json:"dt,omitempty"`
	// The total volume unscaled for this day.
	TotalVolume *string `json:"total_volume,omitempty"`
	// The volume in `quote-currency` denomination.
	VolumeQuote *float64 `json:"volume_quote,omitempty"`
	// A prettier version of the volume quote for rendering purposes.
	PrettyVolumeQuote *string `json:"pretty_volume_quote,omitempty"`
}
type LiquidityTokenTimeseries struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The current date.
	Dt *time.Time `json:"dt,omitempty"`
	// The total liquidity unscaled up to this day.
	TotalLiquidity *string `json:"total_liquidity,omitempty"`
	// The liquidity in `quote-currency` denomination.
	LiquidityQuote *float64 `json:"liquidity_quote,omitempty"`
	// A prettier version of the liquidity quote for rendering purposes.
	PrettyLiquidityQuote *string `json:"pretty_liquidity_quote,omitempty"`
}
type PriceTokenTimeseries struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The current date.
	Dt *time.Time `json:"dt,omitempty"`
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *string `json:"quote_currency,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// A prettier version of the exchange rate for rendering purposes.
	PrettyQuoteRate *string `json:"pretty_quote_rate,omitempty"`
}
type SupportedDexesResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// List of response items.
	Items []SupportedDex `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type SingleNetworkExchangeTokenResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []PoolWithTimeseries `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type TransactionsForAccountAddressResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []ExchangeTransaction `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type ExchangeTransaction struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	Act    *string `json:"act,omitempty"`
	// The requested address.
	Address *string `json:"address,omitempty"`
	// A list of explorers for this transaction.
	Explorers     *[]genericmodels.Explorer `json:"explorers,omitempty"`
	Amount0       *string                   `json:"amount_0,omitempty"`
	Amount1       *string                   `json:"amount_1,omitempty"`
	Amount0In     *string                   `json:"amount_0_in,omitempty"`
	Amount0Out    *string                   `json:"amount_0_out,omitempty"`
	Amount1In     *string                   `json:"amount_1_in,omitempty"`
	Amount1Out    *string                   `json:"amount_1_out,omitempty"`
	ToAddress     *string                   `json:"to_address,omitempty"`
	FromAddress   *string                   `json:"from_address,omitempty"`
	SenderAddress *string                   `json:"sender_address,omitempty"`
	TotalQuote    *float64                  `json:"total_quote,omitempty"`
	// A prettier version of the total quote for rendering purposes.
	PrettyTotalQuote *string `json:"pretty_total_quote,omitempty"`
	// The value attached to this tx.
	Value *utils.BigInt `json:"value,omitempty"`
	// The value attached in `quote-currency` to this tx.
	ValueQuote *float64 `json:"value_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyValueQuote *string `json:"pretty_value_quote,omitempty"`
	// The requested chain native gas token metadata.
	GasMetadata *genericmodels.ContractMetadata `json:"gas_metadata,omitempty"`
	// The amount of gas supplied for this tx.
	GasOffered *int64 `json:"gas_offered,omitempty"`
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
	// The requested quote currency eg: `USD`.
	QuoteCurrency   *string    `json:"quote_currency,omitempty"`
	Token0          *PoolToken `json:"token_0,omitempty"`
	Token1          *PoolToken `json:"token_1,omitempty"`
	Token0QuoteRate *float64   `json:"token_0_quote_rate,omitempty"`
	Token1QuoteRate *float64   `json:"token_1_quote_rate,omitempty"`
}
type PoolToken struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The string returned by the `name()` method.
	ContractName *string `json:"contract_name,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc *bool `json:"supports_erc,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
}
type TransactionsForTokenAddressResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []ExchangeTransaction `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type TransactionsForExchangeResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []ExchangeTransaction `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type NetworkTransactionsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []ExchangeTransaction `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type EcosystemChartDataResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []UniswapLikeEcosystemCharts `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type UniswapLikeEcosystemCharts struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string `json:"chain_id,omitempty"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency      *string  `json:"quote_currency,omitempty"`
	GasTokenPriceQuote *float64 `json:"gas_token_price_quote,omitempty"`
	TotalSwaps24h      *int64   `json:"total_swaps_24h,omitempty"`
	TotalActivePairs7d *int64   `json:"total_active_pairs_7d,omitempty"`
	TotalFees24h       *float64 `json:"total_fees_24h,omitempty"`
	// A prettier version of the gas quote for rendering purposes.
	PrettyGasTokenPriceQuote *string `json:"pretty_gas_token_price_quote,omitempty"`
	// A prettier version of the 24h total fees for rendering purposes.
	PrettyTotalFees24h *string                    `json:"pretty_total_fees_24h,omitempty"`
	VolumeChart7d      *[]VolumeEcosystemChart    `json:"volume_chart_7d,omitempty"`
	VolumeChart30d     *[]VolumeEcosystemChart    `json:"volume_chart_30d,omitempty"`
	LiquidityChart7d   *[]LiquidityEcosystemChart `json:"liquidity_chart_7d,omitempty"`
	LiquidityChart30d  *[]LiquidityEcosystemChart `json:"liquidity_chart_30d,omitempty"`
}
type VolumeEcosystemChart struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string    `json:"chain_id,omitempty"`
	Dt      *time.Time `json:"dt,omitempty"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency *string  `json:"quote_currency,omitempty"`
	VolumeQuote   *float64 `json:"volume_quote,omitempty"`
	// A prettier version of the volume quote for rendering purposes.
	PrettyVolumeQuote *string `json:"pretty_volume_quote,omitempty"`
	SwapCount24       *int64  `json:"swap_count_24,omitempty"`
}
type LiquidityEcosystemChart struct {
	// The name of the DEX, eg: `uniswap_v2`.
	DexName *string `json:"dex_name,omitempty"`
	// The requested chain ID eg: `1`.
	ChainId *string    `json:"chain_id,omitempty"`
	Dt      *time.Time `json:"dt,omitempty"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency  *string  `json:"quote_currency,omitempty"`
	LiquidityQuote *float64 `json:"liquidity_quote,omitempty"`
	// A prettier version of the liquidity quote for rendering purposes.
	PrettyLiquidityQuote *string `json:"pretty_liquidity_quote,omitempty"`
}
type HealthDataResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// List of response items.
	Items []HealthData `json:"items"`
	// Pagination metadata.
	Pagination genericmodels.Pagination `json:"pagination"`
}
type HealthData struct {
	SyncedBlockHeight   *int       `json:"synced_block_height,omitempty"`
	SyncedBlockSignedAt *time.Time `json:"synced_block_signed_at,omitempty"`
	LatestBlockHeight   *int       `json:"latest_block_height,omitempty"`
	LatestBlockSignedAt *time.Time `json:"latest_block_signed_at,omitempty"`
}

type GetPoolsQueryParamOpts struct {
	// Ending date to define a block range (YYYY-MM-DD). Omitting this parameter defaults to the current date.
	Date *string `json:"date,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetPoolsForTokenAddressQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// The DEX name eg: `uniswap_v2`.
	DexName *string `json:"dexName,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
}
type GetPoolsForWalletAddressQueryParamOpts struct {
	// The token contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.
	TokenAddress *string `json:"tokenAddress,omitempty"`
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// The DEX name eg: `uniswap_v2`.
	DexName *string `json:"dexName,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
}
type GetNetworkExchangeTokensQueryParamOpts struct {
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetLpTokenViewQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
}
type GetSingleNetworkExchangeTokenQueryParamOpts struct {
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetTransactionsForTokenAddressQueryParamOpts struct {
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetTransactionsForExchangeQueryParamOpts struct {
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}
type GetTransactionsForDexQueryParamOpts struct {
	// The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.
	QuoteCurrency *quotes.Quote `json:"quoteCurrency,omitempty"`
	// Number of items per page. Omitting this parameter defaults to 100.
	PageSize *int `json:"pageSize,omitempty"`
	// 0-indexed page number to begin pagination.
	PageNumber *int `json:"pageNumber,omitempty"`
}

func NewXykServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) XykService {

	return &xykServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type XykService interface {

	// Commonly used to get all the pools of a particular DEX. Supports most common DEXs (Uniswap, SushiSwap, etc), and returns detailed trading data (volume, liquidity, swap counts, fees, LP token prices).
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	GetPools(chainName chains.Chain, dexName string, queryParamOpts ...GetPoolsQueryParamOpts) (*utils.Response[PoolResponse], error)

	// Commonly used to get the corresponding supported DEX given a pool address, along with the swap fees, DEX's logo url, and factory addresses. Useful to identifying the specific DEX to which a pair address is associated.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// poolAddress: The requested pool address.. Type: string
	GetDexForPoolAddress(chainName chains.Chain, poolAddress string) (*utils.Response[PoolToDexResponse], error)

	// Commonly used to get the 7 day and 30 day time-series data (volume, liquidity, price) of a particular liquidity pool in a DEX. Useful for building time-series charts on DEX trading activity.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// poolAddress: The pool contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetPoolByAddress(chainName chains.Chain, dexName string, poolAddress string) (*utils.Response[PoolByAddressResponse], error)

	// Commonly used to get all pools and the supported DEX for a token. Useful for building a table of top pairs across all supported DEXes that the token is trading on.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// tokenAddress: The token contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	// page: The requested 0-indexed page number.. Type: int
	GetPoolsForTokenAddress(chainName chains.Chain, tokenAddress string, page int, queryParamOpts ...GetPoolsForTokenAddressQueryParamOpts) (*utils.Response[PoolsDexDataResponse], error)

	// Commonly used to return balance of a wallet/contract address on a specific DEX.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// accountAddress: The account address.. Type: string
	GetAddressExchangeBalances(chainName chains.Chain, dexName string, accountAddress string) (*utils.Response[AddressExchangeBalancesResponse], error)

	// Commonly used to get all pools and supported DEX for a wallet. Useful for building a personal DEX UI showcasing pairs and supported DEXes associated to the wallet.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The account address.. Type: string
	// page: The requested 0-indexed page number.. Type: int
	GetPoolsForWalletAddress(chainName chains.Chain, walletAddress string, page int, queryParamOpts ...GetPoolsForWalletAddressQueryParamOpts) (*utils.Response[PoolsDexDataResponse], error)

	// Commonly used to retrieve all network exchange tokens for a specific DEX. Useful for building a top tokens table by total liquidity within a particular DEX.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	GetNetworkExchangeTokens(chainName chains.Chain, dexName string, queryParamOpts ...GetNetworkExchangeTokensQueryParamOpts) (*utils.Response[NetworkExchangeTokensResponse], error)

	// Commonly used to get a detailed view for a single liquidity pool token. Includes time series data.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// tokenAddress: The token contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetLpTokenView(chainName chains.Chain, dexName string, tokenAddress string, queryParamOpts ...GetLpTokenViewQueryParamOpts) (*utils.Response[NetworkExchangeTokenViewResponse], error)

	// Commonly used to get all the supported DEXs available for the xy=k endpoints, along with the swap fees and factory addresses.
	//   Parameters:

	GetSupportedDEXes() (*utils.Response[SupportedDexesResponse], error)

	// Commonly used to get historical daily swap count for a single network exchange token.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// tokenAddress: The token contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetSingleNetworkExchangeToken(chainName chains.Chain, dexName string, tokenAddress string, queryParamOpts ...GetSingleNetworkExchangeTokenQueryParamOpts) (*utils.Response[SingleNetworkExchangeTokenResponse], error)

	// Commonly used to get all the DEX transactions of a wallet. Useful for building tables of DEX activity segmented by wallet.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// accountAddress: The account address. Passing in an `ENS` or `RNS` resolves automatically.. Type: string
	GetTransactionsForAccountAddress(chainName chains.Chain, dexName string, accountAddress string) (*utils.Response[TransactionsForAccountAddressResponse], error)

	// Commonly used to get all the transactions of a token within a particular DEX. Useful for getting a per-token view of DEX activity.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// tokenAddress: The token contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTransactionsForTokenAddress(chainName chains.Chain, dexName string, tokenAddress string, queryParamOpts ...GetTransactionsForTokenAddressQueryParamOpts) (*utils.Response[TransactionsForTokenAddressResponse], error)

	// Commonly used for getting all the transactions of a particular DEX liquidity pool. Useful for building a transactions history table for an individual pool.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	// poolAddress: The pool contract address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetTransactionsForExchange(chainName chains.Chain, dexName string, poolAddress string, queryParamOpts ...GetTransactionsForExchangeQueryParamOpts) (*utils.Response[TransactionsForExchangeResponse], error)

	// Commonly used to get all the the transactions for a given DEX. Useful for building DEX activity views.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	GetTransactionsForDex(chainName chains.Chain, dexName string, queryParamOpts ...GetTransactionsForDexQueryParamOpts) (*utils.Response[NetworkTransactionsResponse], error)

	// Commonly used to get a 7d and 30d time-series chart of DEX activity. Includes volume and swap count.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	GetEcosystemChartData(chainName chains.Chain, dexName string) (*utils.Response[EcosystemChartDataResponse], error)

	// Commonly used to ping the health of xy=k endpoints to get the synced block height per chain.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// dexName: The DEX name eg: `uniswap_v2`.. Type: string
	GetHealthData(chainName chains.Chain, dexName string) (*utils.Response[HealthDataResponse], error)
}

type xykServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *xykServiceImpl) GetPools(chainName chains.Chain, dexName string, queryParamOpts ...GetPoolsQueryParamOpts) (*utils.Response[PoolResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/pools/", chainName, dexName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[PoolResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

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
		return &utils.Response[PoolResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[PoolResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[PoolResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[PoolResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[PoolResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetDexForPoolAddress(chainName chains.Chain, poolAddress string) (*utils.Response[PoolToDexResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/address/%s/dex_name/", chainName, poolAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[PoolToDexResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[PoolToDexResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[PoolToDexResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[PoolToDexResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolToDexResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolToDexResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[PoolToDexResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[PoolToDexResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetPoolByAddress(chainName chains.Chain, dexName string, poolAddress string) (*utils.Response[PoolByAddressResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/pools/address/%s/", chainName, dexName, poolAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[PoolByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[PoolByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[PoolByAddressResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[PoolByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolByAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[PoolByAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[PoolByAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetPoolsForTokenAddress(chainName chains.Chain, tokenAddress string, page int, queryParamOpts ...GetPoolsForTokenAddressQueryParamOpts) (*utils.Response[PoolsDexDataResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/tokens/address/%s/pools/page/%d/", chainName, tokenAddress, page)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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

		if opts.DexName != nil {
			params.Add("dex-name", fmt.Sprintf("%v", *opts.DexName))
		}

		if opts.PageSize != nil {
			params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
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
		return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[PoolsDexDataResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[PoolsDexDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[PoolsDexDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetAddressExchangeBalances(chainName chains.Chain, dexName string, accountAddress string) (*utils.Response[AddressExchangeBalancesResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/address/%s/balances/", chainName, dexName, accountAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[AddressExchangeBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[AddressExchangeBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[AddressExchangeBalancesResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[AddressExchangeBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[AddressExchangeBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[AddressExchangeBalancesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[AddressExchangeBalancesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[AddressExchangeBalancesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetPoolsForWalletAddress(chainName chains.Chain, walletAddress string, page int, queryParamOpts ...GetPoolsForWalletAddressQueryParamOpts) (*utils.Response[PoolsDexDataResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/address/%s/pools/page/%d/", chainName, walletAddress, page)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.TokenAddress != nil {
			params.Add("token-address", fmt.Sprintf("%v", *opts.TokenAddress))
		}

		if opts.QuoteCurrency != nil {
			params.Add("quote-currency", fmt.Sprintf("%v", *opts.QuoteCurrency))
		}

		if opts.DexName != nil {
			params.Add("dex-name", fmt.Sprintf("%v", *opts.DexName))
		}

		if opts.PageSize != nil {
			params.Add("page-size", fmt.Sprintf("%v", *opts.PageSize))
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
		return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[PoolsDexDataResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[PoolsDexDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[PoolsDexDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[PoolsDexDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetNetworkExchangeTokens(chainName chains.Chain, dexName string, queryParamOpts ...GetNetworkExchangeTokensQueryParamOpts) (*utils.Response[NetworkExchangeTokensResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/tokens/", chainName, dexName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NetworkExchangeTokensResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NetworkExchangeTokensResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NetworkExchangeTokensResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NetworkExchangeTokensResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NetworkExchangeTokensResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NetworkExchangeTokensResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NetworkExchangeTokensResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NetworkExchangeTokensResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetLpTokenView(chainName chains.Chain, dexName string, tokenAddress string, queryParamOpts ...GetLpTokenViewQueryParamOpts) (*utils.Response[NetworkExchangeTokenViewResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/tokens/address/%s/view/", chainName, dexName, tokenAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NetworkExchangeTokenViewResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NetworkExchangeTokenViewResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NetworkExchangeTokenViewResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NetworkExchangeTokenViewResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NetworkExchangeTokenViewResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NetworkExchangeTokenViewResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NetworkExchangeTokenViewResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NetworkExchangeTokenViewResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetSupportedDEXes() (*utils.Response[SupportedDexesResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/xy=k/supported_dexes/")

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[SupportedDexesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[SupportedDexesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[SupportedDexesResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[SupportedDexesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[SupportedDexesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[SupportedDexesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[SupportedDexesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[SupportedDexesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetSingleNetworkExchangeToken(chainName chains.Chain, dexName string, tokenAddress string, queryParamOpts ...GetSingleNetworkExchangeTokenQueryParamOpts) (*utils.Response[SingleNetworkExchangeTokenResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/tokens/address/%s/", chainName, dexName, tokenAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[SingleNetworkExchangeTokenResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[SingleNetworkExchangeTokenResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetTransactionsForAccountAddress(chainName chains.Chain, dexName string, accountAddress string) (*utils.Response[TransactionsForAccountAddressResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/address/%s/transactions/", chainName, dexName, accountAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsForAccountAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[TransactionsForAccountAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[TransactionsForAccountAddressResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsForAccountAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsForAccountAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsForAccountAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsForAccountAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsForAccountAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetTransactionsForTokenAddress(chainName chains.Chain, dexName string, tokenAddress string, queryParamOpts ...GetTransactionsForTokenAddressQueryParamOpts) (*utils.Response[TransactionsForTokenAddressResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/tokens/address/%s/transactions/", chainName, dexName, tokenAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsForTokenAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[TransactionsForTokenAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[TransactionsForTokenAddressResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsForTokenAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsForTokenAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsForTokenAddressResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsForTokenAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsForTokenAddressResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetTransactionsForExchange(chainName chains.Chain, dexName string, poolAddress string, queryParamOpts ...GetTransactionsForExchangeQueryParamOpts) (*utils.Response[TransactionsForExchangeResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/pools/address/%s/transactions/", chainName, dexName, poolAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[TransactionsForExchangeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[TransactionsForExchangeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[TransactionsForExchangeResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[TransactionsForExchangeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsForExchangeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[TransactionsForExchangeResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[TransactionsForExchangeResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[TransactionsForExchangeResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetTransactionsForDex(chainName chains.Chain, dexName string, queryParamOpts ...GetTransactionsForDexQueryParamOpts) (*utils.Response[NetworkTransactionsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/transactions/", chainName, dexName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NetworkTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NetworkTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NetworkTransactionsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NetworkTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NetworkTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NetworkTransactionsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NetworkTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NetworkTransactionsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetEcosystemChartData(chainName chains.Chain, dexName string) (*utils.Response[EcosystemChartDataResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/ecosystem/", chainName, dexName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[EcosystemChartDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[EcosystemChartDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[EcosystemChartDataResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[EcosystemChartDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[EcosystemChartDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[EcosystemChartDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[EcosystemChartDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[EcosystemChartDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *xykServiceImpl) GetHealthData(chainName chains.Chain, dexName string) (*utils.Response[HealthDataResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/xy=k/%s/health/", chainName, dexName)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[HealthDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[HealthDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[HealthDataResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[HealthDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[HealthDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[HealthDataResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[HealthDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[HealthDataResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
