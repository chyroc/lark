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
package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/chyroc/lark"
)

// ExampleFile ...
func ExampleFile() {
	ctx := context.Background()
	cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	// upload image
	{
		f, err := os.Open("<FILE_PATH>")
		if err != nil {
			panic(err)
		}
		resp, _, err := cli.File.UploadImage(ctx, &lark.UploadImageReq{
			ImageType: lark.ImageTypeMessage,
			Image:     f,
		})
		fmt.Println(resp, err)
	}

	// download image
	{
		resp, _, err := cli.File.DownloadImage(ctx, &lark.DownloadImageReq{
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
		resp, _, err := cli.File.UploadFile(ctx, &lark.UploadFileReq{
			FileType: lark.FileTypeDoc,
			FileName: "<FILE_NAME>",
			File:     f,
		})
		fmt.Println(resp, err)
	}

	// download file
	{
		resp, _, err := cli.File.DownloadFile(ctx, &lark.DownloadFileReq{
			FileKey: "<FILE_KEY>",
		})
		fmt.Println(resp, err)
	}
}
