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

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

// ExampleChat ...
func ExampleChat() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// create chat
	{
		resp, _, err := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
			Name: ptr.String("<CHAT_NAME>"),
		})
		fmt.Println(resp, err)
	}

	// update chat
	{
		resp, _, err := cli.Chat.UpdateChat(ctx, &lark.UpdateChatReq{
			ChatID: "<CHAT_ID>",
			Name:   ptr.String("<CHAT_NAME>"),
		})
		fmt.Println(resp, err)
	}

	// delete chat
	{
		resp, _, err := cli.Chat.DeleteChat(ctx, &lark.DeleteChatReq{
			ChatID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// get chat
	{
		resp, _, err := cli.Chat.GetChat(ctx, &lark.GetChatReq{
			ChatID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// get chat of self
	{
		resp, _, err := cli.Chat.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
		fmt.Println(resp, err)
	}

	// add member
	{
		resp, _, err := cli.Chat.AddChatMember(ctx, &lark.AddChatMemberReq{
			ChatID:       "<CHAT_ID>",
			MemberIDType: lark.IDTypePtr(lark.IDTypeOpenID),
			IDList:       []string{"<USER_OPEN_ID>"},
		})
		fmt.Println(resp, err)
	}

	// delete member
	{
		resp, _, err := cli.Chat.DeleteChatMember(ctx, &lark.DeleteChatMemberReq{
			ChatID:       "<CHAT_ID>",
			MemberIDType: lark.IDTypePtr(lark.IDTypeOpenID),
			IDList:       []string{"<USER_OPEN_ID>"},
		})
		fmt.Println(resp, err)
	}

	// check bot is in member
	{
		resp, _, err := cli.Chat.IsInChat(ctx, &lark.IsInChatReq{
			ChatID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// search chat
	{
		resp, _, err := cli.Chat.SearchChat(ctx, &lark.SearchChatReq{
			Query: ptr.String("<SEARCH>"),
		})
		fmt.Println(resp, err)
	}
}
