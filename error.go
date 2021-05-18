package lark

import (
	"fmt"
)

type Error struct {
	Scope    string
	FuncName string
	Code     int64
	Msg      string
}

func (r *Error) Error() string {
	if r.Code == 0 {
		return ""
	}
	return fmt.Sprintf("request %s#%s failed: code: %d, msg: %s", r.Scope, r.FuncName, r.Code, r.Msg)
}

func NewError(scope, funcName string, code int64, msg string) error {
	return &Error{
		Scope:    scope,
		FuncName: funcName,
		Code:     code,
		Msg:      msg,
	}
}
