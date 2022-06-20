package comps

type ComponentFPSTracker struct {
	ID int
}

func NewComponentFPSTracker() ComponentFPSTracker {
	c := ComponentFPSTracker{}
	c.ID = C_FPSTRACKER

	return c
}

func (c *ComponentFPSTracker) GetID() int {
	return c.ID
}