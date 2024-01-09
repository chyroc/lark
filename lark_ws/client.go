package lark_ws

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/chyroc/lark"
	"github.com/gorilla/websocket"
)

type Client struct {
	Lark              *lark.Lark
	wsDialer          *websocket.Dialer
	conn              *websocket.Conn
	connURL           *url.URL
	serviceID         string
	connID            string
	autoReconnect     bool          // 是否自动重连，默认开启
	reconnectNonce    int           // 首次重连抖动，单位秒
	reconnectCount    int           // 重连次数，负数无限次
	reconnectInterval time.Duration // 重连间隔
	pingInterval      time.Duration // Ping间隔
	cache             *cache
	mu                sync.Mutex
}

type ClientOption func(cli *Client)

func WithAutoReconnect(b bool) ClientOption {
	return func(cli *Client) {
		cli.autoReconnect = b
	}
}

func New(larkCli *lark.Lark, opts ...ClientOption) *Client {
	cli := &Client{
		Lark: larkCli,
		wsDialer: &websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
		},
		conn:              nil,
		connURL:           nil,
		serviceID:         "",
		connID:            "",
		autoReconnect:     true,
		reconnectNonce:    30,
		reconnectCount:    -1,
		reconnectInterval: 2 * time.Minute,
		pingInterval:      2 * time.Minute,
		cache:             newCache(),
		mu:                sync.Mutex{},
	}

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func (c *Client) Start(ctx context.Context) (err error) {
	err = c.connect(ctx)
	if err != nil {
		c.logError(ctx, "connect failed, err: %s", err)
		if !isRetryErr(err) {
			return
		}
		c.disconnect(ctx)
		if c.autoReconnect {
			if err = c.reconnect(ctx); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	go c.pingLoop(ctx)
	select {}
}

func (c *Client) reconnect(ctx context.Context) (err error) {
	// 首次重连随机抖动
	if c.reconnectNonce > 0 {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(c.reconnectNonce * 1000)
		time.Sleep(time.Duration(num) * time.Millisecond)
	}

	if c.reconnectCount >= 0 {
		for i := 0; i < c.reconnectCount; i++ {
			success, err := c.tryConnect(ctx, i)
			if success || err != nil {
				return err
			}
			time.Sleep(c.reconnectInterval)
		}
		return fmt.Errorf("unable to connect to server after %d retries", c.reconnectCount)
	} else {
		i := 0
		for {
			success, err := c.tryConnect(ctx, i)
			if success || err != nil {
				return err
			}
			time.Sleep(c.reconnectInterval)
			i += 1
		}
	}
}

func (c *Client) tryConnect(ctx context.Context, cnt int) (bool, error) {
	c.logInfo(ctx, "trying to reconnect: %d", cnt+1)
	err := c.connect(ctx)
	if err == nil {
		return true, nil
	} else if !isRetryErr(err) {
		return false, err
	} else {
		c.logError(ctx, "connect failed, err: %v", err)
		return false, nil
	}
}

func (c *Client) pingLoop(ctx context.Context) {
	defer func() {
		if e := recover(); e != nil {
			c.logWarn(ctx, "ping loop panic, panic: %v, stack: %s", e, string(debug.Stack()))
		}
		// TODO: 短时间内一直 panic, 退出
		go c.pingLoop(ctx)
	}()

	for {
		// TODO: 锁
		if c.conn != nil {
			i, _ := strconv.ParseInt(c.serviceID, 10, 32)
			frame := newPingFrame(int32(i))
			bs, _ := frame.Marshal()

			err := c.writeMessage(websocket.BinaryMessage, bs)
			if err != nil {
				c.logWarn(ctx, "ping failed, err: %v", err)
			} else {
				c.logDebug(ctx, "ping success")
			}
		}
		time.Sleep(c.pingInterval)
	}
}
