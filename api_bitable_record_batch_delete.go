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

// BatchDeleteBitableRecord 该接口用于删除数据表中现有的多条记录, 单次调用中最多删除 500 条记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_delete
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-record/batch_delete
func (r *BitableService) BatchDeleteBitableRecord(ctx context.Context, request *BatchDeleteBitableRecordReq, options ...MethodOptionFunc) (*BatchDeleteBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableBatchDeleteBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchDeleteBitableRecord mock enable")
		return r.cli.mock.mockBitableBatchDeleteBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "BatchDeleteBitableRecord",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchDeleteBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableBatchDeleteBitableRecord mock BitableBatchDeleteBitableRecord method
func (r *Mock) MockBitableBatchDeleteBitableRecord(f func(ctx context.Context, request *BatchDeleteBitableRecordReq, options ...MethodOptionFunc) (*BatchDeleteBitableRecordResp, *Response, error)) {
	r.mockBitableBatchDeleteBitableRecord = f
}

// UnMockBitableBatchDeleteBitableRecord un-mock BitableBatchDeleteBitableRecord method
func (r *Mock) UnMockBitableBatchDeleteBitableRecord() {
	r.mockBitableBatchDeleteBitableRecord = nil
}

// BatchDeleteBitableRecordReq ...
type BatchDeleteBitableRecordReq struct {
	AppToken string   `path:"app_token" json:"-"` // base app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string   `path:"table_id" json:"-"`  // table id, 示例值: "tblsRc9GRRXKqhvW"
	Records  []string `json:"records,omitempty"`  // 删除的多条记录id列表, 示例值: ["recIcJBbvC", "recvmiCORa"]
}

// BatchDeleteBitableRecordResp ...
type BatchDeleteBitableRecordResp struct {
	Records []*BatchDeleteBitableRecordRespRecord `json:"records,omitempty"` // 记录
}

// BatchDeleteBitableRecordRespRecord ...
type BatchDeleteBitableRecordRespRecord struct {
	Deleted  bool   `json:"deleted,omitempty"`   // 是否成功删除
	RecordID string `json:"record_id,omitempty"` // 删除的记录 ID
}

// batchDeleteBitableRecordResp ...
type batchDeleteBitableRecordResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeleteBitableRecordResp `json:"data,omitempty"`
}
