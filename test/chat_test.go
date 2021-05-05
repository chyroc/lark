package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/chyroc/go-ptr"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

var ctx = context.Background()

func Test_ChatFailed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		chatCli := cli.Chat()

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetAnnouncement(ctx, &lark.GetAnnouncementReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.UpdateAnnouncement(ctx, &lark.UpdateAnnouncementReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.CreateChat(ctx, &lark.CreateChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.DeleteChat(ctx, &lark.DeleteChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetChat(ctx, &lark.GetChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.AddMember(ctx, &lark.AddMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.DeleteMember(ctx, &lark.DeleteMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetMemberList(ctx, &lark.GetMemberListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.IsInChat(ctx, &lark.IsInChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.SearchChat(ctx, &lark.SearchChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.UpdateChat(ctx, &lark.UpdateChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		chatCli := cli.Chat()

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetAnnouncement(ctx, &lark.GetAnnouncementReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.UpdateAnnouncement(ctx, &lark.UpdateAnnouncementReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.CreateChat(ctx, &lark.CreateChatReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.DeleteChat(ctx, &lark.DeleteChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetChat(ctx, &lark.GetChatReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.AddMember(ctx, &lark.AddMemberReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.DeleteMember(ctx, &lark.DeleteMemberReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.GetMemberList(ctx, &lark.GetMemberListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.IsInChat(ctx, &lark.IsInChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.SearchChat(ctx, &lark.SearchChatReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.UpdateChat(ctx, &lark.UpdateChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_CreateChat(t *testing.T) {
	as := assert.New(t)

	t.Run("CreateChat, AddMember, GetMemberList, DeleteMember, DeleteChat", func(t *testing.T) {
		cli := lark.New(
			lark.WithAppCredential(AppALLPermission.AppID, AppALLPermission.AppSecret),
			lark.WithTimeout(time.Second*10),
		)

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
			resp, _, err := cli.Chat().GetMemberList(ctx, &lark.GetMemberListReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
			})
			spew.Dump("GetMemberList", resp, err)
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

	t.Run("SearchChat, success", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().GetChat(ctx, &lark.GetChatReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Contains(resp.Name, "lark-sdk")
		as.Equal("group", resp.ChatMode)
		as.Equal(lark.ChatTypePrivate, resp.ChatType)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().UpdateChat(ctx, &lark.UpdateChatReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
			Name:   ptr.String("包含「lark-sdk」的群 " + strconv.FormatInt(randInt64(), 10)),
		})
		as.Nil(err)
		as.NotNil(resp)
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

	t.Run("SearchChat, success", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Chat().SearchChat(ctx, &lark.SearchChatReq{
			Query: ptr.String("lark-sdk"),
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
