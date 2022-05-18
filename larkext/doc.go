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

// Doc is doc client
type Doc struct {
	larkClient *lark.Lark
	docToken   string
	docURL     string
	typ        string
}

// NewDoc new doc client
func NewDoc(larkClient *lark.Lark, docToken string) *Doc {
	return newDoc(larkClient, docToken, "")
}

func newDoc(larkClient *lark.Lark, docToken, docURL string) *Doc {
	r := new(Doc)
	r.larkClient = larkClient
	r.docToken = docToken
	r.docURL = docURL
	r.typ = "doc"
	return r
}

// DocToken get doc Token
func (r *Doc) DocToken() string {
	return r.docToken
}

// DocURL get doc URL
func (r *Doc) DocURL() string {
	return r.docURL
}

// Meta get doc meta
func (r *Doc) Meta(ctx context.Context) (*lark.GetDriveDocMetaResp, error) {
	return r.meta(ctx)
}

// Copy copy doc file
func (r *Doc) Copy(ctx context.Context, folderToken, name string) (*Doc, error) {
	return r.copy(ctx, folderToken, name)
}

func (r *Doc) Move(ctx context.Context, folderToken string) (*Task, error) {
	return moveFile(ctx, r.larkClient, folderToken, r.docToken, r.typ)
}

func (r *Doc) Delete(ctx context.Context) (*Task, error) {
	return deleteFile(ctx, r.larkClient, r.docToken, r.typ)
}

// RawContent get doc raw content
func (r *Doc) RawContent(ctx context.Context) (string, error) {
	return r.rawContent(ctx)
}

// RawContent get doc rich content
func (r *Doc) Content(ctx context.Context) (*lark.DocContent, error) {
	return r.content(ctx)
}

// Update update doc
func (r *Doc) Update(ctx context.Context, requests ...*lark.UpdateDocRequest) error {
	return r.update(ctx, requests...)
}

// Revision get doc revision
func (r *Doc) Revision(ctx context.Context) (int64, error) {
	return r.revision(ctx)
}

// Permission grant doc permission
func (r *Doc) Permission() *Permission {
	return newPermission(r.larkClient, r.docToken, r.typ)
}

// GetParagraphElementLocation get paragraph element location
func (r *Doc) GetParagraphElementLocation(ctx context.Context, pe *lark.DocParagraphElement) (*lark.DocLocation, error) {
	content, err := r.content(ctx)
	if err != nil {
		return nil, err
	}

	return content.GetParagraphElementLocation(pe), nil
}

func (r *Doc) GetParagraph(ctx context.Context, p *lark.DocParagraph) (*lark.DocParagraph, error) {
	content, err := r.content(ctx)
	if err != nil {
		return nil, err
	}

	return content.GetParagraph(p), nil
}

func (r *Doc) GetTableBySize(ctx context.Context, p *lark.DocTable) (*lark.DocTable, error) {
	content, err := r.content(ctx)
	if err != nil {
		return nil, err
	}

	return content.GetTableBySize(p), nil
}
