package test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_CalendarACL_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := moduleCli.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
			// 	CalendarID: "x",
			// })
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{
				CalendarID: "x",
				ACLID:      "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_CalendarEventAttendee_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarEventAttendee(ctx, &lark.CreateCalendarEventAttendeeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventAttendeeList(ctx, &lark.GetCalendarEventAttendeeListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarEventAttendee(ctx, &lark.DeleteCalendarEventAttendeeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarEventAttendee(ctx, &lark.CreateCalendarEventAttendeeReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventAttendeeList(ctx, &lark.GetCalendarEventAttendeeListReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarEventAttendee(ctx, &lark.DeleteCalendarEventAttendeeReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_Calendar(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		summary := "summary-test"

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			printData(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.GetCalendar(ctx, &lark.GetCalendarReq{
				CalendarID: calendarID,
			})
			printData(resp, err)
			as.Nil(err)
			as.Equal(summary, resp.Summary)
		}

		{
			resp, _, err := moduleCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			printData(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{
				CalendarID: calendarID,
				Summary:    ptr.String(summary + "-update"),
			})
			printData(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
				CalendarID: calendarID,
			})
			printData(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
				CalendarID: calendarID,
			})
			printData(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
			printData(resp, err)
			as.Nil(err)
		}
	})
}

func Test_CalendarEvent(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		eventID := ""
		summary := "summary-test"
		defer func() {
			_, _, _ = moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
		}()

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			printData(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.CreateCalendarEvent(ctx, &lark.CreateCalendarEventReq{
				CalendarID: calendarID,
				Summary:    &summary,
				StartTime: &lark.CreateCalendarEventReqStartTime{
					Date: ptr.String("2020-09-01"),
				},
				EndTime: &lark.CreateCalendarEventReqEndTime{
					Date: ptr.String("2020-09-02"),
				},
			})
			printData(resp, err)
			as.Nil(err)
			eventID = resp.Event.EventID
		}

		{
			resp, _, err := moduleCli.GetCalendarEvent(ctx, &lark.GetCalendarEventReq{
				CalendarID: calendarID,
				EventID:    eventID,
			})
			printData(resp, err)
			as.Nil(err)
			as.Equal(summary, resp.Event.Summary)
		}

		{
			resp, _, err := moduleCli.GetCalendarEventList(ctx, &lark.GetCalendarEventListReq{
				CalendarID: calendarID,
			})
			printData(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UpdateCalendarEvent(ctx, &lark.UpdateCalendarEventReq{
				CalendarID: calendarID,
				EventID:    eventID,
				Summary:    ptr.String(summary + "-update"),
			})
			printData(resp, err)
			as.Nil(err)
		}

		{
			// TODO: user_access_token
			// resp, _, err := moduleCli.SubscribeCalendarEvent(ctx, &lark.SubscribeCalendarEventReq{
			// 	CalendarID: calendarID,
			// })
			// printData(resp, err)
			// as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendarEvent(ctx, &lark.DeleteCalendarEventReq{
				CalendarID: calendarID,
				EventID:    eventID,
			})
			printData(resp, err)
			as.Nil(err)
		}
	})
}

func Test_CalendarACL(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		aclID := ""
		defer func() {
			_, _, _ = moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
		}()

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			printData(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{
				UserIDType: nil,
				CalendarID: calendarID,
				Role:       lark.CalendarRoleWriter,
				Scope: &lark.CreateCalendarACLReqScope{
					Type:   "user",
					UserID: &UserAdmin.OpenID,
				},
			})
			printData(resp, err)
			as.Nil(err)
			aclID = resp.ACLID
		}

		{
			resp, _, err := moduleCli.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{
				CalendarID: calendarID,
			})
			printData(resp, err)
			as.Nil(err)
		}

		{
			// TODO: user_access_token
			// resp, _, err := moduleCli.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
			// 	CalendarID: calendarID,
			// })
			// printData(resp, err)
			// as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{
				CalendarID: calendarID,
				ACLID:      aclID,
			})
			printData(resp, err)
			as.Nil(err)
		}
	})
}
