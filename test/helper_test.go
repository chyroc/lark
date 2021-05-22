package test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt64() int64 {
	return rand.Int63()
}

func randInt64String() string {
	return strconv.FormatInt(randInt64(), 10)
}

func mockGetTenantAccessTokenFailed(ctx context.Context) (*lark.TokenExpire, *lark.Response, error) {
	return nil, nil, fmt.Errorf("failed")
}

func printData(datas ...interface{}) {
	for _, v := range datas {
		printDataSingle(v)
	}
}

func printDataSingle(v interface{}) {
	vt := reflect.TypeOf(v)
	if vt != nil {
		if vt.Kind() == reflect.Ptr {
			vt = vt.Elem()
		}
		fmt.Printf(vt.Name() + "#")
	}
	if v == nil {
		fmt.Println("nil")
		return
	}
	switch v := v.(type) {
	case int:
		fmt.Println(v)
	case int8:
		fmt.Println(v)
	case int16:
		fmt.Println(v)
	case int32:
		fmt.Println(v)
	case int64:
		fmt.Println(v)
	case uint:
		fmt.Println(v)
	case uint8:
		fmt.Println(v)
	case uint16:
		fmt.Println(v)
	case uint32:
		fmt.Println(v)
	case uint64:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	case error:
		fmt.Println(v)
	default:
		vv, _ := json.Marshal(v)
		fmt.Println(string(vv))
	}
}

func Test_Helper(t *testing.T) {
	as := assert.New(t)

	t.Run("GetErrorCode", func(t *testing.T) {
		as.Equal(int64(-1), lark.GetErrorCode(fmt.Errorf("x")))
	})

	t.Run("UnwrapMessageContent", func(t *testing.T) {
		t.Run("image", func(t *testing.T) {
			_, err := lark.UnwrapMessageContent(lark.MsgTypeInteractive, `{"image_key":"image-x"}`)
			as.NotNil(err)
			as.Contains(err.Error(), "unknown message type")
		})

		t.Run("image", func(t *testing.T) {
			_, err := lark.UnwrapMessageContent(lark.MsgTypeText, "")
			as.NotNil(err)
			as.Contains(err.Error(), "invalid content")
		})

		t.Run("text", func(t *testing.T) {
			res, err := lark.UnwrapMessageContent(lark.MsgTypeText, `{"text":"hi"}`)
			as.Nil(err)
			as.Equal("hi", res.Text.Text)
		})

		t.Run("image", func(t *testing.T) {
			res, err := lark.UnwrapMessageContent(lark.MsgTypeImage, `{"image_key":"image-x"}`)
			as.Nil(err)
			as.Equal("image-x", res.Image.ImageKey)
		})
	})

	t.Run("", func(t *testing.T) {
		printData(nil)
		printData("hi")
		printData(1)
		printData(false)
		printData(lark.MsgTypeText)
		printData(lark.SendRawMessageReq{Content: "x"})
		var x *lark.SendRawMessageReq = nil
		printData(x)
		printData(lark.SendRawMessageReq{Content: "x"}, x)
	})
}

func IsNotInCI() bool {
	isNotInCI := os.Getenv("IN_CI") == ""
	if isNotInCI {
		fmt.Println("NOT IN CI, SKIP")
	}
	return isNotInCI
}

func IsInCI() bool {
	return os.Getenv("IN_CI") != ""
}
