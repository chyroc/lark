package test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	as := assert.New(t)

	as.Equal("TRACE", lark.LogLevelTrace.String())
	as.Equal("DEBUG", lark.LogLevelDebug.String())
	as.Equal("INFO", lark.LogLevelInfo.String())
	as.Equal("WARN", lark.LogLevelWarn.String())
	as.Equal("ERROR", lark.LogLevelError.String())
}
