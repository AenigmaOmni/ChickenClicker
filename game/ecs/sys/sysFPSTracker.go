package sys

import (
	"fmt"

	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemFPSTracker struct {

}

func (sr *SystemFPSTracker) Update(world inter.WorldSpace, entities *[]*ec.Entity, delta float64) {
	for i := range *entities {
		if (*entities)[i].HasComponent(ec.C_FPSTRACKER) && (*entities)[i].HasComponent(ec.C_TEXT) {
			comp := (*entities)[i].GetComponentWithID(ec.C_TEXT)
			var t *ec.ComponentText = comp.(*ec.ComponentText)
			str := fmt.Sprintf("%f", ebiten.CurrentFPS())
			t.Message = str
		}
	}
}