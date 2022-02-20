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
package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

// ExampleSendMessage ...
func ExampleSendMessage() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// send raw message
	{
		resp, _, err := cli.Message.SendRawMessage(ctx, &lark.SendRawMessageReq{
			ReceiveIDType: lark.IDTypeEmail,
			ReceiveID:     "<EMAIL>",
			Content:       `{"text":"text"}`,
			MsgType:       lark.MsgTypeText,
		})
		fmt.Println(resp, err)
	}

	// send text message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendText(ctx, "<TEXT>")
		fmt.Println(resp, err)
	}

	// send image message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendImage(ctx, "<IMAGE_KEY>")
		fmt.Println(resp, err)
	}

	// send post message
	{
		// send post message with string
		{
			resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendPost(ctx, "<POST_MESSAGE_CONTENT>")
			fmt.Println(resp, err)
		}

		// send post message with struct
		{
			data := &lark.MessageContentPostAll{
				ZhCn: &lark.MessageContentPost{
					Title: "<TITLE>",
					Content: [][]lark.MessageContentPostItem{
						{
							lark.MessageContentPostText{Text: "<TEXT>"},
						},
						{
							lark.MessageContentPostAt{UserID: "<OPEN_ID>"},
						},
					},
				},
				JaJp: nil,
				EnUs: nil,
			}
			resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendPost(ctx, data.String())
			fmt.Println(resp, err)
		}
	}

	// send card message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendCard(ctx, "<CARD_MESSAGE_CONTENT>")
		fmt.Println(resp, err)
	}

	// send share chat message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendShareChat(ctx, "<CHAT_ID>")
		fmt.Println(resp, err)
	}

	// send share user message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendShareUser(ctx, "<USER_ID>")
		fmt.Println(resp, err)
	}

	// send audio message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendAudio(ctx, "<FILE_KEY>")
		fmt.Println(resp, err)
	}

	// send media message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendMedia(ctx, "<IMAGE_KEY>", "<FILE_KEY>")
		fmt.Println(resp, err)
	}

	// send file message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendFile(ctx, "<FILE_KEY>")
		fmt.Println(resp, err)
	}

	// send sticker message
	{
		resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendSticker(ctx, "<FILE_KEY>")
		fmt.Println(resp, err)
	}
}
