package lark

type IDType string

const (
	IDTypeUserID  IDType = "user_id"  // 以 user_id 来识别成员
	IDTypeUnionID IDType = "union_id" // 以 union_id 来识别成员
	IDTypeOpenID  IDType = "open_id"  // 以 open_id 来识别成员
	IDTypeAppID   IDType = "app_id"   // 以 app_id 来识别成员
)
