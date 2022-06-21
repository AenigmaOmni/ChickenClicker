package ec

type ComponentPlayer struct {
	ID int
	Eggs int
	parent Entity
	HandBuyCost int
	HandBuyMulti float32
}

func NewComponentPlayer() ComponentPlayer {
	c := ComponentPlayer{}
	c.ID = C_PLAYER
	c.Eggs = 0
	c.HandBuyCost = 50
	c.HandBuyMulti = 1.1
	return c
}

func (c *ComponentPlayer) GetID() int {
	return c.ID
}

func (c *ComponentPlayer) GetEntityID() int64 {
	return c.parent.GetID()
}