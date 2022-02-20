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
package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_FolderExt(t *testing.T) {
	as := assert.New(t)
	ctx := context.Background()
	larkCli := AppAllPermission.Ins()
	f, err := larkext.NewRootFolder(ctx, larkCli)
	as.Nil(err)

	t.Run("meta", func(t *testing.T) {
		meta1, err := larkext.NewFolder(larkCli, "").Meta(ctx)
		as.Nil(err)

		meta2, err := f.Meta(ctx)
		as.Nil(err)

		as.Equal(meta1.Token, meta2.Token)
	})

	t.Run("list", func(t *testing.T) {
		files, err := f.ListFiles(ctx)
		as.Nil(err)
		as.True(len(files) > 0)
	})

	t.Run("new-sheet sheet-self-delete", func(t *testing.T) {
		sheet, err := f.NewSheet(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(sheet.Delete(ctx))
	})

	t.Run("new-sheet folder-delete", func(t *testing.T) {
		sheet, err := f.NewSheet(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(f.DeleteSheet(ctx, sheet.SheetToken()))
	})

	t.Run("new-doc doc-self-delete", func(t *testing.T) {
		doc, err := f.NewDoc(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(doc.Delete(ctx))
	})

	t.Run("new-doc folder-delete", func(t *testing.T) {
		doc, err := f.NewDoc(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(f.DeleteDoc(ctx, doc.DocToken()))
	})
}
