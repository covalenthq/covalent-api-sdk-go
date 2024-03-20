package tests

import (
	"testing"
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	// Test case 1: Valid JSON string
	validJSON := []byte("\"2023-11-23\"")
	ct := &utils.CustomTime{}
	err := ct.UnmarshalJSON(validJSON)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC)
	if !ct.Time.Equal(expected) {
		t.Errorf("Unexpected value. Got: %v, Expected: %v", ct.Time, expected)
	}

	// Test case 2: Empty JSON string
	emptyJSON := []byte("\"\"")
	ct = &utils.CustomTime{}
	err = ct.UnmarshalJSON(emptyJSON)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ct.Time != nil {
		t.Errorf("Expected nil time, got: %v", ct.Time)
	}

	// Test case 3: Null JSON string
	nullJSON := []byte("null")
	ct = &utils.CustomTime{}
	err = ct.UnmarshalJSON(nullJSON)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ct.Time != nil {
		t.Errorf("Expected nil time, got: %v", ct.Time)
	}

	// Test case 4: Invalid JSON string
	invalidJSON := []byte("\"invalid\"")
	ct = &utils.CustomTime{}
	err = ct.UnmarshalJSON(invalidJSON)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
