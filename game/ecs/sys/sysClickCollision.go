package sys

import (
	"fmt"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemClickCollision struct {

}

func (sr *SystemClickCollision) Update(entities *[]ec.Entity, delta float64) {
	for i := range *entities {
		if (*entities)[i].HasComponent(ec.C_SPATIAL) && (*entities)[i].HasComponent(ec.C_CLICKER) {
			spatialC := (*entities)[i].GetComponentWithID(ec.C_SPATIAL)
			var spatial *ec.ComponentSpatial = spatialC.(*ec.ComponentSpatial)

			x := spatial.X
			y := spatial.Y
			width := spatial.Width
			height := spatial.Height
			parentID := spatial.GetEntityID()

			
		}
	}
}