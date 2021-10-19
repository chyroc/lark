package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	s := RandString(10)
	t.Logf(s)
	assert.NotEmpty(t, s)
	assert.Len(t, s, 10)
}
