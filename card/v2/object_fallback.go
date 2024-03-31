package card

import "encoding/json"

func Fallback(text string) *ObjectFallback {
	return &ObjectFallback{
		Text: Text(text),
	}
}

func DropFallback() *ObjectFallback {
	return &ObjectFallback{
		drop: true,
	}
}

//go:generate generate_set_attrs -type=ObjectFallback
//go:generate generate_to_map -type=ObjectFallback
type ObjectFallback struct {
	Text *ObjectText `json:"text,omitempty"`
	drop bool
}

// MarshalJSON ...
func (r ObjectFallback) MarshalJSON() ([]byte, error) {
	if r.drop {
		return []byte(`"drop"`), nil
	}
	m := r.toMap()
	m["tag"] = "fallback_text"
	return json.Marshal(m)
}

// SetText set ObjectFallback.Text attribute
func (r *ObjectFallback) SetText(val *ObjectText) *ObjectFallback {
	r.Text = val
	return r
}

// Setdrop set ObjectFallback.drop attribute
func (r *ObjectFallback) Setdrop(val bool) *ObjectFallback {
	r.drop = val
	return r
}

// toMap conv ObjectFallback to map
func (r *ObjectFallback) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 1)
	if r.Text != nil {
		res["text"] = r.Text
	}
	return res
}
