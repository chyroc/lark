/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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

	resp, _, err := cli.Message.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
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
