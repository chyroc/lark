package main

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

func main() {
	ctx := context.Background()
	r := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	resp, _, err := r.Chat().CreateChat(ctx, &lark.CreateChatReq{
		Name: ptr.String("chat-name"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("chat created: ", resp.ChatID)
}
