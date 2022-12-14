//go:build unit
// +build unit

package errors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotFound(t *testing.T) {
	t.Parallel()

	var e NotFound

	assert.Equal(t, "resource not found", e.Error())
}

func TestErrorInvalidInput(t *testing.T) {
	t.Parallel()

	var e InvalidInput

	assert.Equal(t, "invalid input error", e.Error())
}

func TestErrorUnexpectedStatusCode(t *testing.T) {
	t.Parallel()

	e := NewUnexpectedStatusCode(99, "wat")

	assert.Equal(t, "99 response returned: wat", e.Error())
}

func TestErrorUnauthorized(t *testing.T) {
	t.Parallel()

	e := NewUnauthorizedError()

	assert.Equal(t, 401, e.statusCode)
	assert.True(t, strings.Contains(e.Error(), "Invalid credentials provided"))
}
