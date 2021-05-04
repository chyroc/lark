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

func UnwrapMessageContent(resp *GetMessageRespItem) (*MessageContent, error) {
	switch resp.MsgType {
	case "text", "image":
		content := new(MessageContent)
		err := json.Unmarshal([]byte(resp.Body.Content), content)
		if err != nil {
			return nil, err
		}
		return content, nil
	}
	return nil, fmt.Errorf("unknown message type: %s", resp.MsgType)
}
