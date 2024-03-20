package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

type ApprovalsResponse struct {
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
	Items []TokensApprovalItem `json:"items"`
}
type TokensApprovalItem struct {
	// The address for the token that has approvals.
	TokenAddress *string `json:"token_address,omitempty"`
	// The name for the token that has approvals.
	TokenAddressLabel *string `json:"token_address_label,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	TickerSymbol *string `json:"ticker_symbol,omitempty"`
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals *int `json:"contract_decimals,omitempty"`
	// The contract logo URL.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The exchange rate for the requested quote currency.
	QuoteRate *float64 `json:"quote_rate,omitempty"`
	// Wallet balance of the token.
	Balance *utils.BigInt `json:"balance,omitempty"`
	// Value of the wallet balance of the token.
	BalanceQuote *float64 `json:"balance_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyBalanceQuote *string `json:"pretty_balance_quote,omitempty"`
	// Total amount at risk across all spenders.
	ValueAtRisk *string `json:"value_at_risk,omitempty"`
	// Value of total amount at risk across all spenders.
	ValueAtRiskQuote *float64 `json:"value_at_risk_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyValueAtRiskQuote *string `json:"pretty_value_at_risk_quote,omitempty"`
	// Contracts with non-zero approvals for this token.
	Spenders *[]TokenSpenderItem `json:"spenders,omitempty"`
}
type TokenSpenderItem struct {
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
	// The offset is the position of the tx in the block.
	TxOffset *int64 `json:"tx_offset,omitempty"`
	// The offset is the position of the log entry within an event log."
	LogOffset *int64 `json:"log_offset,omitempty"`
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// Most recent transaction that updated approval amounts for the token.
	TxHash *string `json:"tx_hash,omitempty"`
	// Address of the contract with approval for the token.
	SpenderAddress *string `json:"spender_address,omitempty"`
	// Name of the contract with approval for the token.
	SpenderAddressLabel *string `json:"spender_address_label,omitempty"`
	// Remaining number of tokens granted to the spender by the approval.
	Allowance *string `json:"allowance,omitempty"`
	// Value of the remaining allowance specified by the approval.
	AllowanceQuote *float64 `json:"allowance_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyAllowanceQuote *string `json:"pretty_allowance_quote,omitempty"`
	// Amount at risk for spender.
	ValueAtRisk *string `json:"value_at_risk,omitempty"`
	// Value of amount at risk for spender.
	ValueAtRiskQuote *float64 `json:"value_at_risk_quote,omitempty"`
	// A prettier version of the quote for rendering purposes.
	PrettyValueAtRiskQuote *string `json:"pretty_value_at_risk_quote,omitempty"`
	RiskFactor             *string `json:"risk_factor,omitempty"`
}
type NftApprovalsResponse struct {
	// The timestamp when the response was generated. Useful to show data staleness to users.
	UpdatedAt time.Time `json:"updated_at"`
	// The requested chain ID eg: `1`.
	ChainId int `json:"chain_id"`
	// The requested chain name eg: `eth-mainnet`.
	ChainName string `json:"chain_name"`
	// The requested address.
	Address string `json:"address"`
	// List of response items.
	Items []NftApprovalsItem `json:"items"`
}
type NftApprovalsItem struct {
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress *string `json:"contract_address,omitempty"`
	// The label of the contract address.
	ContractAddressLabel *string `json:"contract_address_label,omitempty"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol *string `json:"contract_ticker_symbol,omitempty"`
	// List of asset balances held by the user.
	TokenBalances *[]NftApprovalBalance `json:"token_balances,omitempty"`
	// Contracts with non-zero approvals for this token.
	Spenders *[]NftApprovalSpender `json:"spenders,omitempty"`
}
type NftApprovalBalance struct {
	// The token's id.
	TokenId *utils.BigInt `json:"token_id,omitempty"`
	// The NFT's token balance.
	TokenBalance *utils.BigInt `json:"token_balance,omitempty"`
}
type NftApprovalSpender struct {
	// The height of the block.
	BlockHeight *int64 `json:"block_height,omitempty"`
	// The offset is the position of the tx in the block.
	TxOffset *int64 `json:"tx_offset,omitempty"`
	// The offset is the position of the log entry within an event log."
	LogOffset *int64 `json:"log_offset,omitempty"`
	// The block signed timestamp in UTC.
	BlockSignedAt *time.Time `json:"block_signed_at,omitempty"`
	// Most recent transaction that updated approval amounts for the token.
	TxHash *string `json:"tx_hash,omitempty"`
	// Address of the contract with approval for the token.
	SpenderAddress *string `json:"spender_address,omitempty"`
	// Name of the contract with approval for the token.
	SpenderAddressLabel *string `json:"spender_address_label,omitempty"`
	// The token ids approved.
	TokenIdsApproved *string `json:"token_ids_approved,omitempty"`
	// Remaining number of tokens granted to the spender by the approval.
	Allowance *string `json:"allowance,omitempty"`
}

func NewSecurityServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) SecurityService {

	return &securityServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type SecurityService interface {

	// Commonly used to get a list of approvals across all token contracts categorized by spenders for a wallet’s assets.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetApprovals(chainName chains.Chain, walletAddress string) (*utils.Response[ApprovalsResponse], error)

	// Commonly used to get a list of NFT approvals across all token contracts categorized by spenders for a wallet’s assets.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// walletAddress: The requested address. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically.. Type: string
	GetNftApprovals(chainName chains.Chain, walletAddress string) (*utils.Response[NftApprovalsResponse], error)
}

type securityServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *securityServiceImpl) GetApprovals(chainName chains.Chain, walletAddress string) (*utils.Response[ApprovalsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/approvals/%s/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[ApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[ApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[ApprovalsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[ApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[ApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[ApprovalsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[ApprovalsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}

func (s *securityServiceImpl) GetNftApprovals(chainName chains.Chain, walletAddress string) (*utils.Response[NftApprovalsResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/%v/nft/approvals/%s/", chainName, walletAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &utils.Response[NftApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
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
		return &utils.Response[NftApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data utils.Response[NftApprovalsResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &utils.Response[NftApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &utils.Response[NftApprovalsResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &utils.Response[NftApprovalsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &utils.Response[NftApprovalsResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
