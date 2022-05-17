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
package doc

import (
	"github.com/chyroc/lark"
)

// TableBlock 格式化表格
// 这个有问题，可以创建，但是无内容
func TableBlock(items ...*lark.DocTableRow) *lark.DocBlock {
	rowSize := int64(len(items))
	columnSize := int64(0)
	for _, row := range items {
		if int64(len(row.TableCells)) > columnSize {
			columnSize = int64(len(row.TableCells))
		}
	}
	return &lark.DocBlock{
		Type: lark.DocBlockTypeTable,
		Table: &lark.DocTable{
			RowSize:    rowSize,
			ColumnSize: columnSize,
			TableRows:  items,
		},
	}
}

// EmptyTableBlock 格式化表格
func EmptyTableBlock(rowSize, columnSize int64) *lark.DocBlock {
	return &lark.DocBlock{
		Type: lark.DocBlockTypeTable,
		Table: &lark.DocTable{
			RowSize:    rowSize,
			ColumnSize: columnSize,
		},
	}
}
