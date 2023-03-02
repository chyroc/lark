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
func copyFile(ctx context.Context, larkClient *lark.Lark, folderToken, fileToken string, typ string, name string) (*FileMeta, error) {
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

func moveFile(ctx context.Context, larkClient *lark.Lark, folderToken, fileToken string, typ string) (*Task, error) {
	resp, _, err := larkClient.Drive.MoveDriveFile(ctx, &lark.MoveDriveFileReq{
		FileToken:   fileToken,
		Type:        &typ,
		FolderToken: &folderToken,
	})
	if err != nil {
		return nil, err
	}
	return newTask(larkClient, resp.TaskID), nil
}

// TODO: shortcut
func deleteFile(ctx context.Context, larkClient *lark.Lark, fileToken string, typ string) (*Task, error) {
	resp, _, err := larkClient.Drive.DeleteDriveFile(ctx, &lark.DeleteDriveFileReq{
		FileToken: fileToken,
		Type:      typ,
	})
	if err != nil {
		return nil, err
	}
	return newTask(larkClient, resp.TaskID), nil
}

func ptrStringNoNonePtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ptrString(s string) *string {
	return &s
}

func ptrInt64(s int64) *int64 {
	return &s
}

func ptrBool(s bool) *bool {
	return &s
}
