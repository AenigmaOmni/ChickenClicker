package comps

type ComponentPosition struct {
	ID int
	X float64
	Y float64
}

func NewComponentPosition() ComponentPosition {
	c := ComponentPosition{}
	c.ID = C_POSITION

	return c
}

func (c *ComponentPosition) GetID() int {
	return c.ID
}