package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func ExampleMiddleware() {
	r := lark.New(
		lark.WithAppCredential("", ""),
		lark.WithApiMiddleware(func(endpoint lark.ApiEndpoint) lark.ApiEndpoint {
			return func(ctx context.Context, rawHttpReq *lark.RawRequestReq, resp interface{}) (*lark.Response, error) {
				fmt.Println("log before request")
				response, err := endpoint(ctx, rawHttpReq, resp)
				fmt.Println("log after request")
				return response, err
			}
		}),
	)

	r.Message.SendRawMessage(context.Background(), &lark.SendRawMessageReq{})
}
