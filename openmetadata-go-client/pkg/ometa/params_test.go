package ometa_test

import (
	"testing"

	"github.com/open-metadata/openmetadata-sdk/openmetadata-go-client/pkg/ometa"
	"github.com/stretchr/testify/assert"
)

func TestEncodeParams(t *testing.T) {
	type TestParams struct {
		Name  string `form:"name,omitempty"`
		Age   int    `form:"age,omitempty"`
		Email string `form:"email,omitempty"`
	}

	params := ometa.EncodeParams(
		TestParams{
			Name:  "Alice",
			Age:   30,
			Email: "alice@example.com",
		})

	assert.Equal(t, "Alice", params.Get("name"), "name should be Alice")
	assert.Equal(t, "30", params.Get("age"), "age should be 30")
	assert.Equal(t, "alice@example.com", params.Get("email"), "email should be alice@example.com")
}

