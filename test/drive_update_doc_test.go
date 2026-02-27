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
	"math/rand"
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/doc"
	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func randString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Test_UpdateDoc(t *testing.T) {
	SkipRealAPITest(t)
	as := assert.New(t)

	ctx := context.Background()
	client := AppAllPermission.Ins()

	f, err := larkext.NewRootFolder(ctx, client)
	if !as.Nil(err) {
		return
	}

	docIns, err := f.NewDoc(ctx, "docs")
	if err != nil && lark.GetErrorCode(err) == 95054 {
		t.Skipf("skip because CreateDriveDoc is deprecated: %v", err)
		return
	}
	if !as.Nil(err) || !as.NotNil(docIns) {
		return
	}
	fmt.Println(docIns.DocURL())

	// 最后删除这个 doc
	defer docIns.Delete(ctx)

	// 更新标题
	as.Nil(docIns.Update(ctx,
		doc.UpdateTitle("[test] update doc"),
	))

	// 在正文最后面，插入文字
	text1 := randString(20)
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocInsertBlocksRequest(doc.Body(
			doc.ParagraphBlock(
				doc.TextRunParagraphElement(text1),
			),
		), doc.EndOfZoneDocLocation()),
	))

	// 找到这个文本的位置
	text1Location, err := docIns.GetParagraphElementLocation(ctx, doc.TextRunParagraphElement(text1))
	if !as.Nil(err) {
		return
	}

	// 插入一段行内元素
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocInsertParagraphElementsRequest(
			doc.Paragraph(
				doc.TextRunParagraphElement("1之后的行内元素"),
			),
			doc.IndexDocLocation(text1Location.EndIndex),
		),
	))

	// 元素，改为素元
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocReplaceAllTextRequest("元素", "素元"),
	))

	text1 = text1 + "1之后的行内素元"
	paragraph, err := docIns.GetParagraph(ctx, doc.Paragraph(
		doc.TextRunParagraphElement(text1),
	))
	if !as.Nil(err) {
		return
	}

	// 更新段落
	as.Nil(docIns.Update(ctx,
		doc.UpdateParagraphStyleRequest(
			doc.Location(paragraph.Location.StartIndex, paragraph.Location.EndIndex),
			doc.SetParagraphStyleQuote(true),
		),
	))

	// 插入表格
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocInsertBlocksRequest(doc.Body(
			doc.EmptyTableBlock(2, 2),
		), doc.EndOfZoneDocLocation()),
	))

	tableSize2, err := docIns.GetTableBySize(ctx, doc.EmptyTableBlock(2, 2).Table)
	if !as.Nil(err) {
		return
	}

	// 插入一行
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocInsertTableRowRequest(
			tableSize2.TableID, 1,
		),
	))

	// 插入一列
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocInsertTableColumnRequest(
			tableSize2.TableID, 1,
		),
	))

	// 删除一行
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocDeleteTableRowRequest(
			tableSize2.TableID, 0, 1,
		),
	))

	// 删除一列
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocDeleteTableColumnsRequest(
			tableSize2.TableID, 0, 1,
		),
	))

	// 修改第一列宽度
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocUpdateTableColumnPropertiesRequest(
			tableSize2.TableID,
			0,
			300,
		),
	))

	// 合并表格
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocMergeTableCellsRequest(
			tableSize2.TableID,
			0, 2,
			0, 2,
		),
	))

	tableSize2, err = docIns.GetTableBySize(ctx, doc.EmptyTableBlock(2, 2).Table)
	if !as.Nil(err) || !as.NotNil(tableSize2) {
		return
	}
	if !as.NotEmpty(tableSize2.MergedCells) {
		return
	}

	// 取消表格
	as.Nil(docIns.Update(ctx,
		doc.UpdateDocUnmergeTableCellsRequest(
			tableSize2.TableID,
			tableSize2.MergedCells[0].MergedCellID,
		),
	))
}
