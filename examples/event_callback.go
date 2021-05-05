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
		lark.WithEventCallbackVerify("<ENCRYPY_KEY>", "<VERIFICATION_TOKEN>"),
	)

	// handle message callback
	cli.EventCallback().HandlerEventIMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeader, event *lark.EventIMMessageReceiveV1) (string, error) {
		_, _, err := cli.Message().Reply(event.Message.MessageID).SendText(ctx, "hi, "+event.Message.Content)
		return "", err
	})

	// handle add bot
	cli.EventCallback().HandlerEventIMChatMemberBotAddedV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeader, event *lark.EventIMChatMemberBotAddedV1) (string, error) {
		return "", nil
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		cli.EventCallback().ListenCallback(r.Context(), r.Body, w)
	})

	fmt.Println("start server ...")
	log.Fatal(http.ListenAndServe(":9726", nil))
}
