package sys

import (
	"fmt"

	comps "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemFPSTracker struct {

}

func (sr *SystemFPSTracker) Update(entities []entity.Entity, delta float64) {
	for i := range entities {
		if entities[i].HasComponent(comps.C_FPSTRACKER) && entities[i].HasComponent(comps.C_TEXT) {
			comp := entities[i].GetComponentWithID(comps.C_TEXT)
			var t *comps.ComponentText = comp.(*comps.ComponentText)
			str := fmt.Sprintf("%f", ebiten.CurrentFPS())
			t.Message = str
		}
	}
}