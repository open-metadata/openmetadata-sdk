package ometa

import (
	"net/http"
	"time"
)

type ClientConfig struct {
    BaseURL string
	APIVersion string
	AccessToken string
    Retry int
	RetryWait time.Duration
	RetryCodes []int
    HTTPClient *http.Client
}

type ClientOption func(*ClientConfig)

func WithToken(token string) ClientOption { 
	return func(c *ClientConfig) {
		c.AccessToken = token
	}
}

func WithAPIVersion(v string) ClientOption { 
	return func(c *ClientConfig) {
		c.APIVersion = v
	}
}

func WithRetry(count int, wait time.Duration) ClientOption { 
	return func(c *ClientConfig) {
		c.Retry = count
		c.RetryWait = wait
	}
}

func WithHTTPClient(c *http.Client) ClientOption { 
	return func(cfg *ClientConfig) {
		cfg.HTTPClient = c
	}
 }