package lark

func (r *apiImpl) Chat() *ChatAPI {
	return &ChatAPI{
		cli: r,
	}
}

type ChatAPI struct {
	cli *apiImpl
}

func (r *apiImpl) Message() *MessageAPI {
	return &MessageAPI{
		cli: r,
	}
}

type MessageAPI struct {
	cli *apiImpl
}

func (r *apiImpl) Token() *TokenAPI {
	return &TokenAPI{
		cli: r,
	}
}

type TokenAPI struct {
	cli *apiImpl
}
