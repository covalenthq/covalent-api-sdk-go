package valuefmt

import (
	"fmt"
	"math/big"
)

func CalculatePrettyBalance(value interface{}, decimals int, roundOff bool, precision int) string {
	var bigDecimalValue *big.Float
	switch v := value.(type) {
	case float64:
		bigDecimalValue = new(big.Float).SetFloat64(v)
	case int:
		bigDecimalValue = new(big.Float).SetInt64(int64(v))
	default:
		return "-1"
	}

	expoValue := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	calculated := new(big.Float).Quo(bigDecimalValue, expoValue)

	if decimals == 0 {
		return fmt.Sprintf("%s", calculated.Text('f', 0))
	}

	decimalFixed := precision
	if precision == 0 {
		decimalFixed = 2
		hundred := new(big.Float).SetFloat64(100)
		if calculated.Cmp(hundred) < 0 {
			decimalFixed = 6
		}
	}

	if roundOff {
		return fmt.Sprintf("%."+fmt.Sprint(decimalFixed)+"f", calculated)
	}

	return fmt.Sprintf("%s", calculated.Text('f', -1))
}
