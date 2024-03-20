package genericmodels

import (
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

type NftCollectionAttribute struct {
	TraitType *string      `json:"trait_type,omitempty"`
	Value     *interface{} `json:"value,omitempty"`
}
type DecodedItem struct {
	Name      *string  `json:"name,omitempty"`
	Signature *string  `json:"signature,omitempty"`
	Params    *[]Param `json:"params,omitempty"`
}
type Param struct {
	Name    *string      `json:"name,omitempty"`
	Type    *string      `json:"type,omitempty"`
	Indexed *bool        `json:"indexed,omitempty"`
	Decoded *bool        `json:"decoded,omitempty"`
	Value   *interface{} `json:"value,omitempty"`
}
type LogEvent struct {
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
	// The offset is the position of the tx in the block.
	TxOffset *int64 `json:"tx_offset,omitempty"`
	// The offset is the position of the log entry within an event log.
	LogOffset *int64 `json:"log_offset,omitempty"`
	// The requested transaction hash.
	TxHash *string `json:"tx_hash,omitempty"`
	// The log topics in raw data.
	RawLogTopics *[]string `json:"raw_log_topics,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	SenderContractDecimals *int `json:"sender_contract_decimals,omitempty"`
	// The name of the sender.
	SenderName                 *string `json:"sender_name,omitempty"`
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
	Decoded *DecodedItem `json:"decoded,omitempty"`
}
type ContractMetadata struct {
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
}
type Explorer struct {
	// The name of the explorer.
	Label *string `json:"label,omitempty"`
	// The URL of the explorer.
	Url *string `json:"url,omitempty"`
}
type Pagination struct {
	// True is there is another page.
	HasMore *bool `json:"has_more,omitempty"`
	// The requested page number.
	PageNumber *int `json:"page_number,omitempty"`
	// The requested number of items on the current page.
	PageSize *int `json:"page_size,omitempty"`
	// The total number of items across all pages for this request.
	TotalCount *int `json:"total_count,omitempty"`
}
type LogoUrls struct {
	// The token logo URL.
	TokenLogoUrl *string `json:"token_logo_url,omitempty"`
	// The protocol logo URL.
	ProtocolLogoUrl *string `json:"protocol_logo_url,omitempty"`
	// The chain logo URL.
	ChainLogoUrl *string `json:"chain_logo_url,omitempty"`
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
	Name               *string                   `json:"name,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	AssetUrl           *string                   `json:"asset_url,omitempty"`
	AssetFileExtension *string                   `json:"asset_file_extension,omitempty"`
	AssetMimeType      *string                   `json:"asset_mime_type,omitempty"`
	AssetSizeBytes     *string                   `json:"asset_size_bytes,omitempty"`
	Image              *string                   `json:"image,omitempty"`
	Image256           *string                   `json:"image_256,omitempty"`
	Image512           *string                   `json:"image_512,omitempty"`
	Image1024          *string                   `json:"image_1024,omitempty"`
	AnimationUrl       *string                   `json:"animation_url,omitempty"`
	ExternalUrl        *string                   `json:"external_url,omitempty"`
	Attributes         *[]NftCollectionAttribute `json:"attributes,omitempty"`
}
