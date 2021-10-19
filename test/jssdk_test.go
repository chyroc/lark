package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Jssdk(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		res, _, err := AppAllPermission.Ins().Jssdk.GetJssdkTicket(ctx, &lark.GetJssdkTicketReq{})
		as.Nil(err)
		as.NotNil(res)
		as.NotEmpty(res.Ticket)
	})

	t.Run("", func(t *testing.T) {
		res, err := AppAllPermission.Ins().Jssdk.GenerateJssdkSignature(ctx, &lark.GenerateJssdkSignatureReq{URL: "https://m.mm.cn/ttc/3541093/3131_1.html"})
		as.Nil(err)
		as.NotEmpty(res)
	})
}
