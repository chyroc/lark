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
	"fmt"
)

// ItoCol int to col
// index >= 1
func ItoCol(index int) string {
	var (
		columnBase = 26
		digitMax   = 7 // ceil(log26(Int32.Max))
		digits     = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	)

	if index <= 0 {
		return ""
	}

	if index <= columnBase {
		return string([]byte{digits[index-1]})
	}

	current := index
	offset := digitMax
	res := make([]byte, digitMax)
	for current > 0 {
		offset -= 1
		current -= 1
		res[offset] = digits[current%columnBase]
		current /= columnBase
	}

	return string(res[offset:])
}

// CellRange 将两个坐标 (x1, y1) 和 (x2, y2) 转成 sheet!A1:D5 的形式
//
// https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview
// https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-guide
//
// a. sheetId ：表示整表应用
// b. sheetId!1:2 ：表示整行应用
// c. sheetId!A:B ：表示整列应用
// d. sheetId!A1:C ：省略结束行，会使用表格的最后行作为结束行
// e. sheetId!A1:B2 ：表示普通的range
func CellRange(sheetID string, x1, y1, x2, y2 int) string {
	// a
	if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
		return sheetID
	}
	// b
	if x1 == 0 && x2 == 0 {
		return fmt.Sprintf("%s!%d:%d", sheetID, y1, y2)
	}
	// c
	if y1 == 0 && y2 == 0 {
		return fmt.Sprintf("%s!%s:%s", sheetID, ItoCol(x1), ItoCol(x2))
	}
	// d
	if y2 == 0 {
		return fmt.Sprintf("%s!%s%d:%s", sheetID, ItoCol(x1), y1, ItoCol(x2))
	}
	// e
	return fmt.Sprintf("%s!%s%d:%s%d", sheetID, ItoCol(x1), y1, ItoCol(x2), y2)
}
