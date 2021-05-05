package test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Helpdesk_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		helpdeskCli := cli.Helpdesk()

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

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		helpdeskCli := cli.Helpdesk()

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.StartService(ctx, &lark.StartServiceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetTicket(ctx, &lark.GetTicketReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetTicketList(ctx, &lark.GetTicketListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := helpdeskCli.UpdateTicket(ctx, &lark.UpdateTicketReq{})
			// as.NotNil(err)
			//   as.True(lark.GetErrorCode(err)>0)
		})
	})
}

func Test_Helpdesk(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().StartService(ctx, &lark.StartServiceReq{
			// HumanService: ptr.Bool(false),
			// AppointedAgents: nil,
			OpenID: UserAdmin.OpenID,
			// CustomizedInfo: ptr.String("test"),
		})
		spew.Dump(resp, err)
		as.Nil(err)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().GetTicketList(ctx, &lark.GetTicketListReq{})
		spew.Dump(resp, err)
		as.Nil(err)
		// 6958447406052540443
		as.True(len(resp.Tickets) > 0)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().GetTicketList(ctx, &lark.GetTicketListReq{})
		spew.Dump(resp, err)
		as.Nil(err)
		// 6958447406052540443
		as.True(len(resp.Tickets) > 0)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().GetTicketMessageList(ctx, &lark.GetTicketMessageListReq{
			TicketID: "6958447406052540443",
		})
		spew.Dump(resp, err)
		as.Nil(err)
		// 6958447406052540443
		// as.True(len(resp.Tickets) > 0)
	})
}
