package ec

type ComponentPlayer struct {
	ID int
	Eggs int
	parent Entity
}

func NewComponentPlayer() ComponentPlayer {
	c := ComponentPlayer{}
	c.ID = C_PLAYER
	c.Eggs = 0
	return c
}

func (c *ComponentPlayer) GetID() int {
	return c.ID
}

func (c *ComponentPlayer) GetEntityID() int64 {
	return c.parent.GetID()
}