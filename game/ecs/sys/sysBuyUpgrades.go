package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemBuyUpgrades struct {
	player *ec.ComponentPlayer
	ran bool
}

func (sr *SystemBuyUpgrades) Update(entities *[]ec.Entity, delta float64) {
	if !sr.ran {
		sr.ran = true
		//Loop through entities
		for i := range *entities {
			entity := &(*entities)[i]
			//Set player for reference
			if entity.GetTag() == "Player" {
				sr.player = entity.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
			}

			UpdateBuyHandText(entities)
		}
		if sr.player == nil {
			panic("Couldn't find player in Buy Upgrades System! Did you add player? Is player entity missing 'Player' tag?")
		}
	}

	//Check the clickers
	for i := range *entities {
		e := (*entities)[i]
		//If trying to buy a petter
		if e.GetTag() == "Buy Hand" {
			//Make sure it has a clicker
			if e.HasComponent(ec.C_CLICKER) {
				cc := e.GetComponentWithID(ec.C_CLICKER)
				var clicker *ec.ComponentClicker = cc.(*ec.ComponentClicker)
				//If it's clicked?
				if clicker.Clicked {
					//If we have enough eggs, buy a petter and increase price of petter
					if sr.player.Eggs >= sr.player.HandBuyCost {
						sr.player.Eggs -= sr.player.HandBuyCost
						temp := float32(sr.player.HandBuyCost) * sr.player.HandBuyMulti
						sr.player.HandBuyCost = int(temp)
						UpdateBuyHandText(entities)
						UpdateEggs(entities)
					//If we don't have enough, do nothing, and reset clicker
					} else {
						clicker.Clicked = false
					}
				}
			}
		}
	}
}