package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Helpdesk(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().StartService(ctx, &lark.StartServiceReq{
			// HumanService: ptr.Bool(false),
			// AppointedAgents: nil,
			OpenID: UserAdmin.OpenID,
			// CustomizedInfo: ptr.String("test"),
		})
		printData(resp, err)
		as.Nil(err)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().GetTicketList(ctx, &lark.GetTicketListReq{})
		printData(resp, err)
		as.Nil(err)
		// 6958447406052540443
		as.True(len(resp.Tickets) > 0)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().GetTicketList(ctx, &lark.GetTicketListReq{})
		printData(resp, err)
		as.Nil(err)
		// 6958447406052540443
		as.True(len(resp.Tickets) > 0)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk().GetTicketMessageList(ctx, &lark.GetTicketMessageListReq{
			TicketID: "6958447406052540443",
		})
		printData(resp, err)
		as.Nil(err)
		// 6958447406052540443
		// as.True(len(resp.Tickets) > 0)
	})
}
