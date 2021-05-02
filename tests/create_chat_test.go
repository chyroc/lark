package tests

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_CreateChat(t *testing.T) {
	as := assert.New(t)

	t.Run("Chat.CreateChat no-permission", func(t *testing.T) {
		cli := lark.New(AppNoPermission.AppID, AppNoPermission.AppSecret)
		ctx := context.Background()
		_, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{})
		spew.Dump(err)
		as.NotNil(err)
		as.Equal(99991672, lark.GetErrorCode(err))
		as.Equal("request Chat#CreateChat failed: code: 99991672, msg: No permission", err.Error())
	})

	t.Run("Chat.CreateChat no-permission", func(t *testing.T) {
		cli := lark.New(AppALLPermission.AppID, AppALLPermission.AppSecret)
		ctx := context.Background()
		_, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{})
		spew.Dump(err)
		as.NotNil(err)
		as.Equal(99991672, lark.GetErrorCode(err))
		as.Equal("request Chat#CreateChat failed: code: 99991672, msg: No permission", err.Error())
	})
}
