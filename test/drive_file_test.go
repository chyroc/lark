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
	"bytes"
	"hash/adler32"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_DriveFile(t *testing.T) {
	as := assert.New(t)

	filename := "./_examples/bot.go"
	bs := readFile(filename)

	ins := AppAllPermission.Ins()

	rootFolderMetaToken := ""
	{
		resp, _, err := ins.Drive.GetDriveRootFolderMeta(ctx, &lark.GetDriveRootFolderMetaReq{})
		as.Nil(err)
		rootFolderMetaToken = resp.Token
	}

	fileToken := ""
	{
		x := adler32.New()
		x.Write(bs)
		resp, _, err := AppAllPermission.Ins().Drive.UploadDriveFile(ctx, &lark.UploadDriveFileReq{
			FileName:   filename,
			ParentType: "explorer",
			ParentNode: rootFolderMetaToken,
			Size:       int64(len(bs)),
			Checksum:   ptrString(strconv.FormatInt(int64(x.Sum32()), 10)),
			File:       bytes.NewReader(bs),
		})
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.FileToken)
		fileToken = resp.FileToken
	}

	{
		resp, _, err := AppAllPermission.Ins().Drive.DownloadDriveFile(ctx, &lark.DownloadDriveFileReq{
			FileToken: fileToken,
		})
		as.Nil(err)
		as.NotNil(resp)

		as.Equal(filename, resp.Filename)
		bs2, err := ioutil.ReadAll(resp.File)
		as.Nil(err)
		as.Equal(bs, bs2)
	}
}
