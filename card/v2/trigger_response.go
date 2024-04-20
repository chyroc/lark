package card

import "encoding/json"

type TriggerResponse struct {
	Toast    *TriggerToast
	Template *TriggerTemplate
	Card     *ComponentCard
}

type TriggerToast struct {
	// 提示类型。可选值：
	//
	// info
	// success
	// error
	// warning
	Type string `json:"type"`
	// 提示文案
	Content string `json:"content"`
	// 多语言提示文案
	I18n map[string]string `json:"i18n"`
}

type TriggerTemplate struct {
	// 卡片模板 ID，即卡片 ID
	ID string `json:"template_id"`
	// 卡片模板的变量
	Variable map[string]interface{} `json:"template_variable"`
	// 卡片模版的版本
	VersionName string `json:"template_version_name"`
}

func (r TriggerResponse) String() string {
	bs, _ := r.MarshalJSON()
	return string(bs)
}

func (r TriggerResponse) MarshalJSON() ([]byte, error) {
	card := map[string]interface{}{}
	if r.Template != nil {
		card["type"] = "template"
		card["data"] = r.Template
	} else if r.Card != nil {
		card["type"] = "raw"
		card["data"] = r.Card
	}
	return json.Marshal(map[string]interface{}{
		"toast": r.Toast,
		"card":  card,
	})
}
