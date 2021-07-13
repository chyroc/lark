package test

import (
	"bytes"
	"hash/adler32"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
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
			Checksum:   strconv.FormatInt(int64(x.Sum32()), 10),
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
			Range:     [2]int64{},
		})
		as.Nil(err)
		as.NotNil(resp)

		as.Equal(filename, resp.Filename)
		bs2, err := ioutil.ReadAll(resp.File)
		as.Nil(err)
		as.Equal(bs, bs2)
	}
}
