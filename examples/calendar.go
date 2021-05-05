package examples

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

func ExampleCalendar() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// create calendar
	{
		resp, _, err := cli.Calendar().CreateCalendar(ctx, &lark.CreateCalendarReq{
			Summary: ptr.String("<SUMMARY>"),
		})
		fmt.Println(resp, err)
	}

	// update calendar
	{
		resp, _, err := cli.Calendar().UpdateCalendar(ctx, &lark.UpdateCalendarReq{
			CalendarID: "<CALENDAR_ID>",
			Summary:    ptr.String("<SUMMARY>"),
		})
		fmt.Println(resp, err)
	}

	// delete calendar
	{
		resp, _, err := cli.Calendar().DeleteCalendar(ctx, &lark.DeleteCalendarReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// get calendar
	{
		resp, _, err := cli.Calendar().GetCalendar(ctx, &lark.GetCalendarReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// get calendar list
	{
		resp, _, err := cli.Calendar().GetCalendarList(ctx, &lark.GetCalendarListReq{})
		fmt.Println(resp, err)
	}

	// search calendar
	{
		resp, _, err := cli.Calendar().SearchCalendar(ctx, &lark.SearchCalendarReq{
			Query: "<SEARCH>",
		})
		fmt.Println(resp, err)
	}

	// subscribe calendar
	{
		resp, _, err := cli.Calendar().SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// unsubscribe calendar
	{
		resp, _, err := cli.Calendar().UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// create calendar event
	{
		resp, _, err := cli.Calendar().CreateCalendarEvent(ctx, &lark.CreateCalendarEventReq{
			Summary: ptr.String("<SUMMARY>"),
		})
		fmt.Println(resp, err)
	}

	// update calendar event
	{
		resp, _, err := cli.Calendar().UpdateCalendarEvent(ctx, &lark.UpdateCalendarEventReq{
			CalendarID: "<CALENDAR_ID>",
			Summary:    ptr.String("<SUMMARY>"),
		})
		fmt.Println(resp, err)
	}

	// delete calendar event
	{
		resp, _, err := cli.Calendar().DeleteCalendarEvent(ctx, &lark.DeleteCalendarEventReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// get calendar event
	{
		resp, _, err := cli.Calendar().GetCalendarEvent(ctx, &lark.GetCalendarEventReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// get calendar list event
	{
		resp, _, err := cli.Calendar().GetCalendarEventList(ctx, &lark.GetCalendarEventListReq{})
		fmt.Println(resp, err)
	}

	// search calendar event
	{
		resp, _, err := cli.Calendar().SearchCalendarEvent(ctx, &lark.SearchCalendarEventReq{
			Query: "<SEARCH>",
		})
		fmt.Println(resp, err)
	}

	// subscribe calendar event
	{
		resp, _, err := cli.Calendar().SubscribeCalendarEvent(ctx, &lark.SubscribeCalendarEventReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// create calendar acl
	{
		resp, _, err := cli.Calendar().CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{
			CalendarID: "<CALENDAR_ID>",
			Role:       lark.CalendarRoleReader,
		})
		fmt.Println(resp, err)
	}

	// get calendar acl list
	{
		resp, _, err := cli.Calendar().GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// delete calendar acl
	{
		resp, _, err := cli.Calendar().DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{
			CalendarID: "<CALENDAR_ID>",
			ACLID:      "<ACL_ID>",
		})
		fmt.Println(resp, err)
	}

	// subscribe calendar acl
	{
		resp, _, err := cli.Calendar().SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
			CalendarID: "<CALENDAR_ID>",
		})
		fmt.Println(resp, err)
	}

	// create calendar event attendee
	{
		resp, _, err := cli.Calendar().CreateCalendarEventAttendee(ctx, &lark.CreateCalendarEventAttendeeReq{
			CalendarID: "<CALENDAR_ID>",
			EventID:    "<EVENT_ID>",
			Attendees: []*lark.CreateCalendarEventAttendeeReqAttendee{
				{
					UserID: ptr.String("<USER_ID>"),
				},
			},
		})
		fmt.Println(resp, err)
	}

	// get calendar event attendee list
	{
		resp, _, err := cli.Calendar().GetCalendarEventAttendeeList(ctx, &lark.GetCalendarEventAttendeeListReq{
			CalendarID: "<CALENDAR_ID>",
			EventID:    "<EVENT_ID>",
		})
		fmt.Println(resp, err)
	}

	// delete calendar event attendee
	{
		resp, _, err := cli.Calendar().DeleteCalendarEventAttendee(ctx, &lark.DeleteCalendarEventAttendeeReq{
			CalendarID: "<CALENDAR_ID>",
			EventID:    "<EVENT_ID>",
			AttendeeIDs: []string{
				"<USER_ID>",
			},
		})
		fmt.Println(resp, err)
	}

	// get calendar event attendee chat member list
	{
		resp, _, err := cli.Calendar().GetCalendarEventAttendeeChatMemberList(ctx, &lark.GetCalendarEventAttendeeChatMemberListReq{
			CalendarID: "<CALENDAR_ID>",
			EventID:    "<EVENT_ID>",
			AttendeeID: "<CHAT_ID>",
		})
		fmt.Println(resp, err)
	}

	// get calendar freebusy list
	{
		resp, _, err := cli.Calendar().GetCalendarFreeBusyList(ctx, &lark.GetCalendarFreeBusyListReq{
			TimeMin: "2020-10-28T12:00:00+08:00",
			TimeMax: "2020-10-29T12:00:00+08:00",
			UserID:  ptr.String("<USER_ID>"),
		})
		fmt.Println(resp, err)
	}

	// create calendar timeoff event
	{
		resp, _, err := cli.Calendar().CreateCalendarTimeoffEvent(ctx, &lark.CreateCalendarTimeoffEventReq{
			UserID:    "<USER_ID>",
			StartTime: "2021-01-01",
			EndTime:   "2021-01-02",
			Title:     ptr.String("<TITLE>"),
		})
		fmt.Println(resp, err)
	}

	// delete calendar timeoff event
	{
		resp, _, err := cli.Calendar().DeleteCalendarTimeoffEvent(ctx, &lark.DeleteCalendarTimeoffEventReq{
			TimeoffEventID: "TIMEOFF_EVENT_ID",
		})
		fmt.Println(resp, err)
	}

	// subscribe calendar acl
	{
		resp, _, err := cli.Calendar().GenerateCaldavConf(ctx, &lark.GenerateCaldavConfReq{
			DeviceName: ptr.String("DEVICE_NAME"),
		})
		fmt.Println(resp, err)
	}
}
