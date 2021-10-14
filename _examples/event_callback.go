package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/chyroc/lark"
)

func ExampleEventCallback() {
	cli := lark.New(
		lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
		lark.WithEventCallbackVerify("<ENCRYPT_KEY>", "<VERIFICATION_TOKEN>"),
	)

	// handle chat create callback
	cli.EventCallback.HandlerEventV1P2PChatCreate(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV1, event *lark.EventV1P2PChatCreate) (string, error) {
		_, _, err := cli.Message.Send().ToChatID(event.ChatID).SendText(ctx, "hi")
		return "", err
	})

	// handle message callback
	cli.EventCallback.HandlerEventV2IMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
		content, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
		if err != nil {
			return "", err
		}
		switch {
		case content.Text != nil:
			_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got text: %s", content.Text.Text))
		case content.File != nil:
			_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got file: %s, key: %s", content.File.FileName, content.File.FileKey))
		case content.Image!=nil:
			_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got image: %s", content.Image.ImageKey))
		}
		return "", err
	})

	// handle add bot
	cli.EventCallback.HandlerEventV2IMChatMemberBotAddedV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMChatMemberBotAddedV1) (string, error) {
		return "", nil
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		cli.EventCallback.ListenCallback(r.Context(), r.Body, w)
	})

	fmt.Println("start server ...")
	log.Fatal(http.ListenAndServe(":9726", nil))
}
