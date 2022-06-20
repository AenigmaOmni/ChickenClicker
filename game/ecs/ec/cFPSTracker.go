package ec

type ComponentFPSTracker struct {
	ID int
	parent Entity
}

func NewComponentFPSTracker() ComponentFPSTracker {
	c := ComponentFPSTracker{}
	c.ID = C_FPSTRACKER

	return c
}

func (c *ComponentFPSTracker) GetID() int {
	return c.ID
}

func (c *ComponentFPSTracker) GetEntityID() int64 {
	return c.parent.GetID()
}