package utils

import (
	"regexp"
)

type ApiKeyValidator struct {
	APIKey          string
	APIKeyV1Pattern *regexp.Regexp
	APIKeyV2Pattern *regexp.Regexp
}

const InvalidAPIKeyMessage string = "invalid or missing API key (sign up at covalenthq.com/platform)"

// NewApiKeyValidator is a constructor function for ApiKeyValidator
func NewApiKeyValidator(apiKey string) *ApiKeyValidator {
	return &ApiKeyValidator{
		APIKey:          apiKey,
		APIKeyV1Pattern: regexp.MustCompile(`^ckey_([a-f0-9]{27})$`),
		APIKeyV2Pattern: regexp.MustCompile(`^cqt_(wF|rQ)([bcdfghjkmpqrtvwxyBCDFGHJKMPQRTVWXY346789]{26})$`),
	}
}

// IsValidApiKey checks if the provided API key matches either of the patterns
func (v *ApiKeyValidator) IsValidApiKey() bool {
	return v.APIKeyV1Pattern.MatchString(v.APIKey) || v.APIKeyV2Pattern.MatchString(v.APIKey)
}
