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
		x := adler32.New()
		x.Write(bs)
		resp, _, err := AppAllPermission.Ins().Drive.UploadDriveFile(ctx, &lark.UploadDriveFileReq{
			FileName:   filename,
			ParentType: "explorer",
			ParentNode: token,
			Size:       int64(len(bs)),
			Checksum:   strconv.FormatInt(int64(x.Sum32()), 10),
			File:       bytes.NewReader(bs),
		})
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.FileToken)
	}
}
