package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

// ExampleBot ...
func ExampleBot() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// get bot info
	{
		resp, _, err := cli.Bot.GetBotInfo(ctx, &lark.GetBotInfoReq{})
		fmt.Println(resp, err)
	}
}
