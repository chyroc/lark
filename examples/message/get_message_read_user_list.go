package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

// get user_list of read message
// 获取已阅消息的人员列表
func example() {
	appID := "<APP_ID>"
	appSecret := "<APP_SECRET>"
	messageID := "<MESSAGE_ID>"
	cli := lark.New(lark.WithAppCredential(appID, appSecret))
	ctx := context.Background()

	resp, _, err := cli.Message().GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
		UserIDType: lark.IDTypeOpenID,
		MessageID:  messageID,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range resp.Items {
		fmt.Println("user_id:", v.UserID)
	}
}
