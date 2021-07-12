package test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_DriveFile(t *testing.T) {
	as := assert.New(t)

	filename := "./_examples/bot.go"
	bs, err := ioutil.ReadFile(filename)
	as.Nil(err)

	ins := AppAllPermission.Ins()

	token := ""
	{
		resp, _, err := ins.Drive.GetDriveRootFolderMeta(ctx, &lark.GetDriveRootFolderMetaReq{})
		as.Nil(err)
		token = resp.Token
	}

	{
		resp, _, err := AppAllPermission.Ins().Drive.UploadDriveFile(ctx, &lark.UploadDriveFileReq{
			FileName:   filename,
			ParentType: "explorer",
			ParentNode: token,
			Size:       int64(len(bs)),
			Checksum:   "",
			File:       bytes.NewReader(bs),
		})
		printData(resp)
		printData(err)
		as.Nil(err)
	}
}
