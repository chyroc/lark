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

// GetBitableRecord 该接口用于根据 record_id 的值检索现有记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-record/get
func (r *BitableService) GetBitableRecord(ctx context.Context, request *GetBitableRecordReq, options ...MethodOptionFunc) (*GetBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableRecord mock enable")
		return r.cli.mock.mockBitableGetBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableRecord",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableRecord mock BitableGetBitableRecord method
func (r *Mock) MockBitableGetBitableRecord(f func(ctx context.Context, request *GetBitableRecordReq, options ...MethodOptionFunc) (*GetBitableRecordResp, *Response, error)) {
	r.mockBitableGetBitableRecord = f
}

// UnMockBitableGetBitableRecord un-mock BitableGetBitableRecord method
func (r *Mock) UnMockBitableGetBitableRecord() {
	r.mockBitableGetBitableRecord = nil
}

// GetBitableRecordReq ...
type GetBitableRecordReq struct {
	AppToken          string  `path:"app_token" json:"-"`            // base app token, 示例值: "bascnCMII2ORej2RItqpZZUNMIe"
	TableID           string  `path:"table_id" json:"-"`             // table id, 示例值: "tblxI2tWaxP5dG7p"
	RecordID          string  `path:"record_id" json:"-"`            // 单条记录的 id, 示例值: "recn0hoyXL"
	TextFieldAsArray  *bool   `query:"text_field_as_array" json:"-"` // 控制多行文本字段数据的返回格式, true 表示以数组形式返回, 示例值: true
	UserIDType        *IDType `query:"user_id_type" json:"-"`        // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DisplayFormulaRef *bool   `query:"display_formula_ref" json:"-"` // 控制公式、查找引用是否显示完整的原样返回结果, 示例值: true
	AutomaticFields   *bool   `query:"automatic_fields" json:"-"`    // 控制是否返回自动计算的字段, 例如 `created_by`/`created_time`/`last_modified_by`/`last_modified_time`, true 表示返回, 示例值: true
}

// GetBitableRecordResp ...
type GetBitableRecordResp struct {
	Record *GetBitableRecordRespRecord `json:"record,omitempty"` // 记录
}

// GetBitableRecordRespRecord ...
type GetBitableRecordRespRecord struct {
	RecordID         string                                    `json:"record_id,omitempty"`          // 一条记录的唯一标识 id [record_id 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#15d8db94)
	CreatedBy        *GetBitableRecordRespRecordCreatedBy      `json:"created_by,omitempty"`         // 该记录的创建人
	CreatedTime      int64                                     `json:"created_time,omitempty"`       // 该记录的创建时间
	LastModifiedBy   *GetBitableRecordRespRecordLastModifiedBy `json:"last_modified_by,omitempty"`   // 该记录最新一次更新的修改人
	LastModifiedTime int64                                     `json:"last_modified_time,omitempty"` // 该记录最近一次的更新时间
	Fields           map[string]interface{}                    `json:"fields,omitempty"`             // 数据表的字段, 即数据表的列, 当前接口支持的字段类型请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#31f78a3c), 不同类型字段的数据结构请参考[数据结构概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/development-guide/bitable-structure)
}

// GetBitableRecordRespRecordCreatedBy ...
type GetBitableRecordRespRecordCreatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户id, id类型等于user_id_type所指定的类型。
	Name      string `json:"name,omitempty"`       // 用户的中文名称
	EnName    string `json:"en_name,omitempty"`    // 用户的英文名称
	Email     string `json:"email,omitempty"`      // 用户的邮箱
	AvatarURL string `json:"avatar_url,omitempty"` // 头像链接, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
}

// GetBitableRecordRespRecordLastModifiedBy ...
type GetBitableRecordRespRecordLastModifiedBy struct {
	ID        string `json:"id,omitempty"`         // 用户id, id类型等于user_id_type所指定的类型。
	Name      string `json:"name,omitempty"`       // 用户的中文名称
	EnName    string `json:"en_name,omitempty"`    // 用户的英文名称
	Email     string `json:"email,omitempty"`      // 用户的邮箱
	AvatarURL string `json:"avatar_url,omitempty"` // 头像链接, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
}

// getBitableRecordResp ...
type getBitableRecordResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableRecordResp `json:"data,omitempty"`
}
