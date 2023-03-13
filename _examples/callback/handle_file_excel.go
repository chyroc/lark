package callback

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/chyroc/lark"
)

func handleFileReceive(ctx context.Context, cli *lark.Lark, event *lark.EventV2IMMessageReceiveV1, content *lark.MessageContent) error {
	fileName := content.File.FileName
	fileKey := content.File.FileKey

	file, _, err := cli.Message.GetMessageFile(ctx, &lark.GetMessageFileReq{
		MessageID: event.Message.MessageID,
		FileKey:   fileKey,
		Type:      "file",
	})
	if err != nil {
		return err
	}

	// handle file
	_ = fileName
	_ = file

	return nil
}

func example() {
	cli := newLarkCli()

	// handle message callback
	cli.EventCallback.HandlerEventV2IMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
		content, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
		if err != nil {
			return "", err
		}
		switch event.Message.MessageType {
		case lark.MsgTypeFile:
			return "", handleFileReceive(ctx, cli, event, content)
		}
		return "", nil
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		cli.EventCallback.ListenCallback(r.Context(), r.Body, w)
	})

	fmt.Println("start server ...")
	log.Fatal(http.ListenAndServe(":9726", nil))
}

func newLarkCli() *lark.Lark {
	cli := lark.New(
		lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
		lark.WithEventCallbackVerify("<ENCRYPT_KEY>", "<VERIFICATION_TOKEN>"),
	)

	cli = lark.New(
		lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
		lark.WithEventCallbackVerify("<ENCRYPT_KEY>", "<VERIFICATION_TOKEN>"),
		// 如果你希望你的回调是非阻塞的
		// if you want your callback is non-blocking
		lark.WithNonBlockingCallback(true),
	)
	return cli
}
