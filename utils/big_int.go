package utils

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// BigInt is a wrapper around math/big.Int that implements json.Unmarshaler.
type BigInt struct {
	*big.Int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *BigInt) UnmarshalJSON(p []byte) error {
	// Unmarshal the JSON number into a string.
	var s string
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}

	// Initialize the Int if it's nil.
	if b.Int == nil {
		b.Int = new(big.Int)
	}

	// Set the value of Int.
	_, ok := b.Int.SetString(s, 10)
	if !ok {
		return fmt.Errorf("cannot set big.Int value: %v", s)
	}
	return nil
}
