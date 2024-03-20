package utils

import (
	"fmt"
	"time"
)

func DebugOutput(url string, responseStatus int, startTime time.Time) {
	// Check if startTime is the zero value, which means it was not initialized
	if startTime.IsZero() {
		return
	}
	// Calculate elapsed time
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	// Log the details
	// Note: Sprintf is used for string formatting, and Printf is used for printing
	fmt.Printf("[DEBUG] | Request URL: %s | Response code: %d | Response time: %.2fms\n",
		url, responseStatus, elapsedTime.Seconds()*1000)
}
