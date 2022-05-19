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

// File is file client
type File struct {
	larkClient *lark.Lark
	token      string
	url        string
	typ        string
}

// NewDoc new file client
func NewFile(larkClient *lark.Lark, token string) *File {
	return newFile(larkClient, token, "")
}

// Copy copy file
func (r *File) Copy(ctx context.Context, folderToken, name string) (*File, error) {
	return r.copy(ctx, folderToken, name)
}

// Move ...
func (r *File) Move(ctx context.Context, folderToken string) error {
	_, err := moveFile(ctx, r.larkClient, folderToken, r.token, r.typ)
	return err
}

// Delete ...
func (r *File) Delete(ctx context.Context) error {
	_, err := deleteFile(ctx, r.larkClient, r.token, r.typ)
	return err
}

// Download ...
func (r *File) Download(ctx context.Context) (*lark.DownloadDriveFileResp, error) {
	resp, _, err := r.larkClient.Drive.DownloadDriveFile(ctx, &lark.DownloadDriveFileReq{
		FileToken: r.token,
	})
	return resp, err
}
