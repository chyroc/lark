package examples

import (
	"context"
	"fmt"
	"time"

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

// RedisClient use this interface to mock redis client
type RedisClient interface {
	Get(ctx context.Context, key string) (string, time.Duration, error)
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

// NewRedisClient ...
func NewRedisClient() RedisClient {
	panic("no-impl")
}

// StoreRedis ...
type StoreRedis struct {
	redisCli RedisClient
}

// Get ...
func (r *StoreRedis) Get(ctx context.Context, key string) (string, time.Duration, error) {
	return r.redisCli.Get(ctx, key)
}

// Set ...
func (r *StoreRedis) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	return r.redisCli.Set(ctx, key, val, ttl)
}

// NewStoreRedis ...
func NewStoreRedis(redisCli RedisClient) lark.Store {
	return &StoreRedis{
		redisCli: redisCli,
	}
}

// ExampleISV ...
func ExampleISV() {
	ctx := context.Background()
	cli := lark.New(
		lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
		lark.WithISV(true),
		lark.WithStore(NewStoreRedis(NewRedisClient())),
	)

	// create chat
	{
		{
			tenantKey1Cli := cli.WithTenant("<TENANT_KEY_1>")

			resp, _, err := tenantKey1Cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
				Name: ptr.String("<CHAT_NAME_1>"),
			})
			fmt.Println(resp, err)
		}

		{
			tenantKey2Cli := cli.WithTenant("<TENANT_KEY_2>")

			resp, _, err := tenantKey2Cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
				Name: ptr.String("<CHAT_NAME_1>"),
			})
			fmt.Println(resp, err)
		}
	}
}
