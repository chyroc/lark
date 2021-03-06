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

// GetBitableFieldList 根据 app_token 和 table_id, 获取数据表的所有字段
//
// 该接口支持调用频率上限为 20 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list
func (r *BitableService) GetBitableFieldList(ctx context.Context, request *GetBitableFieldListReq, options ...MethodOptionFunc) (*GetBitableFieldListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableFieldList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableFieldList mock enable")
		return r.cli.mock.mockBitableGetBitableFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableFieldList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableFieldList mock BitableGetBitableFieldList method
func (r *Mock) MockBitableGetBitableFieldList(f func(ctx context.Context, request *GetBitableFieldListReq, options ...MethodOptionFunc) (*GetBitableFieldListResp, *Response, error)) {
	r.mockBitableGetBitableFieldList = f
}

// UnMockBitableGetBitableFieldList un-mock BitableGetBitableFieldList method
func (r *Mock) UnMockBitableGetBitableFieldList() {
	r.mockBitableGetBitableFieldList = nil
}

// GetBitableFieldListReq ...
type GetBitableFieldListReq struct {
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID   string  `path:"table_id" json:"-"`    // table id, 示例值: "tblsRc9GRRXKqhvW"
	ViewID    *string `query:"view_id" json:"-"`    // 视图 ID, 示例值: "vewOVMEXPF"
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "fldwJ4YrtB"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `100`
}

// GetBitableFieldListResp ...
type GetBitableFieldListResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Total     int64                          `json:"total,omitempty"`      // 总数
	Items     []*GetBitableFieldListRespItem `json:"items,omitempty"`      // 字段信息
}

// GetBitableFieldListRespItem ...
type GetBitableFieldListRespItem struct {
	FieldID   string                               `json:"field_id,omitempty"`   // 多维表格字段 id
	FieldName string                               `json:"field_name,omitempty"` // 多维表格字段名
	Type      int64                                `json:"type,omitempty"`       // 多维表格字段类型, 可选值有: `1`: 多行文本, `2`: 数字, `3`: 单选, `4`: 多选, `5`: 日期, `7`: 复选框, `11`: 人员, `15`: 超链接, `17`: 附件, `18`: 关联, `20`: 公式, `21`: 双向关联, `1001`: 创建时间, `1002`: 最后更新时间, `1003`: 创建人, `1004`: 修改人
	Property  *GetBitableFieldListRespItemProperty `json:"property,omitempty"`   // 字段属性, 具体参考: [字段编辑指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide)
}

// GetBitableFieldListRespItemProperty ...
type GetBitableFieldListRespItemProperty struct {
	Options       []*GetBitableFieldListRespItemPropertyOption `json:"options,omitempty"`         // 单选、多选字段的选项信息
	Formatter     string                                       `json:"formatter,omitempty"`       // 数字、公式字段的显示格式
	DateFormatter string                                       `json:"date_formatter,omitempty"`  // 日期、创建时间、最后更新时间字段的显示格式
	AutoFill      bool                                         `json:"auto_fill,omitempty"`       // 日期字段中新纪录自动填写创建时间
	Multiple      bool                                         `json:"multiple,omitempty"`        // 人员字段中允许添加多个成员, 单向关联、双向关联中允许添加多个记录
	TableID       string                                       `json:"table_id,omitempty"`        // 单向关联、双向关联字段中关联的数据表的id
	TableName     string                                       `json:"table_name,omitempty"`      // 单向关联、双向关联字段中关联的数据表的名字
	BackFieldName string                                       `json:"back_field_name,omitempty"` // 双向关联字段中关联的数据表中对应的双向关联字段的名字
}

// GetBitableFieldListRespItemPropertyOption ...
type GetBitableFieldListRespItemPropertyOption struct {
	Name  string `json:"name,omitempty"`  // 选项名
	ID    string `json:"id,omitempty"`    // 选项 ID, 创建时不允许指定 ID
	Color int64  `json:"color,omitempty"` // 选项颜色
}

// getBitableFieldListResp ...
type getBitableFieldListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableFieldListResp `json:"data,omitempty"`
}
