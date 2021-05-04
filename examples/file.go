package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/chyroc/lark"
)

func ExampleFile() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// upload image
	{
		f, err := os.Open("<FILE_PATH>")
		if err != nil {
			panic(err)
		}
		resp, _, err := cli.File().UploadImage(ctx, &lark.UploadImageReq{
			ImageType: lark.ImageTypeMessage,
			Image:     f,
		})
		fmt.Println(resp, err)
	}

	// download image
	{
		resp, _, err := cli.File().DownloadImage(ctx, &lark.DownloadImageReq{
			ImageKey: "<IMAGE_KEY>",
		})
		fmt.Println(resp, err)
	}

	// upload file
	{
		f, err := os.Open("<FILE_PATH.doc>")
		if err != nil {
			panic(err)
		}
		resp, _, err := cli.File().UploadFile(ctx, &lark.UploadFileReq{
			FileType: lark.FileTypeDoc,
			FileName: "<FILE_NAME>",
			File:     f,
		})
		fmt.Println(resp, err)
	}

	// download file
	{
		resp, _, err := cli.File().DownloadFile(ctx, &lark.DownloadFileReq{
			FileKey: "<FILE_KEY>",
		})
		fmt.Println(resp, err)
	}
}
