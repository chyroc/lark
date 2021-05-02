package tests

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_GetMessage(t *testing.T) {
	as := assert.New(t)

	r := lark.New(AppALLPermission.AppID, AppALLPermission.AppSecret)
	ctx := context.Background()

	t.Run("ids not existed", func(t *testing.T) {
		_, _, err := r.Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "these ids not existed")
	})

	t.Run("ids not existed", func(t *testing.T) {
		r = lark.New(AppNoPermission.AppID, AppNoPermission.AppSecret)
		_, _, err := r.Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("ids not existed", func(t *testing.T) {
		resp, _, err := r.Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		spew.Dump(resp, err)
	})
}
