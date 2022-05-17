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
package lark

import (
	"encoding/json"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uYDM2YjL2AjN24iNwYjN

type UpdateDocRequests []*UpdateDocRequest

func (udr UpdateDocRequests) ToString() []string {
	res := []string{}
	for _, v := range udr {
		bs, _ := json.Marshal(v)
		res = append(res, string(bs))
	}
	return res
}

// UpdateDocRequest ...
type UpdateDocRequest struct {
	RequestType                        UpdateDocRequestType                         `json:"requestType"`
	UpdateTitleRequest                 *UpdateDocUpdateTitleRequest                 `json:"updateTitleRequest,omitempty"`                 // 更新标题
	DeleteContentRangeRequest          *UpdateDocDeleteContentRangeRequest          `json:"deleteContentRangeRequest,omitempty"`          // 范围删除
	InsertBlocksRequest                *UpdateDocInsertBlocksRequest                `json:"insertBlocksRequest,omitempty"`                // 插入 blocks
	InsertParagraphElementsRequest     *UpdateDocInsertParagraphElementsRequest     `json:"insertParagraphElementsRequest,omitempty"`     // 插入行内元素
	InsertTableRowRequest              *UpdateDocInsertTableRowRequest              `json:"insertTableRowRequest,omitempty"`              // 格式化表格增加单行
	InsertTableColumnRequest           *UpdateDocInsertTableColumnRequest           `json:"insertTableColumnRequest,omitempty"`           // 格式化表格增加单列
	DeleteTableRowsRequest             *UpdateDocDeleteTableRowsRequest             `json:"deleteTableRowsRequest,omitempty"`             // 格式化表格删除多行
	DeleteTableColumnsRequest          *UpdateDocDeleteTableColumnsRequest          `json:"deleteTableColumnsRequest,omitempty"`          // 格式化表格删除多列
	UpdateTableColumnPropertiesRequest *UpdateDocUpdateTableColumnPropertiesRequest `json:"updateTableColumnPropertiesRequest,omitempty"` // 格式化表格修改列宽度
	MergeTableCellsRequest             *UpdateDocMergeTableCellsRequest             `json:"mergeTableCellsRequest,omitempty"`             // 格式化表格合并单元格
	UnmergeTableCellsRequest           *UpdateDocUnmergeTableCellsRequest           `json:"unmergeTableCellsRequest,omitempty"`           // 格式化表格拆分单元格
	ReplaceAllTextRequest              *UpdateDocReplaceAllTextRequest              `json:"replaceAllTextRequest,omitempty"`              // 查找替换文本内容
	UpdateParagraphStyleRequest        *UpdateDocUpdateParagraphStyleRequest        `json:"updateParagraphStyleRequest,omitempty"`        // 更新段落样式
}

func (r *UpdateDocRequest) SetReplaceAllTextRequestContainsTextMatchCase(val bool) *UpdateDocRequest {
	if r.ReplaceAllTextRequest != nil && r.ReplaceAllTextRequest.ContainsText != nil {
		r.ReplaceAllTextRequest.ContainsText.MatchCase = val
	}
	return r
}

// UpdateDocRequestType ...
type UpdateDocRequestType string

// UpdateDocRequestTypeUpdateTitle ...
const (
	UpdateDocRequestTypeUpdateTitle                 UpdateDocRequestType = "UpdateTitleRequestType"                 //  更新标题
	UpdateDocRequestTypeDeleteContentRange          UpdateDocRequestType = "DeleteContentRangeRequestType"          //  范围删除
	UpdateDocRequestTypeInsertBlocks                UpdateDocRequestType = "InsertBlocksRequestType"                //  插入 blocks
	UpdateDocRequestTypeInsertParagraphElements     UpdateDocRequestType = "InsertParagraphElementsRequestType"     //  插入行内元素
	UpdateDocRequestTypeInsertTableRow              UpdateDocRequestType = "InsertTableRowRequestType"              //  格式化表格增加单行
	UpdateDocRequestTypeInsertTableColumn           UpdateDocRequestType = "InsertTableColumnRequestType"           //  格式化表格增加单列
	UpdateDocRequestTypeDeleteTableRows             UpdateDocRequestType = "DeleteTableRowsRequestType"             //  格式化表格删除多行
	UpdateDocRequestTypeDeleteTableColumns          UpdateDocRequestType = "DeleteTableColumnsRequestType"          //  格式化表格删除多列
	UpdateDocRequestTypeUpdateTableColumnProperties UpdateDocRequestType = "UpdateTableColumnPropertiesRequestType" //  格式化表格修改列宽度
	UpdateDocRequestTypeMergeTableCells             UpdateDocRequestType = "MergeTableCellsRequestType"             //  格式化表格合并单元格
	UpdateDocRequestTypeUnmergeTableCells           UpdateDocRequestType = "UnmergeTableCellsRequestType"           //  格式化表格拆分单元格
	UpdateDocRequestTypeReplaceAllText              UpdateDocRequestType = "ReplaceAllTextRequestType"              //  查找替换文本内容
	UpdateDocRequestTypeUpdateParagraphStyle        UpdateDocRequestType = "UpdateParagraphStyleRequestType"        //  更新段落样式
)

// UpdateDocUpdateTitleRequest 更新标题
type UpdateDocUpdateTitleRequest struct {
	Payload string `json:"payload"` // 文档数据结构概述 Body结构体序列化string
}

// UpdateDocDeleteContentRangeRequest 范围删除
type UpdateDocDeleteContentRangeRequest struct {
	DeleteRange *DocLocation `json:"deleteRange"`
}

// UpdateDocInsertBlocksRequest 插入 blocks
type UpdateDocInsertBlocksRequest struct {
	Payload  string             `json:"payload,omitempty"`  // 文档数据结构概述 Body结构体序列化string
	Location *UpdateDocLocation `json:"location,omitempty"` // 插入位置
}

// UpdateDocInsertParagraphElementsRequest 插入行内元素
type UpdateDocInsertParagraphElementsRequest struct {
	Payload  string             `json:"payload"`  // 文档数据结构概述 Body结构体序列化string
	Location *UpdateDocLocation `json:"location"` // 插入位置
}

// UpdateDocInsertTableRowRequest 格式化表格增加单行
type UpdateDocInsertTableRowRequest struct {
	TableID  string `json:"tableId"`  // 表格 id, 详见 文档数据结构概述 Table 结构体定义
	RowIndex int64  `json:"rowIndex"` // 插入行索引，从 0 开始，往第一行前插入使用 0
}

// UpdateDocInsertTableColumnRequest 格式化表格增加单列
type UpdateDocInsertTableColumnRequest struct {
	TableID     string `json:"tableId"`     // 表格 id, 详见 文档数据结构概述 Table 结构体定义
	ColumnIndex int64  `json:"columnIndex"` // 	插入列索引，从 0 开始，往第一列前插入使用 0
}

// UpdateDocDeleteTableRowsRequest 格式化表格删除多行
// 删除范围 [RowStartIndex,RowEndIndex)
type UpdateDocDeleteTableRowsRequest struct {
	TableID       string `json:"tableId"`       // 表格id, 详见 文档数据结构概述 Table结构体定义
	RowStartIndex int64  `json:"rowStartIndex"` // 删除行起始索引，从 0 开始
	RowEndIndex   int64  `json:"rowEndIndex"`   // 删除行结束索引
}

// UpdateDocDeleteTableColumnsRequest 格式化表格删除多列
// 删除范围 [ColumnStartIndex,ColumnEndIndex)
type UpdateDocDeleteTableColumnsRequest struct {
	TableID          string `json:"tableId"`          //	表格id, 详见 文档数据结构概述 Table结构体定义
	ColumnStartIndex int64  `json:"columnStartIndex"` // 删除列起始索引，从 0 开始
	ColumnEndIndex   int64  `json:"columnEndIndex"`   //	删除列结束索引
}

// UpdateDocUpdateTableColumnPropertiesRequest 格式化表格修改列宽度
type UpdateDocUpdateTableColumnPropertiesRequest struct {
	TableID     string `json:"tableId"`     // 表格id, 详见 文档数据结构概述 Table结构体定义
	ColumnIndex int64  `json:"columnIndex"` // 列索引，从 0 开始
	ColumnWidth int64  `json:"columnWidth"` // 列宽度，单位 px
}

// UpdateDocMergeTableCellsRequest 格式化表格合并单元格
type UpdateDocMergeTableCellsRequest struct {
	TableID          string `json:"tableId"`          // 表格 id, 详见 文档数据结构概述 Table 结构体定义
	RowStartIndex    int64  `json:"rowStartIndex"`    // 合并行起始索引，从 0 开始
	RowEndIndex      int64  `json:"rowEndIndex"`      // 合并行结束索引
	ColumnStartIndex int64  `json:"columnStartIndex"` // 合并列起始索引，从 0 开始
	ColumnEndIndex   int64  `json:"columnEndIndex"`   // 合并列结束索引
}

// UpdateDocUnmergeTableCellsRequest 格式化表格拆分单元格
type UpdateDocUnmergeTableCellsRequest struct {
	TableID      string `json:"tableId"`      // 表格 id, 详见 文档数据结构概述 Table结构体定义
	MergedCellID string `json:"mergedCellId"` // 合并单元格 id
}

// UpdateDocReplaceAllTextRequest 查找替换文本内容
type UpdateDocReplaceAllTextRequest struct {
	ReplaceText  string                           `json:"replaceText"`  // 替换文字，可为空，不可包含换行符、tab，最多1000个字
	ContainsText *UpdateDocSubstringMatchCriteria `json:"containsText"` // 查找文字
}

// UpdateDocUpdateParagraphStyleRequest 更新段落样式
type UpdateDocUpdateParagraphStyleRequest struct {
	Payload string              `json:"payload"` //	详见文档数据结构参考 UpdateParagraphStyleRequest 子命令使用 ParagraphStyle 结构体序列化 string
	Range   *DocLocation        `json:"range"`   // 修改范围
	Fields  *UpdateDocFieldMask `json:"fields"`  // 修改字段
}

// UpdateDocLocation ...
type UpdateDocLocation struct {
	ZoneID      string `json:"zoneId"`      // 编辑区域，包括两种，正文是 "0"，表格单元格 ID，例如 "xr1m4jw7egd9nefz1s0mdsetenl5fbe3lygxc1azupv81i5t2rjmosw5ta0esuwtn8ksya"
	Index       int64  `json:"index"`       // 字符下标，标题占用索引，从 0 开始
	StartOfZone bool   `json:"startOfZone"` // InsertBlocksRequestType支持便捷操作，插入zone开始位置，正文头部（不算标题）或单元格起始位置。true时index入参无效
	EndOfZone   bool   `json:"endOfZone"`   // InsertBlocksRequestType支持便捷操作，插入zone尾部，正文尾部或单元格尾部。true时index入参无效
}

// === UpdateDocLocation ===

func (r *UpdateDocLocation) SetZoneID(val string) *UpdateDocLocation {
	if val == "" {
		val = "0"
	}
	r.ZoneID = val
	return r
}

func (r *UpdateDocLocation) SetIndex(val int64) *UpdateDocLocation {
	r.Index = val
	return r
}

func (r *UpdateDocLocation) SetStartOfZone(val bool) *UpdateDocLocation {
	r.StartOfZone = val
	return r
}

func (r *UpdateDocLocation) SetEndOfZone(val bool) *UpdateDocLocation {
	r.EndOfZone = val
	return r
}

// === UpdateDocLocation ===

// UpdateDocSubstringMatchCriteria ...
type UpdateDocSubstringMatchCriteria struct {
	Text      string `json:"text"`      // 查找内容，不可为空，不可包含换行、tab，最多1000个字
	MatchCase bool   `json:"matchCase"` // 是否匹配大小写，默认false，忽略大小写
}

// UpdateDocFieldMask ...
type UpdateDocFieldMask struct {
	Masks []UpdateDocMaskEnum `json:"masks"` // 修改字段，包括"list"、"headingLevel"、"align"、"quote"
}

type UpdateDocMaskEnum string

const (
	UpdateDocMaskEnumList         UpdateDocMaskEnum = "list"
	UpdateDocMaskEnumHeadingLevel UpdateDocMaskEnum = "headingLevel"
	UpdateDocMaskEnumAlign        UpdateDocMaskEnum = "align"
	UpdateDocMaskEnumQuote        UpdateDocMaskEnum = "quote"
)
