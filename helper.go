package lark

import (
	"encoding/json"
	"fmt"
)

func GetErrorCode(err error) int {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.Code
		}
	}
	return -1
}

func UnwrapMessageContent(msgType MsgType, content string) (*MessageContent, error) {
	switch msgType {
	case "text", "image":
		msg := new(MessageContent)
		err := json.Unmarshal([]byte(content), msg)
		if err != nil {
			return nil, fmt.Errorf("invalid content: %s, type: %v", content, msgType)
		}
		return msg, nil
	}
	return nil, fmt.Errorf("unknown message type: %s", msgType)
}
