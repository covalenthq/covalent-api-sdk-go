package utils

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func PaginateEndpoint(urlStr, apiKey string, urlParams url.Values, page int, debug bool, threadCount int) (*http.Response, error) {

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Check if "page-number" parameter exists and get its value
	if urlParams.Has("page-number") {
		urlParams.Set("page-number", strconv.Itoa(page))
	} else {
		urlParams.Add("page-number", strconv.Itoa(page))
	}

	parsedURL.RawQuery = urlParams.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", `Bearer `+apiKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	backoff := NewExponentialBackoff(apiKey, debug, 0)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		return resp, nil
	}
}

func PaginateEndpointUsingLinks(urlStr, apiKey string, urlParams url.Values, debug bool, threadCount int) (*http.Response, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	parsedURL.RawQuery = urlParams.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", `Bearer `+apiKey)

	var startTime time.Time // Declares startTime, initially set to zero value of time.Time

	if debug {
		startTime = time.Now() // Initialize startTime with the current time
	}

	backoff := NewExponentialBackoff(apiKey, debug, 0)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	DebugOutput(resp.Request.URL.String(), resp.StatusCode, startTime)

	if resp.StatusCode == 429 {
		res, err := backoff.BackOff(resp.Request.URL.String())
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		return resp, nil
	}
}
