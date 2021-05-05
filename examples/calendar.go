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
}
