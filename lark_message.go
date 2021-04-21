package lark

func (r *apiImpl) Message() *MessageAPI {
	return &MessageAPI{
		cli: r,
	}
}

type MessageAPI struct {
	cli *apiImpl
}
