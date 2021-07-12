package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

var ctx = context.Background()

func Test_CreateChat(t *testing.T) {
	as := assert.New(t)

	t.Run("CreateChat, AddMember, GetMemberList, DeleteMember, DeleteChat", func(t *testing.T) {
		cli := lark.New(
			lark.WithAppCredential(AppAllPermission.AppID, AppAllPermission.AppSecret),
			lark.WithTimeout(time.Second*20),
		)

		chatID := ""
		{
			resp, _, err := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{})
			printData("CreateChat", resp, err)
			as.Nil(err)
			chatID = resp.ChatID
		}

		{
			resp, _, err := cli.Chat.AddChatMember(ctx, &lark.AddChatMemberReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
				IDList:       []string{UserAdmin.UserID},
			})
			printData("AddMember", resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := cli.Chat.GetChatMemberList(ctx, &lark.GetChatMemberListReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
			})
			printData("GetMemberList", resp, err)
			as.Nil(err)
			as.Len(resp.Items, 1)
			as.Equal(UserAdmin.UserID, resp.Items[0].MemberID)
		}

		{
			resp, _, err := cli.Chat.DeleteChatMember(ctx, &lark.DeleteChatMemberReq{
				MemberIDType: lark.IDTypePtr(lark.IDTypeUserID),
				ChatID:       chatID,
				IDList:       []string{UserAdmin.UserID},
			})
			printData("DeleteMember", resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := cli.Chat.DeleteChat(ctx, &lark.DeleteChatReq{
				ChatID: chatID,
			})
			printData("DeleteChat", resp, err)
			as.Nil(err)
		}
	})
}

func Test_Chat_member(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Chat.IsInChat(ctx, &lark.IsInChatReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		printData("IsInChat", resp, err)
		as.Nil(err)
		as.Equal(true, resp.IsInChat)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Chat.IsInChat(ctx, &lark.IsInChatReq{
			ChatID: ChatNotContainALLPermissionApp.ChatID,
		})
		printData("IsInChat", resp, err)
		as.Nil(err)
		as.Equal(false, resp.IsInChat)
	})
}

func Test_GetChat(t *testing.T) {
	as := assert.New(t)

	t.Run("SearchChat, success", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Chat.GetChat(ctx, &lark.GetChatReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Contains(resp.Name, "lark-sdk")
		as.Equal("group", resp.ChatMode)
		as.Equal(lark.ChatTypePrivate, resp.ChatType)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Chat.UpdateChat(ctx, &lark.UpdateChatReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
			Name:   ptr.String("包含「lark-sdk」的群 " + strconv.FormatInt(randInt64(), 10)),
		})
		as.Nil(err)
		as.NotNil(resp)
	})

	t.Run("GetChatListOfSelf, success", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Chat.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{
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
		resp, _, err := AppAllPermission.Ins().Chat.SearchChat(ctx, &lark.SearchChatReq{
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
		resp, _, err := AppAllPermission.Ins().Chat.GetChatAnnouncement(ctx, &lark.GetChatAnnouncementReq{
			ChatID: ChatContainALLPermissionApp.ChatID,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Contains(resp.Content, "群公告")
	})
}
