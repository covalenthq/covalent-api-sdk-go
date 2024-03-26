package valuefmt

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/covalenthq/covalent-api-sdk-go/quotes"
)

const (
	LessThanZero = "0.01"
	Zero         = "0.00"
)

var currencyMap = map[quotes.Quote]string{
	"USD": "$",
	"CAD": "CA$",
	"EUR": "€",
	"SGD": "S$",
	"INR": "₹",
	"JPY": "¥",
	"VND": "₫",
	"CNY": "CN¥",
	"KRW": "₩",
	"RUB": "₽",
	"TRY": "₺",
	"NGN": "₦",
	"ARS": "ARS",
	"AUD": "A$",
	"CHF": "CHF",
	"GBP": "£",
}

type CurrencyOptions struct {
	Decimals         int
	Currency         quotes.Quote
	IgnoreSmallValue bool
	IgnoreMinus      bool
	IgnoreZero       bool
}

func NewCurrencyOptions() *CurrencyOptions {
	return &CurrencyOptions{
		Decimals:         2,          // Default value for Decimals
		Currency:         quotes.USD, // Default value for Currency
		IgnoreSmallValue: false,      // Default value for IgnoreSmallValue
		IgnoreMinus:      true,       // Default value for IgnoreMinus
		IgnoreZero:       false,      // Default value for IgnoreZero
	}
}

func formatNumberWithCommas(value float64, decimals int) string {
	// Split the number into integer and fractional parts.
	intPart := int64(value)
	fracPart := value - float64(intPart)

	// Convert the integer part to a string and insert commas.
	intStr := fmt.Sprintf("%d", intPart)
	if intPart >= 1000 || intPart <= -1000 {
		var result strings.Builder
		sign := ""
		if intPart < 0 {
			sign = "-"
			intStr = intStr[1:]
		}
		for i, digit := range intStr {
			if i > 0 && (len(intStr)-i)%3 == 0 {
				result.WriteString(",")
			}
			result.WriteRune(digit)
		}
		intStr = sign + result.String()
	}

	// Format the fractional part.
	fracStr := ""
	if decimals > 0 {
		fracFormat := fmt.Sprintf("%%.%df", decimals)
		fracStr = fmt.Sprintf(fracFormat, fracPart)[1:] // Remove leading "0" from fractional part.
	}

	return intStr + fracStr
}

func PrettifyCurrency(value interface{}, opts *CurrencyOptions) string {
	var floatValue float64
	var err error

	switch v := value.(type) {
	case string:
		floatValue, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return currencyMap[opts.Currency] + Zero
		}
	case float64:
		floatValue = v
	default:
		return currencyMap[opts.Currency] + Zero
	}

	minus := ""
	currencySuffix := ""
	if !opts.IgnoreMinus && floatValue < 0 {
		floatValue = math.Abs(floatValue)
		minus = "-"
	}

	if floatValue == 0 {
		if opts.IgnoreZero {
			return "<" + currencyMap[opts.Currency] + LessThanZero
		} else {
			return currencyMap[opts.Currency] + Zero
		}
	} else if floatValue < 1 {
		if floatValue < 0.01 && opts.IgnoreSmallValue {
			return "<" + currencyMap[opts.Currency] + LessThanZero
		}
	} else if floatValue > 999999999 {
		floatValue /= 1000000000
		currencySuffix = "B"
	} else if floatValue > 999999 {
		floatValue /= 1000000
		currencySuffix = "M"
	}

	expo := math.Pow(10, float64(opts.Decimals))
	floatValue = math.Floor(floatValue*expo) / expo

	currencySymbol := currencyMap[opts.Currency]
	formatter := formatNumberWithCommas(floatValue, opts.Decimals)
	formattedValue := fmt.Sprintf("%s%s%s%s", minus, currencySymbol, formatter, currencySuffix)

	return formattedValue
}
