package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	
)

type SystemBuyUpgrades struct {
	player *ec.ComponentPlayer
	ran bool
}

func (sr *SystemBuyUpgrades) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float32) {
	if !sr.ran {
		sr.ran = true
		//Loop through entities
		for i := range *entities {
			entity := (*entities)[i]
			//Set player for reference
			if entity.GetTag() == "Player" {
				sr.player = entity.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
			}
		}
		if sr.player == nil {
			panic("Couldn't find player in Buy Upgrades System! Did you add player? Is player entity missing 'Player' tag?")
		}
	}

	//Check the clickers
	for i := range *entities {
		e := (*entities)[i]
		//Make sure it has a clicker
		if e.HasComponent(ec.C_CLICKER) {
			cc := e.GetComponentWithID(ec.C_CLICKER)
			var clicker *ec.ComponentClicker = cc.(*ec.ComponentClicker)
			//If it's clicked?
			if clicker.Clicked {
				//Which clicker was clicked?
				if e.GetTag() == "Buy Hand" {
					//If we have enough eggs, buy a petter and increase price of petter
					if sr.player.Eggs >= sr.player.HandBuyCost {
						clicker.Clicked = false

						//Buy petter
						sr.player.Eggs -= sr.player.HandBuyCost
						temp := float32(sr.player.HandBuyCost) * sr.player.HandBuyMulti
						sr.player.HandBuyCost = int(temp)

						//Add petter
						sr.player.PetterCount++
	
					//If we don't have enough, do nothing, and reset clicker
					} else {
						clicker.Clicked = false
					}
				} else if e.GetTag() == "Buy Farmer" {
					//If we have enough eggs, buy a farmer and increase price of farmer
					if sr.player.Eggs >= sr.player.FarmerBuyCost {
						clicker.Clicked = false

						//Buy farmer
						sr.player.Eggs -= sr.player.FarmerBuyCost
						temp := float32(sr.player.FarmerBuyCost) * sr.player.FarmerBuyMulti
						sr.player.FarmerBuyCost = int(temp)

						//Add petter
						sr.player.FarmerCount++
	
					//If we don't have enough, do nothing, and reset clicker
					} else {
						clicker.Clicked = false
					}					
				}
			}
		}
	}
}