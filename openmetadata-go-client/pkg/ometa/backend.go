package ometa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Backend interface {
	Call(ctx context.Context, method, path string, body []byte, params url.Values) ([]byte, error)
}

type HTTPBackend struct {
	config     *ClientConfig
	httpClient *http.Client
	baseURL    string
}

func NewHTTPBackend(config *ClientConfig) *HTTPBackend {
	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &HTTPBackend{
		config:     config,
		httpClient: httpClient,
		baseURL:    fmt.Sprintf("%s/api/%s", config.BaseURL, config.APIVersion),
	}
}


func (b *HTTPBackend) Call(ctx context.Context, method, path string, body []byte, params url.Values) ([]byte, error) {
	fullURL := fmt.Sprintf("%s/%s", b.baseURL, path)

	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequestWithContext(ctx, method, fullURL, bytes.NewReader(body))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, fullURL, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("openmetadata: failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if b.config.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.config.AccessToken))
	}
	if method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json-patch+json")
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	rawBody, err := b.dispatchWithRetry(req, body)
	if err != nil {
		return nil, err
	}

	return rawBody, nil
}

func (b *HTTPBackend) dispatchWithRetry(req *http.Request, body []byte) ([]byte, error) {
	retries := b.config.Retry + 1 // Initial attempt + retries
	var resp *http.Response
	var err error

	for retries > 0 {
		if body != nil {
			req.Body = io.NopCloser(bytes.NewReader(body))
		}
		resp, err = b.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("openmetadata: request failed: %w", err)
		}

		if resp.StatusCode >= 400 && b.shouldRetry(resp.StatusCode) {
			_ = resp.Body.Close()
			retries--
			time.Sleep(b.config.RetryWait)
			continue
		}
		break
	}

	defer func() { _ = resp.Body.Close() }()
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("openmetadata: failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode}
		if len(rawBody) > 0 {
			_ = json.Unmarshal(rawBody, apiErr)
		}
		if apiErr.Message == "" {
			apiErr.Message = http.StatusText(resp.StatusCode)
		}
		return nil, apiErr
	}

	if len(rawBody) == 0 {
		return nil, nil
	}

	return rawBody, nil
}

func (b *HTTPBackend) shouldRetry(statusCode int) bool {
	for _, code := range b.config.RetryCodes {
		if statusCode == code {
			return true
		}
	}
	return false
}
