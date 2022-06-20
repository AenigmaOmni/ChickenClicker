package ec

type ComponentPosition struct {
	ID int
	X float64
	Y float64
	parent Entity
}

func NewComponentPosition(tx float64, ty float64) ComponentPosition {
	c := ComponentPosition{}
	c.ID = C_POSITION
	c.X = tx
	c.Y = ty
	return c
}

func (c *ComponentPosition) GetID() int {
	return c.ID
}

func (c *ComponentPosition) GetEntityID() int64 {
	return c.parent.GetID()
}