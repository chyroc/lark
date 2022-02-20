/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package test

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_GetMessage(t *testing.T) {
	as := assert.New(t)

	msgIDs := []string{}
	defer func() {
		for _, v := range msgIDs {
			_, _, _ = AppAllPermission.Ins().Message.DeleteMessage(ctx, &lark.DeleteMessageReq{MessageID: v})
		}
	}()

	t.Run("", func(t *testing.T) {
		AppAllPermission.Ins().Message.Send().ToChatID(ChatForSendMessage.ChatID).SendCard(ctx, (&lark.MessageContentCard{
			Header: nil,
			Config: nil,
			Modules: []lark.MessageContentCardModule{
				&lark.MessageContentCardModuleDIV{
					Text: &lark.MessageContentCardObjectText{
						Tag:     lark.MessageContentCardTextTypeLarkMd,
						Content: lark.MdBuilder.AtAll(),
						Lines:   0,
					},
					Fields: nil,
					Extra:  nil,
				},
			},
		}).String())
	})

	t.Run("send-message", func(t *testing.T) {
		t.Run("raw", func(t *testing.T) {
			messageID := ""
			{
				resp, _, err := AppAllPermission.Ins().Message.SendRawMessage(ctx, &lark.SendRawMessageReq{
					ReceiveIDType: lark.IDTypeChatID,
					ReceiveID:     ChatForSendMessage.ChatID,
					Content:       fmt.Sprintf(`{"text":"%d"}`, time.Now().Unix()),
					MsgType:       lark.MsgTypeText,
				})
				as.Nil(err)
				messageID = resp.MessageID
			}

			{
				resp, _, err := AppAllPermission.Ins().Message.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
					MessageID: messageID,
					Content:   fmt.Sprintf(`{"text":"reply-%d"}`, time.Now().Unix()),
					MsgType:   lark.MsgTypeText,
				})
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			}

			{
				_, _, err := AppAllPermission.Ins().Message.DeleteMessage(ctx, &lark.DeleteMessageReq{
					MessageID: messageID,
				})
				as.Nil(err)
			}
		})

		t.Run("", func(t *testing.T) {
			messageID := ""
			{
				resp, _, err := AppAllPermission.Ins().Message.Send().ToChatID(ChatForSendMessage.ChatID).SendText(ctx, strconv.FormatInt(time.Now().Unix(), 10))
				as.Nil(err)
				messageID = resp.MessageID
				msgIDs = append(msgIDs, messageID)
			}

			t.Run("text", func(t *testing.T) {
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendText(ctx, strconv.FormatInt(time.Now().Unix(), 10))
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})

			// 这个图，竟然报：The content of the message contains sensitive information.，暂时不测这个
			t.Run("image", func(t *testing.T) {
				// _, _, err := AppAllPermission.Ins().Message().Reply(messageID).SendImage(ctx, File1.Key)
				// as.Nil(err)
				// msgIDs= append(msgIDs, resp.MessageID)
			})

			t.Run("post-1", func(t *testing.T) {
				s := `{"zh_cn": {"title": "我是一个标题","content": [[{"tag": "text","text": "文本"}]]}}`
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendPost(ctx, s)
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})

			t.Run("post-2", func(t *testing.T) {
				s2 := lark.MessageContentPostAll{
					ZhCn: &lark.MessageContentPost{
						Title: "标题",
						Content: [][]lark.MessageContentPostItem{
							{
								lark.MessageContentPostText{
									Text: "文本",
								},
							},
							{
								lark.MessageContentPostAt{
									UserID: UserAdmin.OpenID,
								},
							},
						},
					},
				}
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendPost(ctx, s2.String())
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})

			t.Run("card", func(t *testing.T) {
				s := `{"config": { "wide_screen_mode": true },"i18n_elements": {"zh_cn": [{"tag": "div","text": { "tag": "lark_md", "content": "文本"}}]}}`
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendCard(ctx, s)
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})

			t.Run("file", func(t *testing.T) {
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendFile(ctx, File2.Key)
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})

			t.Run("chat", func(t *testing.T) {
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendShareChat(ctx, ChatForSendMessage.ChatID)
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})

			t.Run("user", func(t *testing.T) {
				resp, _, err := AppAllPermission.Ins().Message.Reply(messageID).SendShareUser(ctx, UserAdmin.OpenID)
				as.Nil(err)
				msgIDs = append(msgIDs, resp.MessageID)
			})
		})
	})

	t.Run("get-message-read", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Message.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
			UserIDType: lark.IDTypeUserID,
			MessageID:  MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		})
		printData(resp, err)
		as.NotNil(err)
		// as.Contains(err.Error(), "Bot is NOT the sender of the message")
		as.Contains(err.Error(), "these ids not existed")
	})

	t.Run("get-message-read", func(t *testing.T) {
		// resp, _, err := AppAllPermission.Ins().Message().GetMessageRead(ctx, &lark.GetMessageReadReq{
		// 	UserIDType: lark.IDTypeUserID,
		// 	MessageID: MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		// })
		// printData(resp, err)
		// as.NotNil(err)
		// as.Contains(err.Error(), "Bot is NOT the sender of the message")
	})

	t.Run("get-message-text", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Message.GetMessage(ctx, &lark.GetMessageReq{
			MessageID: MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Len(resp.Items, 1)
		as.Equal(lark.MsgTypeText, resp.Items[0].MsgType)
		as.Equal(MessageAdminSendTextInChatContainAllPermissionApp.ChatID, resp.Items[0].ChatID)
		msgContent, err := lark.UnwrapMessageContent(resp.Items[0].MsgType, resp.Items[0].Body.Content)
		as.Nil(err)
		as.Equal("test", msgContent.Text.Text)
	})

	t.Run("get-message-image", func(t *testing.T) {
		messageFile := ""
		{
			resp, _, err := AppAllPermission.Ins().Message.GetMessage(ctx, &lark.GetMessageReq{
				MessageID: MessageAdminSendImageInChatContainAllPermissionApp.MessageID,
			})
			printData(resp, err)
			as.Nil(err)
			as.NotNil(resp)
			as.Len(resp.Items, 1)
			as.Equal(lark.MsgTypeImage, resp.Items[0].MsgType)
			as.Equal(MessageAdminSendImageInChatContainAllPermissionApp.ChatID, resp.Items[0].ChatID)
			as.Contains(resp.Items[0].Body.Content, "image_key")
			msgContent, err := lark.UnwrapMessageContent(resp.Items[0].MsgType, resp.Items[0].Body.Content)
			as.Nil(err)
			messageFile = msgContent.Image.ImageKey
		}

		{
			resp, _, err := AppAllPermission.Ins().Message.GetMessageFile(ctx, &lark.GetMessageFileReq{
				Type:      "image",
				MessageID: MessageAdminSendImageInChatContainAllPermissionApp.MessageID,
				FileKey:   messageFile,
			})
			as.Nil(err)
			as.NotNil(resp)
			bs, err := ioutil.ReadAll(resp.File)
			as.Nil(err)
			as.NotEmpty(bs)
		}
	})

	t.Run("get-message-list", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Message.GetMessageList(ctx, &lark.GetMessageListReq{
			ContainerIDType: lark.ContainerIDTypeChat,
			ContainerID:     ChatContainALLPermissionApp.ChatID,
			StartTime:       nil,
			EndTime:         nil,
			PageToken:       nil,
			PageSize:        nil,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.True(len(resp.Items) > 0)
	})
}

func Test_SendCustomBotMessage(t *testing.T) {
	as := assert.New(t)

	t.Run("AppCustomBotNoValid", func(t *testing.T) {
		cli := AppCustomBotNoValid.CustomBot().Message.Send()
		t.Run("text", func(t *testing.T) {
			_, _, err := cli.SendText(ctx, "hi")
			as.Nil(err)
		})

		t.Run("post", func(t *testing.T) {
			post := lark.MessageContentPostAll{
				ZhCn: &lark.MessageContentPost{
					Title: "title",
					Content: [][]lark.MessageContentPostItem{
						{
							lark.MessageContentPostText{
								Text:     "text",
								UnEscape: false,
							},
						},
					},
				},
			}
			_, _, err := cli.SendPost(ctx, post.String())
			as.Nil(err)
		})

		t.Run("chat", func(t *testing.T) {
			_, _, err := cli.SendShareChat(ctx, ChatForSendMessage.ChatID)
			as.Nil(err)
		})

		t.Run("image", func(t *testing.T) {
			imageKey := "img_v2_094a8a5c-ae93-4602-9416-4d875cc9a96g"
			_, _, err := cli.SendImage(ctx, imageKey)
			as.Nil(err)
		})

		t.Run("card", func(t *testing.T) {
			card := lark.MessageContentCard{
				Header: &lark.MessageContentCardHeader{
					Template: "",
					Title: &lark.MessageContentCardObjectText{
						Tag:     "plain_text",
						Content: "1",
					},
				},
				Config: &lark.MessageContentCardConfig{
					EnableForward: true,
				},
				Modules: []lark.MessageContentCardModule{
					lark.MessageContentCardModuleDIV{
						Text: &lark.MessageContentCardObjectText{
							Tag:     "plain_text",
							Content: "1",
						},
						Fields: nil,
						Extra:  nil,
					},
				},
			}
			_, _, err := cli.SendCard(ctx, card.String())
			as.Nil(err)
		})
	})

	t.Run("AppCustomBotCheckCanSendWord", func(t *testing.T) {
		cli := AppCustomBotCheckCanSendWord.CustomBot().Message.Send()
		t.Run("text", func(t *testing.T) {
			_, _, err := cli.SendText(ctx, "hi")
			as.NotNil(err)
			as.Contains(err.Error(), "Key Words Not Found")
		})

		t.Run("text", func(t *testing.T) {
			_, _, err := cli.SendText(ctx, "hi, can-send")
			as.Nil(err)
		})

		t.Run("post", func(t *testing.T) {
			post := lark.MessageContentPostAll{
				ZhCn: &lark.MessageContentPost{
					Title: "title",
					Content: [][]lark.MessageContentPostItem{
						{
							lark.MessageContentPostText{
								Text:     "title",
								UnEscape: false,
							},
						},
					},
				},
			}
			_, _, err := cli.SendPost(ctx, post.String())
			as.NotNil(err)
			as.Contains(err.Error(), "Key Words Not Found")
		})

		t.Run("post", func(t *testing.T) {
			post := lark.MessageContentPostAll{
				ZhCn: &lark.MessageContentPost{
					Title: "title",
					Content: [][]lark.MessageContentPostItem{
						{
							lark.MessageContentPostText{
								Text:     "can-send",
								UnEscape: false,
							},
						},
					},
				},
			}
			_, _, err := cli.SendPost(ctx, post.String())
			as.Nil(err)
		})

		t.Run("chat", func(t *testing.T) {
			_, _, err := cli.SendShareChat(ctx, ChatForSendMessage.ChatID)
			as.NotNil(err)
			as.Contains(err.Error(), "Key Words Not Found")
		})

		t.Run("image", func(t *testing.T) {
			imageKey := "img_v2_094a8a5c-ae93-4602-9416-4d875cc9a96g"
			_, _, err := cli.SendImage(ctx, imageKey)
			as.NotNil(err)
			as.Contains(err.Error(), "Key Words Not Found")
		})

		t.Run("card", func(t *testing.T) {
			card := lark.MessageContentCard{
				Header: &lark.MessageContentCardHeader{
					Template: "",
					Title: &lark.MessageContentCardObjectText{
						Tag:     "plain_text",
						Content: "1",
					},
				},
				Config: &lark.MessageContentCardConfig{
					EnableForward: true,
				},
				Modules: []lark.MessageContentCardModule{
					lark.MessageContentCardModuleDIV{
						Text: &lark.MessageContentCardObjectText{
							Tag:     "plain_text",
							Content: "1",
						},
						Fields: nil,
						Extra:  nil,
					},
				},
			}
			_, _, err := cli.SendCard(ctx, card.String())
			as.NotNil(err)
			as.Contains(err.Error(), "Key Words Not Found")
		})

		t.Run("card", func(t *testing.T) {
			card := lark.MessageContentCard{
				Header: &lark.MessageContentCardHeader{
					Template: "",
					Title: &lark.MessageContentCardObjectText{
						Tag:     "plain_text",
						Content: "1",
					},
				},
				Config: &lark.MessageContentCardConfig{
					EnableForward: true,
				},
				Modules: []lark.MessageContentCardModule{
					lark.MessageContentCardModuleDIV{
						Text: &lark.MessageContentCardObjectText{
							Tag:     "plain_text",
							Content: "1, can-send",
						},
						Fields: nil,
						Extra:  nil,
					},
				},
			}
			_, _, err := cli.SendCard(ctx, card.String())
			as.Nil(err)
		})
	})

	t.Run("AppCustomBotCheckSign", func(t *testing.T) {
		cli := AppCustomBotCheckSign.CustomBot().Message.Send()
		t.Run("text", func(t *testing.T) {
			_, _, err := cli.SendText(ctx, "hi")
			as.Nil(err)
		})

		t.Run("post", func(t *testing.T) {
			post := lark.MessageContentPostAll{
				ZhCn: &lark.MessageContentPost{
					Title: "title",
					Content: [][]lark.MessageContentPostItem{
						{
							lark.MessageContentPostText{
								Text:     "text",
								UnEscape: false,
							},
						},
					},
				},
			}
			_, _, err := cli.SendPost(ctx, post.String())
			as.Nil(err)
		})

		t.Run("chat", func(t *testing.T) {
			_, _, err := cli.SendShareChat(ctx, ChatForSendMessage.ChatID)
			as.Nil(err)
		})

		t.Run("image", func(t *testing.T) {
			imageKey := "img_v2_094a8a5c-ae93-4602-9416-4d875cc9a96g"
			_, _, err := cli.SendImage(ctx, imageKey)
			as.Nil(err)
		})

		t.Run("card", func(t *testing.T) {
			card := lark.MessageContentCard{
				Header: &lark.MessageContentCardHeader{
					Template: "",
					Title: &lark.MessageContentCardObjectText{
						Tag:     "plain_text",
						Content: "1",
					},
				},
				Config: &lark.MessageContentCardConfig{
					EnableForward: true,
				},
				Modules: []lark.MessageContentCardModule{
					lark.MessageContentCardModuleDIV{
						Text: &lark.MessageContentCardObjectText{
							Tag:     "plain_text",
							Content: "1",
						},
						Fields: nil,
						Extra:  nil,
					},
				},
			}
			_, _, err := cli.SendCard(ctx, card.String())
			as.Nil(err)
		})
	})
}

func Test_EphemeralMessage(t *testing.T) {
	as := assert.New(t)

	resp, res, err := AppAllPermission.Ins().Message.SendEphemeralMessage(ctx, &lark.SendEphemeralMessageReq{
		ChatID:  ChatForSendMessage.ChatID,
		UserID:  UserAdmin.UserID,
		MsgType: "interactive",
		Card: &lark.MessageContentCard{
			Header: &lark.MessageContentCardHeader{
				Template: lark.MessageContentCardHeaderTemplateBlue,
				Title: &lark.MessageContentCardObjectText{
					Tag:     "lark_md",
					Content: "title",
				},
			},
			Modules: []lark.MessageContentCardModule{
				lark.MessageContentCardModuleDIV{
					Text: &lark.MessageContentCardObjectText{
						Tag:     "lark_md",
						Content: "- text1\n- text2",
					},
				},
				lark.MessageContentCardModuleNote{
					Elements: []lark.MessageContentCardElement{
						lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "hi",
						},
					},
				},
			},
		},
	})
	as.Nil(err)
	as.NotNil(resp)
	as.NotEmpty(resp.MessageID)
	as.NotNil(res)
	as.NotEmpty(res.RequestID)
}

func Test_BatchSend(t *testing.T) {
	as := assert.New(t)

	card := lark.MessageContentCard{
		Header: &lark.MessageContentCardHeader{
			Template: "",
			Title: &lark.MessageContentCardObjectText{
				Tag:     "plain_text",
				Content: "1",
			},
		},
		Config: &lark.MessageContentCardConfig{
			EnableForward: true,
		},
		Modules: []lark.MessageContentCardModule{
			lark.MessageContentCardModuleDIV{
				Text: &lark.MessageContentCardObjectText{
					Tag:     "plain_text",
					Content: "1",
				},
				Fields: nil,
				Extra:  nil,
			},
		},
	}

	resp, _, err := AppAllPermission.Ins().Message.BatchSendOldRawMessage(ctx, &lark.BatchSendOldRawMessageReq{
		MsgType: lark.MsgTypeInteractive,
		Card:    card,
		OpenIDs: []string{UserAdmin.OpenID},
	})
	as.Nil(err)
	as.NotNil(resp)
	as.NotEmpty(resp.MessageID)
	as.Empty(resp.InvalidOpenIDs)
}
