package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
)

type SystemTimer struct {

}

func (sr *SystemTimer) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float32) {
	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		if e.HasComponent(ec.C_TIMER) {
			timer := e.GetComponentWithID(ec.C_TIMER).(*ec.ComponentTimer)
			if !timer.Alarm {
				d := delta
				timer.Timer += d
				if timer.Timer >= timer.AlarmTime {
					timer.Alarm = true
					timer.Timer = 0.0
				}
			}
		}
	}
}