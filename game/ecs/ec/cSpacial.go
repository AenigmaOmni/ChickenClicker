package ec

type ComponentSpatial struct {
	ID int
	X float64
	Y float64
	Width float64
	Height float64
	parent Entity
}

func NewComponentSpatial(tx float64, ty float64, w float64, h float64) ComponentSpatial {
	c := ComponentSpatial{}
	c.ID = C_SPATIAL
	c.X = tx
	c.Y = ty
	c.Width = w
	c.Height = h
	return c
}

func (c *ComponentSpatial) GetID() int {
	return c.ID
}

func (c *ComponentSpatial) GetEntityID() int64 {
	return c.parent.GetID()
}