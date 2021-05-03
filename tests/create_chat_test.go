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

		chatID := ""
		{
			resp, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{})
			spew.Dump("CreateChat", resp, err)
			as.Nil(err)
			chatID = resp.ChatID
		}

		{
			resp, _, err := cli.Chat().AddMember(ctx, &lark.AddMemberReq{
				// MemberIDType: lark.IDTypeUserID,
				ChatID: chatID,
				IDList: []string{UserAdmin.UserID},
			})
			spew.Dump("AddMember", resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := cli.Chat().DeleteMember(ctx, &lark.DeleteMemberReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
				IDList:       []string{UserAdmin.UserID},
			})
			spew.Dump("DeleteMember", resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := cli.Chat().DeleteChat(ctx, &lark.DeleteChatReq{
				ChatID: chatID,
			})
			spew.Dump("DeleteChat", resp, err)
			as.Nil(err)
		}
	})
}
