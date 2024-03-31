package card

func Confirm(title, text string) *ObjectConfirm {
	return &ObjectConfirm{
		Title: Text(title),
		Text:  Text(text),
	}
}

//go:generate generate_set_attrs -type=ObjectConfirm
//go:generate generate_to_map -type=ObjectConfirm
type ObjectConfirm struct {
	Title *ObjectText `json:"title,omitempty"`
	Text  *ObjectText `json:"text,omitempty"`
}

// SetTitle set ObjectConfirm.Title attribute
func (r *ObjectConfirm) SetTitle(val *ObjectText) *ObjectConfirm {
	r.Title = val
	return r
}

// SetText set ObjectConfirm.Text attribute
func (r *ObjectConfirm) SetText(val *ObjectText) *ObjectConfirm {
	r.Text = val
	return r
}

// toMap conv ObjectConfirm to map
func (r *ObjectConfirm) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 2)
	if r.Title != nil {
		res["title"] = r.Title
	}
	if r.Text != nil {
		res["text"] = r.Text
	}
	return res
}
