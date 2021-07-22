package test

import (
	"testing"

	"github.com/chyroc/lark/internal"
	"github.com/stretchr/testify/assert"
)

func Test_LarkCallbackSignature(t *testing.T) {
	as := assert.New(t)

	timestamp := "1626947018"
	nonce := "64607134"
	encryptKey := "fake-xx"
	body := []byte(`{fake}`)
	signature := "7db78fdbc2fb617c46e713c3ec3f4b4658d94166682a80434fe557f23afcfd06"
	res := internal.CalculateLarkCallbackSignature(timestamp, nonce, encryptKey, body)
	as.Equal(signature, res)
}
