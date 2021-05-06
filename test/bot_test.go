package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Bot_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		fileCli := cli.Bot()

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.GetBotInfo(ctx, &lark.GetBotInfoReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		fileCli := cli.Bot()

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.GetBotInfo(ctx, &lark.GetBotInfoReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_Bot(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Bot().GetBotInfo(ctx, &lark.GetBotInfoReq{})
		bs, _ := json.Marshal(resp)
		fmt.Println(string(bs))
		// printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("lark-sdk", resp.AppName)
	})
}
