package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

func ExampleChat() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// create chat
	{
		resp, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{
			Name: ptr.String("<CHAT_NAME>"),
		})
		fmt.Println(resp, err)
	}

	// update chat
	{
		resp, _, err := cli.Chat().UpdateChat(ctx, &lark.UpdateChatReq{
			ChatID: "<CHAT_ID>",
			Name:   ptr.String("<CHAT_NAME>"),
		})
		fmt.Println(resp, err)
	}

	// delete chat
	{
		resp, _, err := cli.Chat().DeleteChat(ctx, &lark.DeleteChatReq{
			ChatID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// get chat
	{
		resp, _, err := cli.Chat().GetChat(ctx, &lark.GetChatReq{
			ChatID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// get chat of self
	{
		resp, _, err := cli.Chat().GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
		fmt.Println(resp, err)
	}

	// add member
	{
		resp, _, err := cli.Chat().AddMember(ctx, &lark.AddMemberReq{
			ChatID:       "<CHAT_ID>",
			MemberIDType: lark.IDTypePtr(lark.IDTypeOpenID),
			IDList:       []string{"<USER_OPEN_ID>"},
		})
		fmt.Println(resp, err)
	}

	// delete member
	{
		resp, _, err := cli.Chat().DeleteMember(ctx, &lark.DeleteMemberReq{
			ChatID:       "<CHAT_ID>",
			MemberIDType: lark.IDTypePtr(lark.IDTypeOpenID),
			IDList:       []string{"<USER_OPEN_ID>"},
		})
		fmt.Println(resp, err)
	}

	// check bot is in member
	{
		resp, _, err := cli.Chat().IsInChat(ctx, &lark.IsInChatReq{
			ChatID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// search chat
	{
		resp, _, err := cli.Chat().SearchChat(ctx, &lark.SearchChatReq{
			Query: ptr.String("<SEARCH>"),
		})
		fmt.Println(resp, err)
	}
}
