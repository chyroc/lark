package lark_ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	ws "github.com/gorilla/websocket"
)

func (c *Client) receiveMessageLoop(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.logError(ctx, "receive message loop panic, err: %v, stack: %s", err, string(debug.Stack()))
		}
		c.disconnect(ctx)
		if c.autoReconnect {
			if err := c.reconnect(ctx); err != nil {
				c.logError(ctx, err.Error())
			}
		}
	}()

	for {
		if c.conn == nil {
			c.logError(ctx, "connection is closed, receive message loop exit")
			return
		}

		messageType, msg, err := c.conn.ReadMessage()
		if err != nil {
			c.logError(ctx, "receive message failed, err: %v", err)
			return
		}

		if messageType != ws.BinaryMessage {
			c.logWarn(ctx, "receive unknown message, message_type: %d, message: %s", messageType, msg)
			continue
		}

		go c.handleMessage(ctx, msg)
	}
}

func (c *Client) handleMessage(ctx context.Context, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			c.logError(ctx, "handle message panic, err: %v, stack: %s", err, string(debug.Stack()))
		}
	}()

	var frame Frame
	if err := frame.Unmarshal(msg); err != nil {
		c.logError(ctx, "unmarshal message failed, error: %v", err)
		return
	}

	switch frameType(frame.Method) {
	case frameTypeControl:
		c.handleControlFrame(ctx, frame)
	case frameTypeData:
		c.handleDataFrame(ctx, frame)
	default:
	}
}

func (c *Client) handleControlFrame(ctx context.Context, frame Frame) {
	hs := wsHeader(frame.Headers)
	t := hs.GetString("type")

	switch messageType(t) {
	case messageTypePong:
		c.logDebug(ctx, "receive pong")
		if len(frame.Payload) == 0 {
			return
		}
		conf := &endpointClientConfig{}
		if err := json.Unmarshal(frame.Payload, conf); err != nil {
			c.logWarn(ctx, "unmarshal client config failed, err: %v", err)
			return
		}
		c.saveClientConfig(conf)
	default:
	}
}

func (c *Client) handleDataFrame(ctx context.Context, frame Frame) {
	hs := wsHeader(frame.Headers)
	sum := hs.GetInt("sum") // 拆包数, 未拆包为 1
	seq := hs.GetInt("seq") // 包序号, 未拆包为 0
	msgID := hs.GetString("message_id")
	traceID := hs.GetString("trace_id")
	type_ := hs.GetString("type")

	pl := frame.Payload
	if sum > 1 {
		if pl = c.combine(msgID, sum, seq, pl); pl == nil {
			return
		}
	}

	c.logDebug(ctx, "receive message, message_type: %s, message_id: %s, trace_id: %s, payload: %s", type_, msgID, traceID, pl)

	start := time.Now().UnixMilli()
	switch messageType(type_) {
	case messageTypeEvent:
		c.Lark.EventCallback.ListenCallback(ctx, bytes.NewReader(pl), c.newHTTPResponse(ctx, frame, start))
	case messageTypeCard:
		c.Lark.EventCallback.ListenCallback(ctx, bytes.NewReader(pl), c.newHTTPResponse(ctx, frame, start))
	default:
		return
	}
}

func (c *Client) combine(msgID string, sum, seq int, bs []byte) []byte {
	val, ok := c.cache.get(msgID)
	if !ok {
		buf := make([][]byte, sum)
		buf[seq] = bs
		c.cache.set(msgID, buf)
		return nil
	}

	buf := val.([][]byte)
	buf[seq] = bs
	capacity := 0
	for _, v := range buf {
		if len(v) == 0 {
			c.cache.set(msgID, buf)
			return nil
		}
		capacity += len(v)
	}

	pl := make([]byte, 0, capacity)
	for _, v := range buf {
		pl = append(pl, v...)
	}

	return pl
}

func (c *Client) writeMessage(messageType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return fmt.Errorf("connection is closed")
	}
	return c.conn.WriteMessage(messageType, data)
}
