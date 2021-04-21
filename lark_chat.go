package lark

func (r *apiImpl) Chat() *ChatAPI {
	return &ChatAPI{
		cli: r,
	}
}

type ChatAPI struct {
	cli *apiImpl
}
