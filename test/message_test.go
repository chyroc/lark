package test

import (
	"fmt"
	"io"
	"strconv"
	"testing"
	"time"

	"github.com/chyroc/go-ptr"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Message_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		msgCli := cli.Message()

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.DeleteMessage(ctx, &lark.DeleteMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessage(ctx, &lark.GetMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessageFile(ctx, &lark.GetMessageFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessageReadUser(ctx, &lark.GetMessageReadUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessageList(ctx, &lark.GetMessageListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.SendRawMessage(ctx, &lark.SendRawMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.UpdateMessage(ctx, &lark.UpdateMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		msgCli := cli.Message()

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.DeleteMessage(ctx, &lark.DeleteMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessage(ctx, &lark.GetMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessageFile(ctx, &lark.GetMessageFileReq{
				Type:      "",
				MessageID: "x",
				FileKey:   "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessageReadUser(ctx, &lark.GetMessageReadUserReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.GetMessageList(ctx, &lark.GetMessageListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.SendRawMessage(ctx, &lark.SendRawMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := msgCli.UpdateMessage(ctx, &lark.UpdateMessageReq{
				MessageID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})
	})
}

func Test_GetMessage(t *testing.T) {
	as := assert.New(t)

	msgIDs := []string{}
	defer func() {
		for _, v := range msgIDs {
			_, _, _ = AppALLPermission.Ins().Message().DeleteMessage(ctx, &lark.DeleteMessageReq{MessageID: v})
		}
	}()

	t.Run("send-message", func(t *testing.T) {
		t.Run("raw", func(t *testing.T) {
			messageID := ""
			{
				resp, _, err := AppALLPermission.Ins().Message().SendRawMessage(ctx, &lark.SendRawMessageReq{
					ReceiveIDType: lark.IDTypePtr(lark.IDTypeChatID),
					ReceiveID:     ptr.String(ChatForSendMessage.ChatID),
					Content:       fmt.Sprintf(`{"text":"%d"}`, time.Now().Unix()),
					MsgType:       lark.MsgTypeText,
				})
				as.Nil(err)
				messageID = resp.MessageID
			}

			{
				resp, _, err := AppALLPermission.Ins().Message().ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
					MessageID: messageID,
					Content:   fmt.Sprintf(`{"text":"reply-%d"}`, time.Now().Unix()),
					MsgType:   lark.MsgTypeText,
				})
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			}

			{
				_, _, err := AppALLPermission.Ins().Message().DeleteMessage(ctx, &lark.DeleteMessageReq{
					MessageID: messageID,
				})
				as.Nil(err)
			}
		})

		t.Run("text", func(t *testing.T) {
			resp, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendText(ctx, strconv.FormatInt(time.Now().Unix(), 10))
			as.Nil(err)
			msgIDs = append(msgIDs, resp.MessageID)
		})

		// 这个图，竟然报：The content of the message contains sensitive information.，暂时不测这个
		t.Run("image", func(t *testing.T) {
			// _, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendImage(ctx, File1.Key)
			// as.Nil(err)
			// msgIDs= append(msgIDs, resp.MessageID)
		})

		t.Run("post", func(t *testing.T) {
			s := `{"zh_cn": {"title": "我是一个标题","content": [[{"tag": "text","text": "文本"}]]}}`
			resp, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendPost(ctx, s)
			as.Nil(err)
			msgIDs = append(msgIDs, resp.MessageID)
		})

		t.Run("card", func(t *testing.T) {
			s := `{"config": { "wide_screen_mode": true },"i18n_elements": {"zh_cn": [{"tag": "div","text": { "tag": "lark_md", "content": "文本"}}]}}`
			resp, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendCard(ctx, s)
			as.Nil(err)
			msgIDs = append(msgIDs, resp.MessageID)
		})

		t.Run("file", func(t *testing.T) {
			resp, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendFile(ctx, File2.Key)
			as.Nil(err)
			msgIDs = append(msgIDs, resp.MessageID)
		})

		t.Run("chat", func(t *testing.T) {
			resp, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendShareChat(ctx, ChatForSendMessage.ChatID)
			as.Nil(err)
			msgIDs = append(msgIDs, resp.MessageID)
		})

		t.Run("user", func(t *testing.T) {
			resp, _, err := AppALLPermission.Ins().Message().Send().ToChatID(ChatForSendMessage.ChatID).SendShareUser(ctx, UserAdmin.OpenID)
			as.Nil(err)
			msgIDs = append(msgIDs, resp.MessageID)
		})
	})

	t.Run("get-message-read", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessageReadUser(ctx, &lark.GetMessageReadUserReq{
			UserIDType: lark.IDTypeUserID,
			MessageID:  MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		})
		spew.Dump(resp, err)
		as.NotNil(err)
		as.Contains(err.Error(), "Bot is NOT the sender of the message")
	})

	t.Run("get-message-read", func(t *testing.T) {
		// resp, _, err := AppALLPermission.Ins().Message().GetMessageRead(ctx, &lark.GetMessageReadReq{
		// 	UserIDType: lark.IDTypeUserID,
		// 	MessageID: MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		// })
		// spew.Dump(resp, err)
		// as.NotNil(err)
		// as.Contains(err.Error(), "Bot is NOT the sender of the message")
	})

	t.Run("get-message-text", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Len(resp.Items, 1)
		as.Equal(lark.MsgTypeText, resp.Items[0].MsgType)
		as.Equal(MessageAdminSendTextInChatContainAllPermissionApp.ChatID, resp.Items[0].ChatID)
		msgContent, err := lark.UnwrapMessageContent(resp.Items[0])
		as.Nil(err)
		as.Equal("test", msgContent.Text)
	})

	t.Run("get-message-image", func(t *testing.T) {
		messageFile := ""
		{
			resp, _, err := AppALLPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
				MessageID: MessageAdminSendImageInChatContainAllPermissionApp.MessageID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			as.NotNil(resp)
			as.Len(resp.Items, 1)
			as.Equal(lark.MsgTypeImage, resp.Items[0].MsgType)
			as.Equal(MessageAdminSendImageInChatContainAllPermissionApp.ChatID, resp.Items[0].ChatID)
			as.Contains(resp.Items[0].Body.Content, "image_key")
			msgContent, err := lark.UnwrapMessageContent(resp.Items[0])
			as.Nil(err)
			messageFile = msgContent.ImageKey
		}

		{
			resp, _, err := AppALLPermission.Ins().Message().GetMessageFile(ctx, &lark.GetMessageFileReq{
				Type:      "image",
				MessageID: MessageAdminSendImageInChatContainAllPermissionApp.MessageID,
				FileKey:   messageFile,
			})
			as.Nil(err)
			as.NotNil(resp)
			bs, err := io.ReadAll(resp.File)
			as.Nil(err)
			as.NotEmpty(bs)
		}
	})

	t.Run("get-message-list", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessageList(ctx, &lark.GetMessageListReq{
			ContainerIDType: lark.ContainerIDTypeChat,
			ContainerID:     ChatContainALLPermissionApp.ChatID,
			StartTime:       nil,
			EndTime:         nil,
			PageToken:       nil,
			PageSize:        nil,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.True(len(resp.Items) > 0)
	})
}
