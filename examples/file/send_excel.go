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
