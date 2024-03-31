package card

import "encoding/json"

func Person(id string) *ComponentPerson {
	return &ComponentPerson{
		UserID: id,
	}
}

// ComponentPerson 人员
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/user-profile
//
// 人员组件支持展示人员的用户名和头像。你可通过传入人员的 open_id、user_id 或 union_id 使用该组件。
//
// 注意事项
// 若你要使用指定应用发送含有人员组件的卡片, 你需保证该应用有访问用户 ID 的权限。否则卡片中的人员组件无法展示人员信息。
//
//go:generate generate_set_attrs -type=ComponentPerson
//go:generate generate_to_map -type=ComponentPerson
type ComponentPerson struct {
	componentBase

	// 人员的头像尺寸。可取值：
	//
	// extra_small：超小尺寸
	// small：小尺寸
	// medium：中尺寸
	// large：大尺寸
	Size AvatarSize `json:"size,omitempty"`

	// 人员的 ID。可选值有：
	//
	// 人员的 Open ID：标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。详情参考如何获取 Open ID
	// 人员的 Union ID：标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。详情参考如何获取 Union ID
	// 人员的 User ID ：标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。详情参考如何获取User ID
	UserID string `json:"user_id,omitempty"`
}

// MarshalJSON ...
func (r ComponentPerson) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "person"
	return json.Marshal(m)
}

// SetSize set ComponentPerson.Size attribute
func (r *ComponentPerson) SetSize(val AvatarSize) *ComponentPerson {
	r.Size = val
	return r
}

// SetSizeExtraSmall set ComponentPerson.Size attribute to AvatarSizeExtraSmall
func (r *ComponentPerson) SetSizeExtraSmall() *ComponentPerson {
	r.Size = AvatarSizeExtraSmall
	return r
}

// SetSizeSmall set ComponentPerson.Size attribute to AvatarSizeSmall
func (r *ComponentPerson) SetSizeSmall() *ComponentPerson {
	r.Size = AvatarSizeSmall
	return r
}

// SetSizeMedium set ComponentPerson.Size attribute to AvatarSizeMedium
func (r *ComponentPerson) SetSizeMedium() *ComponentPerson {
	r.Size = AvatarSizeMedium
	return r
}

// SetSizeLarge set ComponentPerson.Size attribute to AvatarSizeLarge
func (r *ComponentPerson) SetSizeLarge() *ComponentPerson {
	r.Size = AvatarSizeLarge
	return r
}

// SetUserID set ComponentPerson.UserID attribute
func (r *ComponentPerson) SetUserID(val string) *ComponentPerson {
	r.UserID = val
	return r
}

// toMap conv ComponentPerson to map
func (r *ComponentPerson) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 2)
	if r.Size != "" {
		res["size"] = r.Size
	}
	if r.UserID != "" {
		res["user_id"] = r.UserID
	}
	return res
}
