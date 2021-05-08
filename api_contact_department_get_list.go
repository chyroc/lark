package lark

import (
	"context"
)

// GetDepartmentList 该接口用于获取当前部门子部门列表。
//
// - 使用tenant_access_token时,只返回权限范围内的部门。
// - 使用user_access_token时，返回可见性范围内的所有可见部门。当进行递归查询时，只筛查最多1000个部门的可见性。
// - fetch_child字段填写为false：如果填写具体的部门ID，则返回该部门下的一级子部门；如果没有填写部门ID, 若有全员权限，返回根部门信息，若没有全员权限则返回通讯录范围中配置的部门及其一级子部门。
// - fetch_child字段填写为true：如果填写具体的部门ID，则返回该部门下所有子部门；如果没有填写部门ID, 若有全员权限，返回根部门信息，可以根据根部门ID获取其下的一级子部门，若没有全员权限则返回通讯录范围中配置的部门及其子部门。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/list
func (r *ContactAPI) GetDepartmentList(ctx context.Context, request *GetDepartmentListReq) (*GetDepartmentListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/departments",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getDepartmentListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Contact", "GetDepartmentList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetDepartmentListReq struct {
	UserIDType         *IDType `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`，字段权限要求: 获取用户 userid
	DepartmentIDType   *IDType `query:"department_id_type" json:"-"`   // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	ParentDepartmentID *string `query:"parent_department_id" json:"-"` // 父部门的ID，填上获取部门下所有子部门, 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	FetchChild         *bool   `query:"fetch_child" json:"-"`          // 是否递归获取子部门, 示例值: 是否递归获取子部门，默认值：false
	PageToken          *string `query:"page_token" json:"-"`           // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值: "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ%2BG8JG6tK7%2BZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR"
	PageSize           *int    `query:"page_size" json:"-"`            // 分页大小, 示例值: 10, 最大值：`50`
}

type getDepartmentListResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetDepartmentListResp `json:"data,omitempty"` //
}

type GetDepartmentListResp struct {
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                       `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetDepartmentListRespItem `json:"items,omitempty"`
}

type GetDepartmentListRespItem struct {
	Name               string                             `json:"name,omitempty"`                 // 部门名称, 最小长度：`1` 字符
	I18nName           *GetDepartmentListRespItemI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称
	ParentDepartmentID string                             `json:"parent_department_id,omitempty"` // 父部门的ID,* 创建根部门，该参数值为 “0”
	DepartmentID       string                             `json:"department_id,omitempty"`        // 本部门的自定义部门ID, 最大长度：`128` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`
	OpenDepartmentID   string                             `json:"open_department_id,omitempty"`   // 部门的open_id
	LeaderUserID       string                             `json:"leader_user_id,omitempty"`       // 部门主管用户ID
	ChatID             string                             `json:"chat_id,omitempty"`              // 部门群ID
	Order              string                             `json:"order,omitempty"`                // 部门的排序，即部门在其同级部门的展示顺序
	UnitIDs            []string                           `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表，当前只支持一个
	MemberCount        int                                `json:"member_count,omitempty"`         // 部门下用户的个数
	Status             *GetDepartmentListRespItemStatus   `json:"status,omitempty"`               // 部门状态
}

type GetDepartmentListRespItemI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

type GetDepartmentListRespItemStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}
