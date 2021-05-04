package test

import (
	"io"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_File_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		fileCli := cli.File()

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.UploadImage(ctx, &lark.UploadImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.DownloadImage(ctx, &lark.DownloadImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.UploadFile(ctx, &lark.UploadFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.DownloadFile(ctx, &lark.DownloadFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		fileCli := cli.File()

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.UploadImage(ctx, &lark.UploadImageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.DownloadImage(ctx, &lark.DownloadImageReq{
				ImageKey: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.UploadFile(ctx, &lark.UploadFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := fileCli.DownloadFile(ctx, &lark.DownloadFileReq{
				FileKey: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})
	})
}

func Test_File(t *testing.T) {
	as := assert.New(t)

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

	t.Run("", func(t *testing.T) {
		f, err := os.Open("./file_2.docx")
		as.Nil(err)
		defer f.Close()
		resp, _, err := AppALLPermission.Ins().File().UploadFile(ctx, &lark.UploadFileReq{
			FileType: lark.FileTypeDoc,
			FileName: "file2.docx",
			Duration: nil,
			File:     f,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.NotEmpty(resp.FileKey)
	})

	t.Run("", func(t *testing.T) {
		// ./test/file_2.docx
		resp, _, err := AppALLPermission.Ins().File().DownloadFile(ctx, &lark.DownloadFileReq{
			FileKey: File2.Key,
		})
		// spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		bs, err := io.ReadAll(resp.File)
		as.Nil(err)
		as.Len(bs, 3247)
	})
}
