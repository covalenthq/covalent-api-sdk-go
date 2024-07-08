package covalentclient

import (
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

type CovalentClientSettings struct {
	// Toggle to analyze the execution of each api request.
	Debug *bool `json:"debug,omitempty"`
	//  The number of concurrent requests allowed.
	ThreadCount *int `json:"thread_count,omitempty"`
	// The base URL for the Covalent API
	BaseURL string `json:"base_url,omitempty"`
}

type CovalentClientType struct {
	SecurityService    services.SecurityService
	BalanceService     services.BalanceService
	BaseService        services.BaseService
	NftService         services.NftService
	PricingService     services.PricingService
	TransactionService services.TransactionService
	XykService         services.XykService
	Debug              bool
	ThreadCount        int
}

var defaultDebug bool = false
var defaultThreadCount int = 3

func CovalentClient(apiKey string, settings ...CovalentClientSettings) *CovalentClientType {
	client := &CovalentClientType{}
	validator := utils.NewApiKeyValidator(apiKey)
	isValidKey := validator.IsValidApiKey()
	baseURL := "https://api.covalenthq.com"

	if len(settings) == 0 {
		client.Debug = defaultDebug
		client.ThreadCount = defaultThreadCount
	} else {
		// Settings were provided, apply them
		setting := settings[0] // Assuming only one settings struct is passed

		if setting.Debug == nil {
			client.Debug = defaultDebug
		} else {
			client.Debug = *setting.Debug
		}

		if setting.ThreadCount == nil {
			client.ThreadCount = defaultThreadCount
		} else {
			client.ThreadCount = *setting.ThreadCount
		}

		if setting.BaseURL != "" {
			baseURL = setting.BaseURL
		}
	}

	client.SecurityService = services.NewSecurityServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)
	client.BalanceService = services.NewBalanceServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)
	client.BaseService = services.NewBaseServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)
	client.NftService = services.NewNftServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)
	client.PricingService = services.NewPricingServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)
	client.TransactionService = services.NewTransactionServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)
	client.XykService = services.NewXykServiceImpl(baseURL, apiKey, client.Debug, client.ThreadCount, isValidKey)

	return client

}
