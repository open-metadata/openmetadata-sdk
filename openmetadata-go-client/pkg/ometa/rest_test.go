package ometa

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetHeader(t *testing.T) {
	expectedHeader := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer myFakeAccessToken"},
	}

	header := setHeader(nil, "Bearer myFakeAccessToken")

	assert.Equal(t, expectedHeader, header, "header should be equal")
}

func TestSetCustomHeader(t *testing.T) {
	customHeader := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer myFakeAccessToken"},
		"X-Test":        {"test"},
		"X-Test2":       {"test2"},
	}

	header := setHeader(customHeader, "Bearer myFakeAccessToken")

	assert.Equal(t, customHeader, header, "header should be equal")
}

func TestSetExtraHeader(t *testing.T) {
	expectedHeader := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer myFakeAccessToken"},
		"X-Test":        {"test"},
		"X-Test2":       {"test2"},
	}

	header := setHeader(nil, "Bearer myFakeAccessToken")
	var extraHeader map[string][]string
	extraHeader = make(map[string][]string)
	extraHeader["X-Test"] = []string{"test"}
	extraHeader["X-Test2"] = []string{"test2"}
	setExtraHeader(extraHeader, header)

	assert.Equal(t, expectedHeader, header, "header should be equal")
}

func TestSetQueryParameter(t *testing.T) {
	restConfig := NewRestConfig(
		"http://localhost:8080",
		"",
		0,
		0,
		nil,
		"myFakeAccessToken",
	)
	rest := NewRest(restConfig)
	rest.req, _ = http.NewRequest("GET", "http://localhost:8080", nil)

	queryParams := make(map[string]string)
	queryParams["limit"] = "10"
	queryParams["deleted"] = "false"

	rest.setQueryParams(queryParams)
	assert.Equal(t, rest.req.URL.RawQuery, "deleted=false&limit=10", "query parameters should be equal")
}

func TestNewRestSetsDefaults(t *testing.T) {
	restConfig := &RestConfig{
		BaseURL: "http://localhost:8080",
	}
	rest := NewRest(restConfig)

	assert.NotNil(t, rest.restConfig, "config should not be nil")
	assert.Equal(t, "http://localhost:8080", rest.restConfig.BaseURL, "BaseURL should be set")
	assert.NotZero(t, rest.restConfig.APIVersion, "API version should have a default value")
	assert.NotZero(t, rest.restConfig.Retry, "Retry should have a default value")
	assert.NotZero(t, rest.restConfig.RetryWait, "RetryWait should have a default value")
	assert.NotEmpty(t, rest.restConfig.RetryCodes, "RetryCodes should have a default")
}
