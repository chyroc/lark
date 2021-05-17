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
		resp, _, err := AppALLPermission.Ins().File.UploadImage(ctx, &lark.UploadImageReq{
			ImageType: lark.ImageTypeMessage,
			Image:     f,
		})
		printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.ImageKey)
	})

	t.Run("", func(t *testing.T) {
		// this is file of ./test/file_1.png
		resp, _, err := AppALLPermission.Ins().File.DownloadImage(ctx, &lark.DownloadImageReq{
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
		resp, _, err := AppALLPermission.Ins().File.UploadFile(ctx, &lark.UploadFileReq{
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
		resp, _, err := AppALLPermission.Ins().File.DownloadFile(ctx, &lark.DownloadFileReq{
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
