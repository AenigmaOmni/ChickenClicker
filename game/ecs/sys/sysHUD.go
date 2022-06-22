package sys

import (
	"fmt"
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemHUD struct {
	ran bool
	player *ec.ComponentPlayer
}

func (sr *SystemHUD) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float32) {
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
			panic("Couldn't find player in HUD System! Did you add player? Is player entity missing 'Player' tag?")
		}
	}


	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		//Get the hud timer entity and component
		if e.HasComponent(ec.C_TIMER) && e.GetTag() == "HUD Timer" {
			timer := e.GetComponentWithID(ec.C_TIMER).(*ec.ComponentTimer)
			if timer.Alarm {
				//Set alarm to false and commence updating hud
				timer.Alarm = false
				//Update interface
				updateBuyHandText(entities)
				updateEggsAndIncome(entities)
				updateBuyFarmerText(entities)
				updateUpgradeCountText(entities)
			}
		}
	}
}

func updateUpgradeCountText(entities *[]*ec.Entity) {
	for i := range *entities {
		e := (*entities)[i]
		//Get player
		if e.GetTag() == "Player" {
			var player *ec.ComponentPlayer = e.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
			//Loop over entities
			for j := range *entities {
				e2 := (*entities)[j]
				if e2.HasComponent(ec.C_TEXT) {
					var text *ec.ComponentText = e2.GetComponentWithID(ec.C_TEXT).(*ec.ComponentText)
					
					if e2.GetTag() == "Petter Upgrade Count" {
						text.Message = fmt.Sprintf("Petters: %v", player.PetterCount)	
					} else if e2.GetTag() == "Farmer Upgrade Count" {
						text.Message = fmt.Sprintf("Farmers: %v", player.FarmerCount)
					}
				}
			}		
		}
	}
}

func updateEggsAndIncome(entities *[]*ec.Entity) {
	for i := range *entities {
		e := (*entities)[i]
		//Get player
		if e.GetTag() == "Player" {
			//Get egg counter and/or egg income
			for j := range *entities {
				e2 := (*entities)[j]
				if e2.GetTag() == "Egg Counter" {
					if e2.HasComponent(ec.C_TEXT) {
						//Update egg text
						var text *ec.ComponentText = e2.GetComponentWithID(ec.C_TEXT).(*ec.ComponentText)
						var player *ec.ComponentPlayer = e.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
						text.Message = fmt.Sprintf("Eggs: %v", player.Eggs)
					} else {
						panic("Egg Counter does not have a text component :(")
					}
				} else if e2.GetTag() == "Egg Income" {
					if e2.HasComponent(ec.C_TEXT) {
						//Update income text
						var text *ec.ComponentText = e2.GetComponentWithID(ec.C_TEXT).(*ec.ComponentText)
						var player *ec.ComponentPlayer = e.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)

						income := 0
						income += (player.HandProduction * player.PetterCount)
						income += (player.FarmerProduction * player.FarmerCount)

						text.Message = fmt.Sprintf("Income: %v", income)
					} else {
						panic("Egg Income does not have a text component :(")
					}
				}
			}
			return	
		}
	}

	panic("Couldn't find player in UpdateEggs function")
}

func updateBuyHandText(entities *[]*ec.Entity) {
	for i := range *entities {
		e := (*entities)[i]
		//Get player
		if e.GetTag() == "Player" {
			//Get hand buy text
			for j := range *entities {
				e2 := (*entities)[j]
				if e2.GetTag() == "Hand Buy Text" {
					if e2.HasComponent(ec.C_TEXT) {
						//Update hand buy text
						var text *ec.ComponentText = e2.GetComponentWithID(ec.C_TEXT).(*ec.ComponentText)
						var player *ec.ComponentPlayer = e.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
						text.Message = fmt.Sprintf("Buy Petter: %v", player.HandBuyCost)
						return
					} else {
						panic("Buy hand text does not have a text component :(")
					}
				}
			}	
		}
	}

	panic("Couldn't find player in UpdateBuyHandText function")
}

func updateBuyFarmerText(entities *[]*ec.Entity) {
	for i := range *entities {
		e := (*entities)[i]
		//Get player
		if e.GetTag() == "Player" {
			//Get farmer buy text
			for j := range *entities {
				e2 := (*entities)[j]
				if e2.GetTag() == "Farmer Buy Text" {
					if e2.HasComponent(ec.C_TEXT) {
						//Update farmer buy text
						var text *ec.ComponentText = e2.GetComponentWithID(ec.C_TEXT).(*ec.ComponentText)
						var player *ec.ComponentPlayer = e.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
						text.Message = fmt.Sprintf("Buy Farmer: %v", player.FarmerBuyCost)
						return
					} else {
						panic("Buy farmer text does not have a text component :(")
					}
				}
			}	
		}
	}

	panic("Couldn't find player in UpdateBuyFarmerText function")
}