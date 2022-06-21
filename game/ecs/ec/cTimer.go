package ec

type ComponentTimer struct {
	ID int
	Timer float32
	AlarmTime float32
	Alarm bool
	parent Entity
}

func NewComponentTimer(time float32) ComponentTimer {
	c := ComponentTimer{}
	c.ID = C_TIMER
	c.Alarm = false
	c.Timer = 0.0
	c.AlarmTime = time
	return c
}

func (c *ComponentTimer) GetID() int {
	return c.ID
}

func (c *ComponentTimer) GetEntityID() int64 {
	return c.parent.GetID()
}