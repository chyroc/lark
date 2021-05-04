package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

func ExampleSendMessage() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// send raw message
	{
		resp, _, err := cli.Message().SendRawMessage(ctx, &lark.SendRawMessageReq{
			ReceiveIDType: lark.IDTypePtr(lark.IDTypeEmail),
			ReceiveID:     ptr.String("<EMAIL>"),
			Content:       `{"text":"text"}`,
			MsgType:       lark.MsgTypeText,
		})
		fmt.Println(resp, err)
	}

	// send text message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendText(ctx, "<TEXT>")
		fmt.Println(resp, err)
	}

	// send image message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendImage(ctx, "<IMAGE_KEY>")
		fmt.Println(resp, err)
	}

	// send post message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendPost(ctx, "<POST>")
		fmt.Println(resp, err)
	}

	// send card message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendCard(ctx, "<CARD>")
		fmt.Println(resp, err)
	}

	// send share chat message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendShareChat(ctx, "<CHAT_ID>")
		fmt.Println(resp, err)
	}

	// send share user message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendShareUser(ctx, "<USER_ID>")
		fmt.Println(resp, err)
	}

	// send audio message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendAudio(ctx, "<FILE_KEY>")
		fmt.Println(resp, err)
	}

	// send media message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendMedia(ctx, "<IMAGE_KEY>", "<FILE_KEY>")
		fmt.Println(resp, err)
	}

	// send file message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendFile(ctx, "<FILE_KEY>")
		fmt.Println(resp, err)
	}

	// send sticker message
	{
		resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendSticker(ctx, "<FILE_KEY>")
		fmt.Println(resp, err)
	}
}
