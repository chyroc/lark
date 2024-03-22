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

func newFile(larkClient *lark.Lark, token, url string) *File {
	r := new(File)
	r.larkClient = larkClient
	r.token = token
	r.url = url
	r.typ = "file"
	return r
}

func (r *Docx) copy(ctx context.Context, folderToken, name string) (*Docx, error) {
	res, err := copyFile(ctx, r.larkClient, folderToken, r.token, r.typ, name)
	if err != nil {
		return nil, err
	}
	return newDocx(r.larkClient, res.Token, res.URL), nil
}

func (r *Docx) rawContent(ctx context.Context) (string, error) {
	resp, _, err := r.larkClient.Drive.GetDocxDocumentRawContent(ctx, &lark.GetDocxDocumentRawContentReq{
		DocumentID: r.token,
		Lang:       nil,
	})
	if err != nil {
		return "", err
	}
	return resp.Content, nil
}

func (r *Docx) blocks(ctx context.Context) ([]*lark.DocxBlock, error) {
	token := ""
	size := int64(200)
	blocks := []*lark.DocxBlock{}
	for {
		resp, _, err := r.larkClient.Drive.GetDocxBlockListOfDocument(ctx, &lark.GetDocxBlockListOfDocumentReq{
			DocumentID:         r.token,
			PageSize:           &size,
			PageToken:          &token,
			DocumentRevisionID: nil,
			UserIDType:         nil,
		})
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, resp.Items...)
		if !resp.HasMore {
			break
		}
		token = resp.PageToken
	}
	return blocks, nil
}
