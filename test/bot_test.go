package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Bot(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Bot.GetBotInfo(ctx, &lark.GetBotInfoReq{})
		bs, _ := json.Marshal(resp)
		fmt.Println(string(bs))
		// printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("lark-sdk", resp.AppName)
	})
}
