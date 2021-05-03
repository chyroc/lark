package test

import (
	"bytes"
	"os"
	"testing"

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
		f, err := os.Open("./screenshot.png")
		as.Nil(err)
		defer f.Close()
		resp, _, err := AppALLPermission.Ins().File().UploadImage(ctx, &lark.UploadImageReq{
			ImageType: lark.ImageTypeMessage,
			Image:     f,
		})
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.ImageKey)
	})
}
