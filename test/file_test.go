package test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_File(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().File().UploadImage(ctx, &lark.UploadImageReq{
			ImageType: lark.ImageTypeMessage,
			Image:     bytes.NewReader(nil),
		})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().File().DownloadImage(ctx, &lark.DownloadImageReq{
			ImageKey: "x",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "the app do not turn on bot")
	})

	t.Run("", func(t *testing.T) {
		f, err := os.Open("./file_1.png")
		as.Nil(err)
		defer f.Close()
		resp, _, err := AppALLPermission.Ins().File().UploadImage(ctx, &lark.UploadImageReq{
			ImageType: lark.ImageTypeMessage,
			Image:     f,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.ImageKey)
	})

	t.Run("", func(t *testing.T) {
		// this is file of ./test/file_1.png
		resp, _, err := AppALLPermission.Ins().File().DownloadImage(ctx, &lark.DownloadImageReq{
			ImageKey: File1.Key,
		})
		as.Nil(err)
		as.NotNil(resp)
		bs, err := io.ReadAll(resp.File)
		as.Nil(err)
		as.Len(bs, 84)
	})
}
