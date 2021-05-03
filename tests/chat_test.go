package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

var ctx = context.Background()

func Test_CreateChat(t *testing.T) {
	as := assert.New(t)

	t.Run("CreateChat, no-permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Chat().CreateChat(ctx, &lark.CreateChatReq{})
		spew.Dump(err)
		as.NotNil(err)
		as.Equal(99991672, lark.GetErrorCode(err))
		as.Contains(err.Error(), "No permission")
	})

	t.Run("CreateChat, AddMember, GetMember, DeleteMember, DeleteChat", func(t *testing.T) {
		cli := lark.New(AppALLPermission.AppID, AppALLPermission.AppSecret)

		chatID := ""
		{
			resp, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{})
			spew.Dump("CreateChat", resp, err)
			as.Nil(err)
			chatID = resp.ChatID
		}

		{
			resp, _, err := cli.Chat().AddMember(ctx, &lark.AddMemberReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
				IDList:       []string{UserAdmin.UserID},
			})
			spew.Dump("AddMember", resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := cli.Chat().GetMember(ctx, &lark.GetMemberReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
			})
			spew.Dump("GetMember", resp, err)
			as.Nil(err)
			as.Len(resp.Items, 1)
			as.Equal(UserAdmin.UserID, resp.Items[0].MemberID)
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

func Test_Chat_member(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().IsInChat(ctx, &lark.IsInChatReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		spew.Dump("IsInChat", resp, err)
		as.Nil(err)
		as.Equal(true, resp.IsInChat)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().IsInChat(ctx, &lark.IsInChatReq{
			ChatID: ChatNotContainALLPermissionApp.ChatID,
		})
		spew.Dump("IsInChat", resp, err)
		as.Nil(err)
		as.Equal(false, resp.IsInChat)
	})
}

func Test_GetChat(t *testing.T) {
	as := assert.New(t)

	t.Run("GetChatListOfSelf, no-permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Chat().GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("GetChatListBySearch, no-permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Chat().GetChatListBySearch(ctx, &lark.GetChatListBySearchReq{})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("GetChatListOfSelf, success", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{
			UserIDType: lark.IDTypePtr(lark.IDTypeOpenID),
		})
		as.Nil(err)
		as.True(len(resp.Items) > 0)
		containThisChat := true
		for _, v := range resp.Items {
			if v.ChatID == ChatContainALLPermissionApp.ChatID {
				containThisChat = true
			}
		}
		as.True(containThisChat, fmt.Sprintf("shou contain chat: %s: %#v", ChatContainALLPermissionApp.ChatID, resp.Items))
	})

	t.Run("GetChatListBySearch, success", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().GetChatListBySearch(ctx, &lark.GetChatListBySearchReq{
			Query: ptrString("lark-sdk"),
		})
		as.Nil(err)
		as.True(len(resp.Items) > 0)
		containThisChat := true
		for _, v := range resp.Items {
			if v.ChatID == ChatContainALLPermissionApp.ChatID {
				containThisChat = true
			}
		}
		as.True(containThisChat, fmt.Sprintf("shou contain chat: %s: %#v", ChatContainALLPermissionApp.ChatID, resp.Items))
	})
}

func Test_ChatAnnouncement(t *testing.T) {
	as := assert.New(t)

	t.Run("GetAnnouncement, no-permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Chat().GetAnnouncement(ctx, &lark.GetAnnouncementReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		spew.Dump(err)
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("GetAnnouncement, all-permission", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().GetAnnouncement(ctx, &lark.GetAnnouncementReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Contains(resp.Content, "群公告")
	})
}
