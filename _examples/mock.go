package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

// ExampleMock ...
func ExampleMock() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	cli.Mock().MockChatCreateChat(func(ctx context.Context, request *lark.CreateChatReq, options ...lark.MethodOptionFunc) (*lark.CreateChatResp, *lark.Response, error) {
		panic("mocked")
	})
	defer cli.Mock().UnMockChatCreateChat()

	// get bot info
	{
		resp, _, err := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{})
		fmt.Println(resp, err)
	}
}
