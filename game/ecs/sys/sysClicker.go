package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemClicker struct {
	player *ec.ComponentPlayer
	ran bool
}

func NewSystemClicker() *SystemClicker {
	sys := SystemClicker{}
	sys.player = nil
	sys.ran = false

	return &sys
}

func (sr *SystemClicker) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float32) {
	if !sr.ran {
		sr.ran = true
		for i := range *entities {
			entity := (*entities)[i]
			if entity.GetTag() == "Player" {
				sr.player = entity.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
				return
			}
		}
		panic("Couldn't find player in Clicker Egg System! Did you add player? Is player entity missing 'Player' tag?")
	}
	//Loop through entities
	for i := range *entities {
		//Check if has clicker
		if (*entities)[i].HasComponent(ec.C_CLICKER) {
			//Check what tag this entity
			if (*entities)[i].GetTag() == "Chicken" {
				//Check if clicker is clicked
				clickC := (*entities)[i].GetComponentWithID(ec.C_CLICKER)
				var click *ec.ComponentClicker = clickC.(*ec.ComponentClicker)
				if click.Clicked {
					//Set clicker to not clicked
					click.Clicked = false
					//Increase egg count on player
					sr.player.Eggs++
				}
			} 
		}
	}
}