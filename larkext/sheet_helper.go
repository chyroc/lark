package larkext

import (
	"fmt"
	"strconv"
)

// int to col
// index >= 1
func ItoCol(index int) string {
	var (
		columnBase = 26
		digitMax   = 7 // ceil(log26(Int32.Max))
		digits     = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	)

	if index <= 0 {
		panic("index <= 0")
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

// CellPoint 左上角第一个单元格坐标是 (1, 1) ，x 表示横坐标，y 表示纵坐标，CellPoint 会将 (1, 1) 转成 A1
//
// https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview
//
// range 的描述方式为 <sheetId>!<开始位置>:<结束位置> ，共有 4 种描述方法，分别为：
// 	a. <sheetId>!<开始单元格>:<结束单元格> 如：0b**12!A1:B5 就表示 0b**12 这个工作表中 A1:B5 的区域
// 	b. <sheetId>!<开始列>:<结束列>，如：0b**12!A:B
// 	c. <sheetId>!<开始单元格>:<结束列>，如：0b**12!A1:B
// 	d. <sheetId>，区域留空，如：0b**12，代表这个表格中非空的最大行列范围内的数据
func CellPoint(x, y int) string {
	// d 情况
	if x <= 0 {
		return ""
	}

	// b 情况
	if y == 0 {
		return ItoCol(x) // A, Z, AA
	}

	// a，c 情况
	return ItoCol(x) + strconv.FormatInt(int64(y), 10)
}

// CellRange 将两个坐标 (x1, y1) 和 (x2, y2) 转成 sheet!A1:D5 的形式
//
// https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview
//
// range 的描述方式为 <sheetId>!<开始位置>:<结束位置> ，共有 4 种描述方法，分别为：
// 	a. <sheetId>!<开始单元格>:<结束单元格> 如：0b**12!A1:B5 就表示 0b**12 这个工作表中 A1:B5 的区域
// 	b. <sheetId>!<开始列>:<结束列>，如：0b**12!A:B
// 	c. <sheetId>!<开始单元格>:<结束列>，如：0b**12!A1:B
// 	d. <sheetId>，区域留空，如：0b**12，代表这个表格中非空的最大行列范围内的数据
func CellRange(sheetID string, x1, y1, x2, y2 int) string {
	// d
	if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
		return sheetID
	}
	if x1 == 0 {
		return ""
	}
	// c
	if y1 != 0 && y2 == 0 {
		return fmt.Sprintf("%s!%s%d:%s", sheetID, ItoCol(x1), y1, ItoCol(x2))
	}
	// b
	if y1 == 0 && y2 == 0 {
		return fmt.Sprintf("%s!%s:%s", sheetID, ItoCol(x1), ItoCol(x2))
	}
	if x2 == 0 {
		return ""
	}
	// a
	return fmt.Sprintf("%s!%s%d:%s%d", sheetID, ItoCol(x1), y1, ItoCol(x2), y2)
}
