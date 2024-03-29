// Code generated by lark_sdk_gen. DO NOT EDIT.
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
	"context"
)

// BatchUpdateSheet 该接口用于根据 spreadsheetToken 操作表格, 如增加工作表, 复制工作表、删除工作表。
//
// ::: note
// 该接口和 [更新工作表属性](https://open.feishu.cn/document/ukTMukTMukTM/ugjMzUjL4IzM14COyMTN) 的请求地址相同, 但参数不同, 调用前请仔细阅读文档。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYTMzUjL2EzM14iNxMTN
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/spreadsheet-sheet/operate-sheets
func (r *DriveService) BatchUpdateSheet(ctx context.Context, request *BatchUpdateSheetReq, options ...MethodOptionFunc) (*BatchUpdateSheetResp, *Response, error) {
	if r.cli.mock.mockDriveBatchUpdateSheet != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#BatchUpdateSheet mock enable")
		return r.cli.mock.mockDriveBatchUpdateSheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "BatchUpdateSheet",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchUpdateSheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveBatchUpdateSheet mock DriveBatchUpdateSheet method
func (r *Mock) MockDriveBatchUpdateSheet(f func(ctx context.Context, request *BatchUpdateSheetReq, options ...MethodOptionFunc) (*BatchUpdateSheetResp, *Response, error)) {
	r.mockDriveBatchUpdateSheet = f
}

// UnMockDriveBatchUpdateSheet un-mock DriveBatchUpdateSheet method
func (r *Mock) UnMockDriveBatchUpdateSheet() {
	r.mockDriveBatchUpdateSheet = nil
}

// BatchUpdateSheetReq ...
type BatchUpdateSheetReq struct {
	SpreadSheetToken string                        `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Requests         []*BatchUpdateSheetReqRequest `json:"requests,omitempty"`        // 请求操作, 支持增、删、复制工作表, 三个操作选一个
}

// BatchUpdateSheetReqRequest ...
type BatchUpdateSheetReqRequest struct {
	UpdateSheet *BatchUpdateSheetReqRequestUpdateSheet `json:"updateSheet,omitempty"` // 更新工作表
	AddSheet    *BatchUpdateSheetReqRequestAddSheet    `json:"addSheet,omitempty"`    // 增加工作表
	CopySheet   *BatchUpdateSheetReqRequestCopySheet   `json:"copySheet,omitempty"`   // 复制工作表
	DeleteSheet *BatchUpdateSheetReqRequestDeleteSheet `json:"deleteSheet,omitempty"` // 删除 sheet
}

// BatchUpdateSheetReqRequestAddSheet ...
type BatchUpdateSheetReqRequestAddSheet struct {
	Properties *BatchUpdateSheetReqRequestAddSheetProperties `json:"properties,omitempty"` // 工作表属性
}

// BatchUpdateSheetReqRequestAddSheetProperties ...
type BatchUpdateSheetReqRequestAddSheetProperties struct {
	Title string `json:"title,omitempty"` // 工作表标题
	Index *int64 `json:"index,omitempty"` // 新增工作表的位置, 不填默认往前增加工作表
}

// BatchUpdateSheetReqRequestCopySheet ...
type BatchUpdateSheetReqRequestCopySheet struct {
	Source      *BatchUpdateSheetReqRequestCopySheetSource      `json:"source,omitempty"`      // 需要复制的工作表资源
	Destination *BatchUpdateSheetReqRequestCopySheetDestination `json:"destination,omitempty"` // 工作表 的属性
}

// BatchUpdateSheetReqRequestCopySheetDestination ...
type BatchUpdateSheetReqRequestCopySheetDestination struct {
	Title *string `json:"title,omitempty"` // 目标工作表名称。不填为 old_title(副本_0)
}

// BatchUpdateSheetReqRequestCopySheetSource ...
type BatchUpdateSheetReqRequestCopySheetSource struct {
	SheetID string `json:"sheetId,omitempty"` // 源 sheetId
}

// BatchUpdateSheetReqRequestDeleteSheet ...
type BatchUpdateSheetReqRequestDeleteSheet struct {
	SheetID string `json:"sheetId,omitempty"` // sheetId
}

// BatchUpdateSheetReqRequestUpdateSheet ...
type BatchUpdateSheetReqRequestUpdateSheet struct {
	Properties *BatchUpdateSheetReqRequestUpdateSheetProperties `json:"properties,omitempty"` // 工作表属性
}

// BatchUpdateSheetReqRequestUpdateSheetProperties ...
type BatchUpdateSheetReqRequestUpdateSheetProperties struct {
	SheetID        string                                                  `json:"sheetId,omitempty"`        // read-only ,作为表格唯一识别参数
	Title          *string                                                 `json:"title,omitempty"`          // 更改工作表标题
	Index          *int64                                                  `json:"index,omitempty"`          // 移动工作表的位置
	Hidden         *bool                                                   `json:"hidden,omitempty"`         // 隐藏表格，默认 false
	FrozenRowCount *int64                                                  `json:"frozenRowCount,omitempty"` // 冻结行数，小于等于工作表的最大行数，0表示取消冻结行
	FrozenColCount *int64                                                  `json:"frozenColCount,omitempty"` // 该 sheet 的冻结列数，小于等于工作表的最大列数，0表示取消冻结列
	Protect        *BatchUpdateSheetReqRequestUpdateSheetPropertiesProtect `json:"protect,omitempty"`        // 锁定表格
}

// BatchUpdateSheetReqRequestUpdateSheetPropertiesProtect ...
type BatchUpdateSheetReqRequestUpdateSheetPropertiesProtect struct {
	Lock     string   `json:"lock,omitempty"`     // LOCK 、UNLOCK 上锁/解锁
	LockInfo *string  `json:"lockInfo,omitempty"` // 锁定信息
	UserIDs  []string `json:"userIDs,omitempty"`  // 除了本人与所有者外，添加其他的可编辑人员,user_id_type不为空时使用该字段
}

// BatchUpdateSheetResp ...
type BatchUpdateSheetResp struct {
	Replies []*BatchUpdateSheetRespReply `json:"replies,omitempty"` // 返回本次相关操作工作表的结果
}

// BatchUpdateSheetRespReply ...
type BatchUpdateSheetRespReply struct {
	UpdateSheet *BatchUpdateSheetRespReplyUpdateSheet `json:"updateSheet,omitempty"` // 更新工作表
	AddSheet    *BatchUpdateSheetRespReplyAddSheet    `json:"addSheet,omitempty"`    // 增加/复制工作表的属性
	CopySheet   *BatchUpdateSheetRespReplyCopySheet   `json:"copySheet,omitempty"`   // 增加/复制工作表的属性
	DeleteSheet *BatchUpdateSheetRespReplyDeleteSheet `json:"deleteSheet,omitempty"` // 删除工作表
}

// BatchUpdateSheetRespReplyAddSheet ...
type BatchUpdateSheetRespReplyAddSheet struct {
	Properties *BatchUpdateSheetRespReplyAddSheetProperties `json:"properties,omitempty"` // 表格属性
}

// BatchUpdateSheetRespReplyAddSheetProperties ...
type BatchUpdateSheetRespReplyAddSheetProperties struct {
	SheetID string `json:"sheetId,omitempty"` // sheetId
	Title   string `json:"title,omitempty"`   // 工作表标题
	Index   int64  `json:"index,omitempty"`   // 工作表位置
}

// BatchUpdateSheetRespReplyCopySheet ...
type BatchUpdateSheetRespReplyCopySheet struct {
	Properties *BatchUpdateSheetRespReplyCopySheetProperties `json:"properties,omitempty"` // 表格属性
}

// BatchUpdateSheetRespReplyCopySheetProperties ...
type BatchUpdateSheetRespReplyCopySheetProperties struct {
	SheetID string `json:"sheetId,omitempty"` // sheetId
	Title   string `json:"title,omitempty"`   // 工作表标题
	Index   int64  `json:"index,omitempty"`   // 工作表位置
}

// BatchUpdateSheetRespReplyDeleteSheet ...
type BatchUpdateSheetRespReplyDeleteSheet struct {
	Result  bool   `json:"result,omitempty"`  // 删除工作表是否成功
	SheetID string `json:"sheetId,omitempty"` // sheetId
}

// BatchUpdateSheetRespReplyUpdateSheet ...
type BatchUpdateSheetRespReplyUpdateSheet struct {
	Properties *BatchUpdateSheetRespReplyUpdateSheetProperties `json:"properties,omitempty"` // 工作表属性
}

// BatchUpdateSheetRespReplyUpdateSheetProperties ...
type BatchUpdateSheetRespReplyUpdateSheetProperties struct {
	SheetID        string                                                 `json:"sheetId,omitempty"`        // read-only ,作为表格唯一识别参数
	Title          string                                                 `json:"title,omitempty"`          // 更改工作表标题
	Index          int64                                                  `json:"index,omitempty"`          // 移动工作表的位置
	Hidden         bool                                                   `json:"hidden,omitempty"`         // 隐藏表格，默认 false
	FrozenRowCount int64                                                  `json:"frozenRowCount,omitempty"` // 冻结行数，小于等于工作表的最大行数，0表示取消冻结行
	FrozenColCount int64                                                  `json:"frozenColCount,omitempty"` // 该 sheet 的冻结列数，小于等于工作表的最大列数，0表示取消冻结列
	Protect        *BatchUpdateSheetRespReplyUpdateSheetPropertiesProtect `json:"protect,omitempty"`        // 锁定表格
}

// BatchUpdateSheetRespReplyUpdateSheetPropertiesProtect ...
type BatchUpdateSheetRespReplyUpdateSheetPropertiesProtect struct {
	Lock     string   `json:"lock,omitempty"`     // LOCK 、UNLOCK 上锁/解锁
	LockInfo string   `json:"lockInfo,omitempty"` // 锁定信息
	UserIDs  []string `json:"userIDs,omitempty"`  // 除了本人与所有者外，添加其他的可编辑人员,user_id_type不为空时使用该字段
}

// batchUpdateSheetResp ...
type batchUpdateSheetResp struct {
	Code int64                 `json:"code,omitempty"`
	Msg  string                `json:"msg,omitempty"`
	Data *BatchUpdateSheetResp `json:"data,omitempty"`
}
