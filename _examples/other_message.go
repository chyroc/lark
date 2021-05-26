package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

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
