package utils

import (
	"strings"
	"time"
)

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
