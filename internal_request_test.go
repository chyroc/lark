package lark

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		type req struct {
			ID string `path:"id"`
		}
		resp, err := parseRequestParam(&requestParam{
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
		resp, err := parseRequestParam(&requestParam{
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
		resp, err := parseRequestParam(&requestParam{
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
		resp, err := parseRequestParam(&requestParam{
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
		resp, err := parseRequestParam(&requestParam{
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
		bs, err := io.ReadAll(resp.Body)
		as.Nil(err)
		as.Equal(`{"name":"lark"}`, string(bs))
	})
}
