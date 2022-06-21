package sys

import (
	"fmt"

	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemClickerEgg struct {
	player *ec.ComponentPlayer
	ran bool
}

func NewSystemClickerEgg() *SystemClickerEgg {
	sys := SystemClickerEgg{}
	sys.player = nil
	sys.ran = false

	return &sys
}

func (sr *SystemClickerEgg) Update(entities *[]ec.Entity, delta float64) {
	if !sr.ran {
		sr.ran = true
		for i := range *entities {
			entity := &(*entities)[i]
			if entity.GetTag() == "Player" {
				sr.player = entity.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
				return
			}
		}
		panic("Couldn't find player in Clicker Egg System! Did you add player? Is player entity missing 'Player' tag?")
	}
	//Loop through entities
	for i := range *entities {
		//Check if this entity is the chicken with the tag
		if (*entities)[i].GetTag() == "Chicken" {
			//Check if chicken has clicker
			if (*entities)[i].HasComponent(ec.C_CLICKER) {
				//Check if clicker is clicked
				clickC := (*entities)[i].GetComponentWithID(ec.C_CLICKER)
				var click *ec.ComponentClicker = clickC.(*ec.ComponentClicker)
				if click.Clicked {
					//Set clicker to not clicked
					click.Clicked = false
					//Increase egg count on player
					sr.player.Eggs++

					//Loop through entities again to find egg counter
					for k := range *entities {
						//Check if this is the egg counter with tag
						if (*entities)[k].GetTag() == "Egg Counter" {
							//Get the text component
							tC := (*entities)[k].GetComponentWithID(ec.C_TEXT)
							var text *ec.ComponentText = tC.(*ec.ComponentText)
							str := fmt.Sprintf("Eggs: %v", sr.player.Eggs)
							text.Message = str
						}
					}
				}
			}
		}
	}
}