package lark_ws

import (
	"context"

	"github.com/chyroc/lark"
)

func (c *Client) logError(ctx context.Context, msg string, args ...interface{}) {
	if len(args) == 0 {
		c.Lark.Log(ctx, lark.LogLevelError, "[lark] Websocket "+msg)
		return
	}
	c.Lark.Log(ctx, lark.LogLevelError, "[lark] Websocket "+msg, args...)
}

func (c *Client) logWarn(ctx context.Context, msg string, args ...interface{}) {
	if len(args) == 0 {
		c.Lark.Log(ctx, lark.LogLevelWarn, "[lark] Websocket "+msg)
		return
	}
	c.Lark.Log(ctx, lark.LogLevelWarn, "[lark] Websocket "+msg, args...)
}

func (c *Client) logInfo(ctx context.Context, msg string, args ...interface{}) {
	if len(args) == 0 {
		c.Lark.Log(ctx, lark.LogLevelInfo, "[lark] Websocket "+msg)
		return
	}
	c.Lark.Log(ctx, lark.LogLevelInfo, "[lark] Websocket "+msg, args...)
}

func (c *Client) logDebug(ctx context.Context, msg string, args ...interface{}) {
	if len(args) == 0 {
		c.Lark.Log(ctx, lark.LogLevelDebug, "[lark] Websocket "+msg)
		return
	}
	c.Lark.Log(ctx, lark.LogLevelDebug, "[lark] Websocket "+msg, args...)
}
