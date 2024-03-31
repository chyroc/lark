package card

import "encoding/json"

func Note(elements ...Component) *ComponentNote {
	return &ComponentNote{
		Elements: elements,
	}
}

// ComponentNote 备注
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/note
//
// 你可以使用备注组件展示卡片内的一些次要信息, 用于辅助说明。备注组件支持添加图标、图片以及文本。
//
//go:generate generate_set_attrs -type=ComponentNote
//go:generate generate_to_map -type=ComponentNote
type ComponentNote struct {
	componentBase

	Elements []Component `json:"elements,omitempty"`
}

// MarshalJSON ...
func (r ComponentNote) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "note"
	return json.Marshal(m)
}

// SetElements set ComponentNote.Elements attribute
func (r *ComponentNote) SetElements(val ...Component) *ComponentNote {
	r.Elements = val
	return r
}

// toMap conv ComponentNote to map
func (r *ComponentNote) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 1)
	if len(r.Elements) != 0 {
		res["elements"] = r.Elements
	}
	return res
}
