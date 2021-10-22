package lark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Request(t *testing.T) {
	as := assert.New(t)
	ins := &Lark{}
	ctx := context.Background()

	fmt.Println(1)

	t.Run("", func(t *testing.T) {
		type req struct {
			ID string `path:"id"`
		}

		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com/:id",
			Body: req{
				ID: string("1234"),
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("GET", resp.Method)
		as.Equal("http://x.com/1234", resp.URL)
		as.Nil(resp.Body)
	})

	t.Run("", func(t *testing.T) {
		type req struct {
			ID int `path:"id"`
		}
		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com/:id",
			Body: req{
				ID: 1234,
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("GET", resp.Method)
		as.Equal("http://x.com/1234", resp.URL)
		as.Nil(resp.Body)
	})

	t.Run("", func(t *testing.T) {
		type req struct {
			Type string `path:"type"`
			ID   int    `path:"id"`
		}
		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com/:type/:id",
			Body: req{
				Type: "x",
				ID:   1234,
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("GET", resp.Method)
		as.Equal("http://x.com/x/1234", resp.URL)
		as.Nil(resp.Body)
	})

	t.Run("", func(t *testing.T) {
		type req struct {
			Type string `query:"type"`
			ID   int    `path:"id"`
		}
		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com/:id",
			Body: req{
				Type: "x",
				ID:   1234,
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("GET", resp.Method)
		as.Equal("http://x.com/1234?type=x", resp.URL)
		as.Nil(resp.Body)
	})

	t.Run("", func(t *testing.T) {
		type req struct {
			Type string `query:"type" json:"-"`
			ID   int    `path:"id" json:"-"`
			Name string `json:"name"`
		}
		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com/:id",
			Body: req{
				Type: "x",
				ID:   1234,
				Name: "lark",
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("GET", resp.Method)
		as.Equal("http://x.com/1234?type=x", resp.URL)
		as.NotNil(resp.Body)
		bs, err := ioutil.ReadAll(resp.Body)
		as.Nil(err)
		as.Equal(`{"name":"lark"}`, string(bs))
	})

	t.Run("", func(t *testing.T) {
		type req struct {
			ImageType ImageType `json:"image_type,omitempty"` // 图片类型,**示例值**："message",**可选值有**：,- `message`：用于发送消息,- `avatar`：用于设置头像
			Image     io.Reader `json:"image,omitempty"`      // 图片内容,**示例值**：二进流
		}
		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com",
			IsFile: true,
			Body: req{
				ImageType: ImageTypeMessage,
				Image:     bytes.NewReader([]byte("hi")),
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("GET", resp.Method)
		as.Equal("http://x.com", resp.URL)
		as.NotNil(resp.Body)
		// bs, err := ioutil.ReadAll(resp.Body)
		// as.Nil(err)
		// as.Equal(`{"name":"lark"}`, string(bs))
	})

	t.Run("", func(t *testing.T) {
		type req struct {
			Types []string `query:"types"`
		}
		resp, err := ins.parseRawHttpRequest(ctx, &RawRequestReq{
			Method: "get",
			URL:    "http://x.com",
			IsFile: true,
			Body: req{
				Types: []string{"a", "b"},
			},
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("http://x.com?types=a&types=b", resp.URL)
	})
}

func Test_eventReq_unmarshalEvent(t *testing.T) {
	as := assert.New(t)

	body := `{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "im.message.receive_v1",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "sender": {
            "sender_id": {
                "union_id": "on_8ed6aa67826108097d9ee143816345",
                "user_id": "e33ggbyz",
                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
            },
            "sender_type": "user",
            "tenant_key": "736588c9260f175e"
        },
        "message": {
            "message_id": "om_5ce6d572455d361153b7cb51da133945",
            "root_id": "om_5ce6d572455d361153b7cb5xxfsdfsdfdsf",
            "parent_id": "om_5ce6d572455d361153b7cb5xxfsdfsdfdsf",
            "create_time": "1609073151345",
            "chat_id": "oc_5ce6d572455d361153b7xx51da133945",
            "chat_type": "group",
            "message_type": "text",
            "content": "{\"text\":\"test\"}",
            "mentions": [
                {
                    "key": "@_user_1",
                    "id": {
                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                        "user_id": "e33ggbyz",
                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                    },
                    "name": "Tom",
                    "tenant_key": "736588c9260f175e"
                }
            ]
        }
    }
}`
	req := &eventReq{}
	as.Nil(json.Unmarshal([]byte(body), &req))

	event := new(EventV2IMMessageReceiveV1)
	as.Nil(req.unmarshalEvent(event))
	as.Equal("om_5ce6d572455d361153b7cb51da133945", event.Message.MessageID)
}
