package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemPetter struct {
	ran bool
	player *ec.ComponentPlayer
}

func (sr *SystemPetter) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float64) {
	if !sr.ran {
		sr.ran = true
		//Loop through entities
		for i := range *entities {
			entity := (*entities)[i]
			//Set player for reference
			if entity.GetTag() == "Player" {
				sr.player = entity.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
			}

			UpdateBuyHandText(entities)
		}
		if sr.player == nil {
			panic("Couldn't find player in Petter System! Did you add player? Is player entity missing 'Player' tag?")
		}
	}


	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		if e.HasComponent(ec.C_TIMER) && e.GetTag() == "Petter" {
			timer := e.GetComponentWithID(ec.C_TIMER).(*ec.ComponentTimer)
			if timer.Alarm {
				//Set alarm to false and increase eggs and reset timer
				timer.Alarm = false
				sr.player.Eggs += sr.player.HandProduction
				timer.AlarmTime = sr.player.HandTimer
				//Update interface
				UpdateEggs(entities)
			}
		}
	}
}