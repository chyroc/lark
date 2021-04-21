package lark

func New() API {
	return &apiImpl{}
}

type API interface {
	Message() *MessageAPI
}
