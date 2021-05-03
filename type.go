package lark

type IDType string

const (
	IDTypeUserID  IDType = "user_id"  // 以 user_id 来识别成员
	IDTypeUnionID IDType = "union_id" // 以 union_id 来识别成员
	IDTypeOpenID  IDType = "open_id"  // 以 open_id 来识别成员
	IDTypeAppID   IDType = "app_id"   // 以 app_id 来识别成员
)

func IDTypePtr(idType IDType) *IDType {
	return &idType
}

type I18nNames struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名,**示例值**："群聊"
	EnUs string `json:"en_us,omitempty"` // 英文名,**示例值**："group chat"
	JaJp string `json:"ja_jp,omitempty"` // 日文名,**示例值**："グループチャット"
}

type Sender struct {
	Id         string `json:"id,omitempty"`          // 该字段标识发送者的id
	IDType     IDType `json:"id_type,omitempty"`     // 该字段标识发送者的id类型
	SenderType string `json:"sender_type,omitempty"` // 该字段标识发送者的类型
}

type MessageBody struct {
	Content string `json:"content,omitempty"` // 消息jsonContent
}

type Mention struct {
	Key    string `json:"key,omitempty"`     // mention key
	Id     string `json:"id,omitempty"`      // 用户open id
	IDType IDType `json:"id_type,omitempty"` // id 可以是open_id，user_id或者union_id
	Name   string `json:"name,omitempty"`    // 被at用户的姓名
}

type HelpDeskCustomizedField struct {
	ID      string `json:"id"`       // id ,示例值："123"
	Value   string `json:"value"`    // value ,示例值："value"
	KeyName string `json:"key_name"` // key name ,示例值："key"
}
