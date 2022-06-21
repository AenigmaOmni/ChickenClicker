package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemTimer struct {

}

func (sr *SystemTimer) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float64) {
	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		if e.HasComponent(ec.C_TIMER) {
			timer := e.GetComponentWithID(ec.C_TIMER).(*ec.ComponentTimer)
			if !timer.Alarm {
				timer.Timer += delta
				if timer.AlarmTime <= timer.Timer {
					timer.Alarm = true
					timer.Timer = 0.0
				}
			}
		}
	}
}