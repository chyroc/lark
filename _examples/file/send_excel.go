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
	"os"

	"github.com/chyroc/lark"
)

// send excel file example
// 发送 excel 文件例子
func example() {
	appID := "<APP_ID>"
	appSecret := "<APP_SECRET>"
	chatID := "<CHAT_ID>"
	cli := lark.New(lark.WithAppCredential(appID, appSecret))
	ctx := context.Background()

	f, err := os.Open("/path/test.xlsx")
	if err != nil {
		panic(err)
	}
	resp, _, err := cli.File.UploadFile(ctx, &lark.UploadFileReq{
		FileType: lark.FileTypeXls,
		FileName: "test.xlsx",
		File:     f,
	})
	if err != nil {
		panic(err)
	}

	_, _, err = cli.Message.Send().ToChatID(chatID).SendFile(ctx, resp.FileKey)
	if err != nil {
		panic(err)
	}
}
