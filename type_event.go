package lark

type EventHeaderV2 struct {
	EventID    string    `json:"event_id,omitempty"`    // 事件 ID
	EventType  EventType `json:"event_type,omitempty"`  // 事件类型
	CreateTime string    `json:"create_time,omitempty"` // 事件创建时间戳（单位：毫秒）
	Token      string    `json:"token,omitempty"`       // 事件 Token
	AppID      string    `json:"app_id,omitempty"`      // 应用 ID
	TenantKey  string    `json:"tenant_key,omitempty"`  // 租户 Key
}

type EventHeaderV1 struct {
	UUID      string    `json:"event_id,omitempty"`    // 事件 ID
	EventType EventType `json:"event_type,omitempty"`  // 事件类型
	TS        string    `json:"create_time,omitempty"` // 事件创建时间戳（单位：毫秒）
	Token     string    `json:"token,omitempty"`       // 事件 Token
}
