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

// ExampleOtherMessage ...
func ExampleOtherMessage() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// delete message
	{
		resp, _, err := cli.Message.DeleteMessage(ctx, &lark.DeleteMessageReq{
			MessageID: "<MESSAGE_ID>",
		})
		fmt.Println(resp, err)
	}

	// get message
	{
		resp, _, err := cli.Message.GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "<MESSAGE_ID>",
		})
		fmt.Println(resp, err)
	}

	// get message file
	{
		resp, _, err := cli.Message.GetMessageFile(ctx, &lark.GetMessageFileReq{
			Type:      "image",
			MessageID: "<MESSAGE_ID>",
			FileKey:   "<FILE_KEY>",
		})
		fmt.Println(resp, err)
	}

	// get message list
	{
		resp, _, err := cli.Message.GetMessageList(ctx, &lark.GetMessageListReq{
			ContainerIDType: lark.ContainerIDTypeChat,
			ContainerID:     "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// get message read users
	{
		resp, _, err := cli.Message.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
			MessageID: "<MESSAGE_ID>",
		})
		fmt.Println(resp, err)
	}

	{
		resp, _, err := cli.Message.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
			MessageID: "<MESSAGE_ID>",
			Content:   `{"text":"reply"}`,
			MsgType:   lark.MsgTypeText,
		})
		fmt.Println(resp, err)
	}
	// update card message
	{
		resp, _, err := cli.Message.UpdateMessage(ctx, &lark.UpdateMessageReq{
			MessageID: "<MESSAGE_ID>",
			Content:   "<CARD>",
		})
		fmt.Println(resp, err)
	}
}
