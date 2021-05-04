package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_HelpDesk(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		helpdeskCli := cli.HelpDesk()

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.StartService(ctx, &lark.StartServiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetTicket(ctx, &lark.GetTicketReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetTicketList(ctx, &lark.GetTicketListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := helpdeskCli.UpdateTicket(ctx, &lark.UpdateTicketReq{})
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})
	})
}
