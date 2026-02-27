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
package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_File(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		f, err := os.Open("./file_1.png")
		as.Nil(err)
		defer f.Close()
		resp, _, err := AppAllPermission.Ins().File.UploadImage(ctx, &lark.UploadImageReq{
			ImageType: string(lark.ImageTypeMessage),
			Image:     f,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.ImageKey)
	})

	t.Run("", func(t *testing.T) {
		// this is file of ./test/file_1.png
		resp, _, err := AppAllPermission.Ins().File.DownloadImage(ctx, &lark.DownloadImageReq{
			ImageKey: File1.Key,
		})
		as.Nil(err)
		as.NotNil(resp)
		bs, err := ioutil.ReadAll(resp.File)
		as.Nil(err)
		as.Len(bs, 84)
	})

	t.Run("", func(t *testing.T) {
		f, err := os.Open("./file_2.docx")
		as.Nil(err)
		defer f.Close()
		resp, _, err := AppAllPermission.Ins().File.UploadFile(ctx, &lark.UploadFileReq{
			FileType: lark.FileTypeDoc,
			FileName: "file2.docx",
			Duration: nil,
			File:     f,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.FileKey)
	})

	t.Run("", func(t *testing.T) {
		// ./test/file_2.docx
		resp, _, err := AppAllPermission.Ins().File.DownloadFile(ctx, &lark.DownloadFileReq{
			FileKey: File2.Key,
		})
		// printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		bs, err := ioutil.ReadAll(resp.File)
		as.Nil(err)
		as.Len(bs, 3247)
	})
}
