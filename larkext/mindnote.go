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

// Mindnote is mindnote client
type Mindnote struct {
	larkClient *lark.Lark
	token      string
	url        string
	typ        string
}

// NewDoc new doc client
func NewMindnote(larkClient *lark.Lark, docToken string) *Mindnote {
	return newMindnote(larkClient, docToken, "")
}

func newMindnote(larkClient *lark.Lark, token, url string) *Mindnote {
	r := new(Mindnote)
	r.larkClient = larkClient
	r.token = token
	r.url = url
	r.typ = "mindnote"
	return r
}

// Move move file
func (r *Mindnote) Move(ctx context.Context, folderToken string) (*Task, error) {
	return moveFile(ctx, r.larkClient, folderToken, r.token, r.typ)
}

func (r *Mindnote) Delete(ctx context.Context) (*Task, error) {
	return deleteFile(ctx, r.larkClient, r.token, r.typ)
}
