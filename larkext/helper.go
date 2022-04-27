package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

// file：文件类型
// doc：云文档类型
// sheet：电子表格类型
// bitable：多维表格类型
// docx：新版云文档类型
// mindnote：思维笔记类型
func copyFile(ctx context.Context, larkClient *lark.Lark, folderToken, fileToken, typ, name string) (*FileMeta, error) {
	resp, _, err := larkClient.Drive.CopyDriveFile(ctx, &lark.CopyDriveFileReq{
		FileToken:   fileToken,
		Name:        name,
		Type:        &typ,
		FolderToken: folderToken,
	})
	if err != nil {
		return nil, err
	}
	return &FileMeta{
		Token:       resp.File.Token,
		Name:        resp.File.Name,
		Type:        resp.File.Type,
		ParentToken: resp.File.ParentToken,
		URL:         resp.File.URL,
	}, nil
}
