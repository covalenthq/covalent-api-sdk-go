package tests

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

func TestBigInt_UnmarshalJSON(t *testing.T) {
	// Test case 1: Valid JSON number
	validJSON := []byte("\"123456789\"")
	b := &utils.BigInt{}
	err := b.UnmarshalJSON(validJSON)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := big.NewInt(123456789)
	if !reflect.DeepEqual(b.Int, expected) {
		t.Errorf("Unexpected value. Got: %v, Expected: %v", b.Int, expected)
	}

	// Test case 2: Invalid JSON number
	invalidJSON := []byte("invalid")
	b = &utils.BigInt{}
	err = b.UnmarshalJSON(invalidJSON)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case 3: Empty JSON number
	emptyJSON := []byte("")
	b = &utils.BigInt{}
	err = b.UnmarshalJSON(emptyJSON)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestBigInt_UnmarshalJSON_NilInt(t *testing.T) {
	// Test case: Initialize with nil Int
	b := &utils.BigInt{}
	err := b.UnmarshalJSON([]byte("\"123\""))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := big.NewInt(123)
	if !reflect.DeepEqual(b.Int, expected) {
		t.Errorf("Unexpected value. Got: %v, Expected: %v", b.Int, expected)
	}
}
