package httpclient

import (
	"incrowd-backend/config"
	"net/http"
	"time"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Do(res *http.Request) (*http.Response, error)
}

// NewHttpClient creates a new HTTP client with specified configuration for connection pooling and timeouts
func NewHttpClient(config config.Http) HTTPClient {
	transport := &http.Transport{
		MaxIdleConns:        config.MaxIdleConns,
		MaxConnsPerHost:     config.MaxConnsPerHost,
		MaxIdleConnsPerHost: config.MaxIdleConnsPerHost,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.TimeoutInSeconds) * time.Second,
	}
}
