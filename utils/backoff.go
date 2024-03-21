package utils

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

const DefaultBackoffMaxRetries = 5
const BaseDelayMs = 1000 // Base delay in milliseconds

type ExponentialBackoff struct {
	RetryCount int
	APIKey     string
	Debug      bool
	MaxRetries int
	UserAgent  string
}

func NewExponentialBackoff(apiKey string, debug bool, maxRetries int, userAgent string) *ExponentialBackoff {
	if maxRetries == 0 {
		maxRetries = DefaultBackoffMaxRetries
	}
	return &ExponentialBackoff{
		APIKey:     apiKey,
		Debug:      debug,
		MaxRetries: maxRetries,
		RetryCount: 1,
		UserAgent:  userAgent,
	}
}

func (e *ExponentialBackoff) BackOff(url string) (*http.Response, error) {
	var startTime time.Time
	if e.Debug {
		startTime = time.Now()
	}

	response, err := e.makeRequest(url)
	if err != nil {
		return nil, err // Return the error if request fails
	}

	defer DebugOutput(url, response.StatusCode, startTime)

	// Check for rate limiting or other errors
	if response.StatusCode == http.StatusTooManyRequests || response.StatusCode < 200 || response.StatusCode >= 300 {
		response.Body.Close()
		if e.RetryCount < e.MaxRetries {
			e.RetryCount++
			delayMs := time.Duration(math.Pow(2, float64(e.RetryCount)) * float64(BaseDelayMs))
			time.Sleep(delayMs * time.Millisecond)
			return e.BackOff(url) // Retry the request
		} else {
			return nil, fmt.Errorf("max retries exceeded: %d", e.MaxRetries)
		}
	}

	return response, nil

}

func (e *ExponentialBackoff) makeRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+e.APIKey)
	req.Header.Set("X-Requested-With", e.UserAgent)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (e *ExponentialBackoff) SetNumAttempts(retryCount int) {
	e.RetryCount = retryCount
}
