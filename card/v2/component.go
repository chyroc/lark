package card

type Component interface {
	Component()
}

type componentBase struct{}

func (componentBase) Component() {}
