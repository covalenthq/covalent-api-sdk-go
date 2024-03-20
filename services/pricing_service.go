package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/genericmodels"
	"github.com/covalenthq/covalent-api-sdk-go/quotes"
	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

type TokenPricesResponse struct {
	// Use contract decimals to format the token balance for display purposes - divide the balance by `10^{contract_decimals}`.
	ContractDecimals int `json:"contract_decimals"`
	// The string returned by the `name()` method.
	ContractName string `json:"contract_name"`
	// The ticker symbol for this contract. This field is set by a developer and non-unique across a network.
	ContractTickerSymbol string `json:"contract_ticker_symbol"`
	// Use the relevant `contract_address` to lookup prices, logos, token transfers, etc.
	ContractAddress string `json:"contract_address"`
	// A list of supported standard ERC interfaces, eg: `ERC20` and `ERC721`.
	SupportsErc []string `json:"supports_erc"`
	// The contract logo URL.
	LogoUrl  string    `json:"logo_url"`
	UpdateAt time.Time `json:"update_at"`
	// The requested quote currency eg: `USD`.
	QuoteCurrency string `json:"quote_currency"`
	// The contract logo URLs.
	LogoUrls LogoUrls `json:"logo_urls"`
	// List of response items.
	Prices []Price `json:"prices"`
	// List of response items.
	Items []Price `json:"items"`
}
type LogoUrls struct {
	// The token logo URL.
	TokenLogoUrl *string `json:"token_logo_url,omitempty"`
	// The protocol logo URL.
	ProtocolLogoUrl *string `json:"protocol_logo_url,omitempty"`
	// The chain logo URL.
	ChainLogoUrl *string `json:"chain_logo_url,omitempty"`
}
type Price struct {
	ContractMetadata *genericmodels.ContractMetadata `json:"contract_metadata,omitempty"`
	// The date of the price capture.
	Date *CustomTime `json:"date,omitempty"`
	// The price in the requested `quote-currency`.
	Price *float64 `json:"price,omitempty"`
	// A prettier version of the price for rendering purposes.
	PrettyPrice *string `json:"pretty_price,omitempty"`
}

type GetTokenPricesQueryParamOpts struct {
	// The start day of the historical price range (YYYY-MM-DD).
	From *string `json:"from,omitempty"`
	// The end day of the historical price range (YYYY-MM-DD).
	To *string `json:"to,omitempty"`
	// Sort the prices in chronological ascending order. By default, it's set to `false` and returns prices in chronological descending order.
	PricesAtAsc *bool `json:"pricesAtAsc,omitempty"`
}

type CustomTime struct {
	*time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Trim quotes since JSON numbers and booleans come in quotes
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		return nil
	}
	// Parse the string to time.Time using the expected layout
	// Adjust "2006-01-02" as needed to match your input format
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	ct.Time = &t
	return nil
}

type Response[T any] struct {
	Data         *[]T    `json:"data,omitempty"`
	Error        bool    `json:"error"`
	ErrorCode    *int    `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
}

func NewPricingServiceImpl(apiKey string, debug bool, threadCount int, isValidKey bool) PricingService {

	return &pricingServiceImpl{APIKey: apiKey, Debug: debug, ThreadCount: threadCount, IskeyValid: isValidKey}
}

type PricingService interface {

	// Commonly used to get historic prices of a token between date ranges. Supports native tokens.
	//   Parameters:
	// chainName: The chain name eg: `eth-mainnet`.. Type: chains.Chain
	// quoteCurrency: The currency to convert. Supports `USD`, `CAD`, `EUR`, `SGD`, `INR`, `JPY`, `VND`, `CNY`, `KRW`, `RUB`, `TRY`, `NGN`, `ARS`, `AUD`, `CHF`, and `GBP`.. Type: quotes.Quote
	// contractAddress: Contract address for the token. Passing in an `ENS`, `RNS`, `Lens Handle`, or an `Unstoppable Domain` resolves automatically. Supports multiple contract addresses separated by commas.. Type: string
	GetTokenPrices(chainName chains.Chain, quoteCurrency quotes.Quote, contractAddress string, queryParamOpts ...GetTokenPricesQueryParamOpts) (*Response[TokenPricesResponse], error)
}

type pricingServiceImpl struct {
	APIKey      string
	Debug       bool
	ThreadCount int
	IskeyValid  bool
}

func (s *pricingServiceImpl) GetTokenPrices(chainName chains.Chain, quoteCurrency quotes.Quote, contractAddress string, queryParamOpts ...GetTokenPricesQueryParamOpts) (*Response[TokenPricesResponse], error) {

	apiURL := fmt.Sprintf("https://api.covalenthq.com/v1/pricing/historical_by_addresses_v2/%v/%v/%s/", chainName, quoteCurrency, contractAddress)

	if !s.IskeyValid {
		errorCode := 401
		errorMessage := utils.InvalidAPIKeyMessage
		return &Response[TokenPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, fmt.Errorf(utils.InvalidAPIKeyMessage)
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if len(queryParamOpts) > 0 {
		opts := queryParamOpts[0]

		if opts.From != nil {
			params.Add("from", fmt.Sprintf("%v", *opts.From))
		}

		if opts.To != nil {
			params.Add("to", fmt.Sprintf("%v", *opts.To))
		}

		if opts.PricesAtAsc != nil {
			params.Add("prices-at-asc", fmt.Sprintf("%v", *opts.PricesAtAsc))
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
		return &Response[TokenPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
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
	var data Response[TokenPricesResponse]

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			errorCode := resp.StatusCode
			errorMessage := err.Error()
			return &Response[TokenPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}

		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			res.Body.Close()
			errorCode := 500
			errorMessage := err.Error()
			return &Response[TokenPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
		res.Body.Close()
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			errorCode := 500
			errorMessage := err.Error()
			return &Response[TokenPricesResponse]{Data: nil, Error: true, ErrorCode: &errorCode, ErrorMessage: &errorMessage}, err
		}
	}

	if data.Error {
		return &Response[TokenPricesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, fmt.Errorf(*data.ErrorMessage)
	}

	return &Response[TokenPricesResponse]{Data: data.Data, Error: data.Error, ErrorCode: data.ErrorCode, ErrorMessage: data.ErrorMessage}, nil
}
