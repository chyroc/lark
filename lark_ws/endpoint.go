package lark_ws

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/chyroc/lark"
)

// endpoint ...
type endpoint struct {
	URL          string                `json:"URL,omitempty"`
	ClientConfig *endpointClientConfig `json:"ClientConfig,omitempty"`
}

// endpointClientConfig ...
type endpointClientConfig struct {
	ReconnectCount    int `json:"ReconnectCount,omitempty"`
	ReconnectInterval int `json:"ReconnectInterval,omitempty"`
	ReconnectNonce    int `json:"ReconnectNonce,omitempty"`
	PingInterval      int `json:"PingInterval,omitempty"`
}

func (c *Client) getEndpoint(ctx context.Context) (*endpoint, error) {
	type GenerateCallbackWebsocketEndpointReq struct {
		AppID     string `json:"AppID,omitempty"`
		AppSecret string `json:"AppSecret,omitempty"`
		Locale    string `query:"locale" json:"-"`
	}
	type GenerateCallbackWebsocketEndpointResp struct {
		Code int64     `json:"code,omitempty"`
		Msg  string    `json:"msg,omitempty"`
		Data *endpoint `json:"data,omitempty"`
	}
	req := &lark.RawRequestReq{
		Scope:  "Callback",
		API:    "GenerateCallbackWebsocketEndpoint",
		Method: http.MethodPost,
		URL:    c.Lark.OpenBaseURL() + "/callback/ws/endpoint",
		Body: GenerateCallbackWebsocketEndpointReq{
			AppID:     c.Lark.AppID(),
			AppSecret: c.Lark.AppSecret(),
			Locale:    "zh",
		},
		MethodOption: &lark.MethodOption{},
	}
	resp := &GenerateCallbackWebsocketEndpointResp{}
	_, err := c.Lark.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) saveURL(endpointURL string) error {
	u, err := url.Parse(endpointURL)
	if err != nil {
		return fmt.Errorf("ws: invalid conn url: '%s'", endpointURL)
	}
	connID := u.Query().Get("device_id")
	serviceID := u.Query().Get("service_id")

	c.connID = connID
	c.serviceID = serviceID
	c.connURL = u

	return nil
}

func (c *Client) saveClientConfig(conf *endpointClientConfig) {
	c.reconnectCount = conf.ReconnectCount
	c.reconnectInterval = time.Duration(conf.ReconnectInterval) * time.Second
	c.reconnectNonce = conf.ReconnectNonce
	c.pingInterval = time.Duration(conf.PingInterval) * time.Second
}
