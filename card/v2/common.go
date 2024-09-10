package card

import (
	"encoding/json"
)

type anyMap = map[string]any

func unmarshalComponent(data json.RawMessage) (Component, error) {
	tag := new(unmarshalComponentTag)
	err := json.Unmarshal(data, tag)
	if err != nil {
		return nil, err
	}
	switch tag.Tag {
	case "img":
		res := &ComponentImage{}
		err = json.Unmarshal(data, res)
		return res, err
	case "hr":
		return &ComponentDivider{}, nil
	case "column_set":
		res := &ComponentColumnSet{}
		err = json.Unmarshal(data, res)
		return res, err
	case "markdown":
		res := &ComponentMarkdown{}
		err = json.Unmarshal(data, res)
		return res, err
	case "input":
		res := &ComponentInput{}
		err = json.Unmarshal(data, res)
		return res, err
	case "button":
		res := &ComponentButton{}
		err = json.Unmarshal(data, res)
		return res, err
	}

	return nil, nil
}
