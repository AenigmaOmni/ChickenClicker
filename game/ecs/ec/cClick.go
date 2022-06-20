package ec

type ComponentClicker struct {
	ID int
	Clicked bool
	parent Entity
}

func NewComponentClicker() ComponentClicker {
	c := ComponentClicker{}
	c.ID = C_CLICKER
	c.Clicked = false
	return c
}

func (c *ComponentClicker) GetID() int {
	return c.ID
}

func (c *ComponentClicker) GetEntityID() int64 {
	return c.parent.GetID()
}