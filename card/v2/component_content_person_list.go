package card

import "encoding/json"

// PersonList 人员列表
func PersonList(ids ...string) *ComponentPersonList {
	persons := make([]*ComponentPerson, 0, len(ids))
	for _, id := range ids {
		persons = append(persons, &ComponentPerson{UserID: id})
	}
	return &ComponentPersonList{
		ShowName:   true,
		ShowAvatar: true,
		Persons:    persons,
	}
}

// ComponentPersonList 人员列表
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/user-list
//
// 人员列表组件支持展示多个人员的用户名和头像。你可通过传入人员的 open_id、user_id 或 union_id 使用该组件。
//
// 注意事项
// 若你要使用指定应用发送含有人员列表组件的卡片, 你需保证该应用有访问用户 ID 的权限。否则卡片中的人员列表组件无法展示人员信息。
//
//go:generate generate_set_attrs -type=ComponentPersonList
//go:generate generate_to_map -type=ComponentPersonList
type ComponentPersonList struct {
	componentBase

	// 最大显示行数, 默认不限制最大显示行数。
	Lines int64 `json:"lines,omitempty"`

	// 是否展示人员的用户名。
	ShowName bool `json:"show_name,omitempty"`

	// 是否展示人员的头像。
	ShowAvatar bool `json:"show_avatar,omitempty"`

	// 人员的头像尺寸。可取值：
	//
	// extra_small：超小尺寸
	// small：小尺寸
	// medium：中尺寸
	// large：大尺寸
	Size AvatarSize `json:"size,omitempty"`

	// 人员列表。
	Persons []*ComponentPerson `json:"-,omitempty"`

	// 添加图标作为文本前缀图标。支持自定义或使用图标库中的图标。
	Icon *ObjectIcon `json:"icon,omitempty"`
}

type AvatarSize = string

const (
	AvatarSizeExtraSmall AvatarSize = "extra_small" // 超小尺寸
	AvatarSizeSmall      AvatarSize = "small"       // 小尺寸
	AvatarSizeMedium     AvatarSize = "medium"      // 中尺寸
	AvatarSizeLarge      AvatarSize = "large"       // 大尺寸
)

// MarshalJSON ...
func (r ComponentPersonList) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "person_list"
	if len(r.Persons) > 0 {
		persons := make([]componentPersonListPerson, 0, len(r.Persons))
		for _, p := range r.Persons {
			persons = append(persons, componentPersonListPerson{ID: p.UserID})
		}
		m["persons"] = persons
	}
	return json.Marshal(m)
}

type componentPersonListPerson struct {
	ID string `json:"id,omitempty"`
}

// SetLines set ComponentPersonList.Lines attribute
func (r *ComponentPersonList) SetLines(val int64) *ComponentPersonList {
	r.Lines = val
	return r
}

// SetShowName set ComponentPersonList.ShowName attribute
func (r *ComponentPersonList) SetShowName(val bool) *ComponentPersonList {
	r.ShowName = val
	return r
}

// SetShowAvatar set ComponentPersonList.ShowAvatar attribute
func (r *ComponentPersonList) SetShowAvatar(val bool) *ComponentPersonList {
	r.ShowAvatar = val
	return r
}

// SetSize set ComponentPersonList.Size attribute
func (r *ComponentPersonList) SetSize(val AvatarSize) *ComponentPersonList {
	r.Size = val
	return r
}

// SetSizeExtraSmall set ComponentPersonList.Size attribute to AvatarSizeExtraSmall
func (r *ComponentPersonList) SetSizeExtraSmall() *ComponentPersonList {
	r.Size = AvatarSizeExtraSmall
	return r
}

// SetSizeSmall set ComponentPersonList.Size attribute to AvatarSizeSmall
func (r *ComponentPersonList) SetSizeSmall() *ComponentPersonList {
	r.Size = AvatarSizeSmall
	return r
}

// SetSizeMedium set ComponentPersonList.Size attribute to AvatarSizeMedium
func (r *ComponentPersonList) SetSizeMedium() *ComponentPersonList {
	r.Size = AvatarSizeMedium
	return r
}

// SetSizeLarge set ComponentPersonList.Size attribute to AvatarSizeLarge
func (r *ComponentPersonList) SetSizeLarge() *ComponentPersonList {
	r.Size = AvatarSizeLarge
	return r
}

// SetPersons set ComponentPersonList.Persons attribute
func (r *ComponentPersonList) SetPersons(val ...*ComponentPerson) *ComponentPersonList {
	r.Persons = val
	return r
}

// SetIcon set ComponentPersonList.Icon attribute
func (r *ComponentPersonList) SetIcon(val *ObjectIcon) *ComponentPersonList {
	r.Icon = val
	return r
}

// toMap conv ComponentPersonList to map
func (r *ComponentPersonList) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 5)
	if r.Lines != 0 {
		res["lines"] = r.Lines
	}
	if r.ShowName != false {
		res["show_name"] = r.ShowName
	}
	if r.ShowAvatar != false {
		res["show_avatar"] = r.ShowAvatar
	}
	if r.Size != "" {
		res["size"] = r.Size
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	return res
}
