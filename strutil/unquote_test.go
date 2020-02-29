package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnquoteBytes(t *testing.T) {
	assert.Equal(t, "test", UnquoteBytes([]byte("test")))
	assert.Equal(t, "test", UnquoteBytes([]byte("\"test\"")))
	assert.Equal(t, "\"test\"", UnquoteBytes([]byte("\"\"test\"\"")))
}
