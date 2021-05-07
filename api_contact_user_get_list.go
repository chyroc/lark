package lark

import (
	"context"
)

// GetUserList 通过该接口向通讯录获取用户信息列表。
//
// 如果没有全员权限，则返回有通讯录范围的用户列表；如果有全员权限且应用可用范围为全员，将department_id字段设置为0，则可返回根部门下所有用户；如果传入一个具体的department_id,则返回该部门下的所有用户。
// 如果没有单独设置用户的可见性，不填写具体的department_id则不会获取到用户数据
// 如果有全员权限，将department_id字段设置为0，但根部门下没有任何用户的情况下，不会获取到用户数据。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list
func (r *ContactAPI) GetUserList(ctx context.Context, request *GetUserListReq) (*GetUserListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/users",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getUserListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Contact", "GetUserList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetUserListReq struct {
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType *IDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值："open_department_type", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	DepartmentID     *string `query:"department_id" json:"-"`      // 填写该字段表示获取部门下所有用户，选填。, 示例值："od-xxxxxxxxxxxxx"
	PageToken        *string `query:"page_token" json:"-"`         // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："true"
	PageSize         *int    `query:"page_size" json:"-"`          // 分页大小, 示例值：10, 最大值：`50`
}

type getUserListResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetUserListResp `json:"data,omitempty"` //
}

type GetUserListResp struct {
	HasMore   bool                   `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                 `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetUserListRespItem `json:"items,omitempty"`      // -
}

type GetUserListRespItem struct {
	UnionID         string                           `json:"union_id,omitempty"`          // 用户的union_id
	UserID          string                           `json:"user_id,omitempty"`           // 租户内用户的唯一标识
	OpenID          string                           `json:"open_id,omitempty"`           // 用户的open_id
	Name            string                           `json:"name,omitempty"`              // 用户名, 最小长度：`1` 字符
	EnName          string                           `json:"en_name,omitempty"`           // 英文名
	Email           string                           `json:"email,omitempty"`             // 邮箱, 字段权限要求:  获取用户邮箱
	Mobile          string                           `json:"mobile,omitempty"`            // 手机号, 字段权限要求:  获取用户手机号
	MobileVisible   bool                             `json:"mobile_visible,omitempty"`    // 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码
	Gender          int                              `json:"gender,omitempty"`            // 性别, 可选值有: `0`：保密, `1`：男, `2`：女
	Avatar          *GetUserListRespItemAvatar       `json:"avatar,omitempty"`            // 用户头像信息
	Status          *GetUserListRespItemStatus       `json:"status,omitempty"`            // 用户状态
	DepartmentIDs   []string                         `json:"department_ids,omitempty"`    // 用户所属部门的ID列表
	LeaderUserID    string                           `json:"leader_user_id,omitempty"`    // 用户的直接主管的用户ID
	City            string                           `json:"city,omitempty"`              // 城市
	Country         string                           `json:"country,omitempty"`           // 国家
	WorkStation     string                           `json:"work_station,omitempty"`      // 工位
	JoinTime        int                              `json:"join_time,omitempty"`         // 入职时间
	IsTenantManager bool                             `json:"is_tenant_manager,omitempty"` // 是否是租户管理员
	EmployeeNo      string                           `json:"employee_no,omitempty"`       // 工号
	EmployeeType    int                              `json:"employee_type,omitempty"`     // 员工类型, 可选值有: `1`：正式员工, `2`：实习生, `3`：外包, `4`：劳务, `5`：顾问
	Orders          []*GetUserListRespItemOrder      `json:"orders,omitempty"`            // 用户排序信息
	CustomAttrs     []*GetUserListRespItemCustomAttr `json:"custom_attrs,omitempty"`      // 自定义属性
	EnterpriseEmail string                           `json:"enterprise_email,omitempty"`  // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务
}

type GetUserListRespItemAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type GetUserListRespItemStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否冻结
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type GetUserListRespItemOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int    `json:"user_order,omitempty"`       // 用户在部门内的排序，数值越大，排序越靠前
	DepartmentOrder int    `json:"department_order,omitempty"` // 用户的部门间的排序，数值越大，排序越靠前
}

type GetUserListRespItemCustomAttr struct {
	Type  string                              `json:"type,omitempty"`  // 自定义属性类型
	ID    string                              `json:"id,omitempty"`    // 自定义属性ID
	Value *GetUserListRespItemCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type GetUserListRespItemCustomAttrValue struct {
	Text  string `json:"text,omitempty"`   // 属性文本
	Url   string `json:"url,omitempty"`    // URL
	PcUrl string `json:"pc_url,omitempty"` // PC上的URL
}
