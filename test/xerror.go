package test

import (
	"github.com/linyuanbin/fmt-conv/xerror"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertError validates that given error
// is of an instance of xerror.Error.
// If so, returns the instance of xerror.Error.
func AssertError(t *testing.T, err error) *xerror.Error {
	assert.NotNil(t, err)
	standardized, ok := err.(*xerror.Error)
	assert.Equal(t, true, ok)
	return standardized
}
