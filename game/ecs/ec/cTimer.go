package ec

type ComponentTimer struct {
	ID int
	Timer float64
	AlarmTime float64
	Alarm bool
	parent Entity
}

func NewComponentTimer(time float64) ComponentTimer {
	c := ComponentTimer{}
	c.ID = C_TIMER
	c.Alarm = false
	c.Timer = 0
	c.AlarmTime = time
	return c
}

func (c *ComponentTimer) GetID() int {
	return c.ID
}

func (c *ComponentTimer) GetEntityID() int64 {
	return c.parent.GetID()
}