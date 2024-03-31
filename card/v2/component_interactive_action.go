package card

import "encoding/json"

func Action(actions ...Component) *ComponentAction {
	return &ComponentAction{
		Layout:  ActionLayoutDefault,
		Actions: actions,
	}
}

//go:generate generate_set_attrs -type=ComponentAction
//go:generate generate_to_map -type=ComponentAction
type ComponentAction struct {
	componentBase

	MultiURL *ObjectMultiURL `json:"multi_url,omitempty"`
	Layout   ActionLayout    `json:"layout,omitempty"`
	Actions  []Component     `json:"actions,omitempty"`
}

type ActionLayout = string

const (
	ActionLayoutDefault    ActionLayout = "default"    // 通栏堆叠
	ActionLayoutBisected   ActionLayout = "bisected"   // 二等分
	ActionLayoutTrisection ActionLayout = "trisection" // 三等分
	ActionLayoutFlow       ActionLayout = "flow"       // 流式
)

func (r ComponentAction) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "action"
	return json.Marshal(m)
}

// SetMultiURL set ComponentAction.MultiURL attribute
func (r *ComponentAction) SetMultiURL(val *ObjectMultiURL) *ComponentAction {
	r.MultiURL = val
	return r
}

// SetLayout set ComponentAction.Layout attribute
func (r *ComponentAction) SetLayout(val ActionLayout) *ComponentAction {
	r.Layout = val
	return r
}

// SetLayoutDefault set ComponentAction.Layout attribute to ActionLayoutDefault
func (r *ComponentAction) SetLayoutDefault() *ComponentAction {
	r.Layout = ActionLayoutDefault
	return r
}

// SetLayoutBisected set ComponentAction.Layout attribute to ActionLayoutBisected
func (r *ComponentAction) SetLayoutBisected() *ComponentAction {
	r.Layout = ActionLayoutBisected
	return r
}

// SetLayoutTrisection set ComponentAction.Layout attribute to ActionLayoutTrisection
func (r *ComponentAction) SetLayoutTrisection() *ComponentAction {
	r.Layout = ActionLayoutTrisection
	return r
}

// SetLayoutFlow set ComponentAction.Layout attribute to ActionLayoutFlow
func (r *ComponentAction) SetLayoutFlow() *ComponentAction {
	r.Layout = ActionLayoutFlow
	return r
}

// SetActions set ComponentAction.Actions attribute
func (r *ComponentAction) SetActions(val ...Component) *ComponentAction {
	r.Actions = val
	return r
}

// toMap conv ComponentAction to map
func (r *ComponentAction) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 3)
	if r.MultiURL != nil {
		res["multi_url"] = r.MultiURL
	}
	if r.Layout != "" {
		res["layout"] = r.Layout
	}
	if len(r.Actions) != 0 {
		res["actions"] = r.Actions
	}
	return res
}
