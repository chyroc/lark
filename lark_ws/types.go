package lark_ws

import "strconv"

type messageType string

const (
	messageTypeEvent messageType = "event"
	messageTypeCard  messageType = "card"
	messageTypePing  messageType = "ping"
	messageTypePong  messageType = "pong"
)

type frameType int

const (
	frameTypeControl frameType = 0
	frameTypeData    frameType = 1
)

type wsHeader []Header

func (h wsHeader) GetString(key string) string {
	for _, header := range h {
		if header.Key == key {
			return header.Value
		}
	}

	return ""
}

func (h wsHeader) GetInt(key string) int {
	for _, header := range h {
		if header.Key == key {
			if val, err := strconv.Atoi(header.Value); err == nil {
				return val
			}
		}
	}

	return 0
}

func (h *wsHeader) Add(key, value string) {
	header := Header{
		Key:   key,
		Value: value,
	}
	*h = append(*h, header)
}

// responseMessage 上行响应消息结构, 置于 Frame.Payload
type responseMessage struct {
	StatusCode int               `json:"code"`
	Headers    map[string]string `json:"headers"`
	Data       []byte            `json:"data"`
}

func newPingFrame(serviceID int32) *Frame {
	return &Frame{
		Method:  int32(frameTypeControl),
		Service: serviceID,
		Headers: []Header{
			{
				Key:   "type",
				Value: string(messageTypePing),
			},
		},
	}
}
