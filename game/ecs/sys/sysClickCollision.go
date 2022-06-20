package sys

import (
	"fmt"

	comps "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemClickCollision struct {

}

func (sr *SystemClickCollision) Update(entities *[]entity.Entity, delta float64) {
	for i := range *entities {
		if (*entities)[i].HasComponent(comps.C_SPATIAL) && (*entities)[i].HasComponent(comps.C_CLICKER) {
			spatialC := (*entities)[i].GetComponentWithID(comps.C_SPATIAL)
			var spatial *comps.ComponentSpatial = spatialC.(*comps.ComponentSpatial)

			x := spatial.X
			y := spatial.Y
			width := spatial.Width
			height := spatial.Height

			
		}
	}
}