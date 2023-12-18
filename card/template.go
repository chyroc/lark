package card

import "github.com/chyroc/lark"

// Template ...
func Template(id string, variable map[string]interface{}) *lark.MessageContentCardTemplate {
	return &lark.MessageContentCardTemplate{
		TemplateID:       id,
		TemplateVariable: variable,
	}
}
