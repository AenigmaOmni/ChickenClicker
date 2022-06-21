package sys

import (
	"fmt"

	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemBuyUpgrades struct {
	handBuyCost int
	handBuyMulti float32
	player *ec.ComponentPlayer
	ran bool
}

func (sr *SystemBuyUpgrades) Update(entities *[]ec.Entity, delta float64) {
	if !sr.ran {
		sr.ran = true

		sr.handBuyCost = 50
		sr.handBuyMulti = 1.1
		
		//Loop through entities
		for i := range *entities {
			entity := &(*entities)[i]
			//Set player for reference
			if entity.GetTag() == "Player" {
				sr.player = entity.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
			}

			if entity.GetTag() == "Hand Buy Text" {
				if entity.HasComponent(ec.C_TEXT) {
					hc := entity.GetComponentWithID(ec.C_TEXT)
					var text *ec.ComponentText = hc.(*ec.ComponentText)
					text.Message = fmt.Sprintf("Buy Petter: %v", sr.handBuyCost)

				} else {
					panic("Hand Buy Text not found!")
				}
			}
		}
		if sr.player == nil {
			panic("Couldn't find player in Buy Upgrades System! Did you add player? Is player entity missing 'Player' tag?")
		}
	}

	for i := range *entities {
		e := (*entities)[i]
		if e.GetTag() == "Buy Hand" {
			if e.HasComponent(ec.C_CLICKER) {
				cc := e.GetComponentWithID(ec.C_CLICKER)
				var clicker *ec.ComponentClicker = cc.(*ec.ComponentClicker)
				if clicker.Clicked {
					if sr.player.Eggs >= sr.handBuyCost {
						sr.player.Eggs -= sr.handBuyCost
						temp := float32(sr.handBuyCost) * sr.handBuyMulti
						sr.handBuyCost = int(temp)
					}
				}
			}
		}
	}
}