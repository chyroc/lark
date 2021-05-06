package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt64() int64 {
	return rand.Int63()
}

func mockGetTenantAccessTokenFailed(ctx context.Context) (*lark.TokenExpire, *lark.Response, error) {
	return nil, nil, fmt.Errorf("failed")
}

func Test_Helper(t *testing.T) {
	as := assert.New(t)

	t.Run("GetErrorCode", func(t *testing.T) {
		as.Equal(-1, lark.GetErrorCode(fmt.Errorf("x")))
	})

	t.Run("UnwrapMessageContent", func(t *testing.T) {
		t.Run("image", func(t *testing.T) {
			_, err := lark.UnwrapMessageContent(lark.MsgTypeShareChat, `{"image_key":"image-x"}`)
			as.NotNil(err)
			as.Contains(err.Error(), "unknown message type")
		})

		t.Run("image", func(t *testing.T) {
			_, err := lark.UnwrapMessageContent(lark.MsgTypeText, "")
			as.NotNil(err)
			as.Contains(err.Error(), "invalid content")
		})

		t.Run("text", func(t *testing.T) {
			res, err := lark.UnwrapMessageContent(lark.MsgTypeText, `{"text":"hi"}`)
			as.Nil(err)
			as.Equal("hi", res.Text)
		})

		t.Run("image", func(t *testing.T) {
			res, err := lark.UnwrapMessageContent(lark.MsgTypeImage, `{"image_key":"image-x"}`)
			as.Nil(err)
			as.Equal("image-x", res.ImageKey)
		})
	})
}
