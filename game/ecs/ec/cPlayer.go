package ec

type ComponentPlayer struct {
	ID int
	Eggs int
	parent Entity
	HandBuyCost int
	HandBuyMulti float32
	HandProduction int
	HandTimer float32
}

func NewComponentPlayer() ComponentPlayer {
	c := ComponentPlayer{}
	c.ID = C_PLAYER
	c.Eggs = 0
	c.HandBuyCost = 10
	c.HandBuyMulti = 1.1
	c.HandProduction = 1
	c.HandTimer = 3.0
	return c
}

func (c *ComponentPlayer) GetID() int {
	return c.ID
}

func (c *ComponentPlayer) GetEntityID() int64 {
	return c.parent.GetID()
}