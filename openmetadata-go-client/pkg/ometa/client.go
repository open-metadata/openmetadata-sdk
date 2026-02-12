package ometa

import "time"

type Client struct {
	config  *ClientConfig
	backend Backend
	Table *TableService
}

func NewClient(
	baseUrl string,
	opts...ClientOption) *Client {
		config := defaultConfig(baseUrl)
		for _, opt := range opts {
			opt(config)
		}

		backend := NewHTTPBackend(config)

		c := &Client{
			config:  config,
			backend: backend,
		}
		c.Table = &TableService{backend: backend}
		return c
	}

func defaultConfig(baseURL string) *ClientConfig {
    return &ClientConfig{
        BaseURL:    baseURL,
        APIVersion: "v1",
        Retry:      3,
        RetryWait:  30 * time.Second,
        RetryCodes: []int{429, 504},
    }
}
