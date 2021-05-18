// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateField 该接口用于在数据表中新增一个字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/create
func (r *BitableService) CreateField(ctx context.Context, request *CreateFieldReq, options ...MethodOptionFunc) (*CreateFieldResp, *Response, error) {
	if r.cli.mock.mockBitableCreateField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateField mock enable")
		return r.cli.mock.mockBitableCreateField(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Bitable#CreateField call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateField request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(createFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#CreateField POST https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#CreateField POST https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "CreateField", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateField success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableCreateField(f func(ctx context.Context, request *CreateFieldReq, options ...MethodOptionFunc) (*CreateFieldResp, *Response, error)) {
	r.mockBitableCreateField = f
}

func (r *Mock) UnMockBitableCreateField() {
	r.mockBitableCreateField = nil
}

type CreateFieldReq struct {
	UserIDType *IDType      `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	AppToken   string       `path:"app_token" json:"-"`     // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID    string       `path:"table_id" json:"-"`      // table id, 示例值："tblsRc9GRRXKqhvW"
	FieldName  string       `json:"field_name,omitempty"`   // 多维表格字段名, 示例值："多行文本"
	Type       int          `json:"type,omitempty"`         // 多维表格字段类型, 示例值：1
	Property   *interface{} `json:"property,omitempty"`     // 字段属性, 示例值：[,                    {,                        "name": "选项A",                    },,                    {,                        "name": "选项B",                    },                ],
}

type createFieldResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *CreateFieldResp `json:"data,omitempty"` //
}

type CreateFieldResp struct {
	Field *CreateFieldRespField `json:"field,omitempty"` // 字段
}

type CreateFieldRespField struct {
	FieldID   string      `json:"field_id,omitempty"`   // 多维表格字段 id
	FieldName string      `json:"field_name,omitempty"` // 多维表格字段名
	Type      int         `json:"type,omitempty"`       // 多维表格字段类型
	Property  interface{} `json:"property,omitempty"`   // 字段属性
}
