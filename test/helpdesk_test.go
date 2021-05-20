package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Helpdesk(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk.StartHelpdeskService(ctx, &lark.StartHelpdeskServiceReq{
			// HumanService: ptr.Bool(false),
			// AppointedAgents: nil,
			OpenID: UserAdmin.OpenID,
			// CustomizedInfo: ptr.String("test"),
		})
		printData(resp, err)
		as.Nil(err)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{})
		printData(resp, err)
		as.Nil(err)
		// 6958447406052540443
		as.True(len(resp.Tickets) > 0)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{})
		printData(resp, err)
		as.Nil(err)
		// 6958447406052540443
		as.True(len(resp.Tickets) > 0)
	})

	t.Run("", func(t *testing.T) {
		resp, _, err := HelpdeskAllPermission.Ins().Helpdesk.GetHelpdeskTicketMessageList(ctx, &lark.GetHelpdeskTicketMessageListReq{
			TicketID: "6958447406052540443",
		})
		printData(resp, err)
		as.Nil(err)
		// 6958447406052540443
		// as.True(len(resp.Tickets) > 0)
	})
}
