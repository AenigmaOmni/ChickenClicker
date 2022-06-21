package sys

import (
	"fmt"

	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

func UpdateEggs(entities *[]ec.Entity) {
	for i := range *entities {
		e := &(*entities)[i]
		//Get player
		if e.GetTag() == "Player" {
			//Get egg counter
			for j := range *entities {
				e2 := &(*entities)[j]
				if e2.GetTag() == "Egg Counter" {
					if e2.HasComponent(ec.C_TEXT) {
						//Update egg text
						var text *ec.ComponentText = e2.GetComponentWithID(ec.C_TEXT).(*ec.ComponentText)
						var player *ec.ComponentPlayer = e.GetComponentWithID(ec.C_PLAYER).(*ec.ComponentPlayer)
						text.Message = fmt.Sprintf("Eggs: %v", player.Eggs)
						return
					} else {
						panic("Egg Counter does not have a text component :(")
					}
				}
			}	
		}
	}

	panic("Couldn't find player in UpdateEggs function")
}

func UpdateBuyHandText(entities *[]ec.Entity) {
	for i := range *entities {
		e := &(*entities)[i]
		//Get player
		if e.GetTag() == "Player" {
			//Get hand buy text
			for j := range *entities {
				e2 := &(*entities)[j]
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