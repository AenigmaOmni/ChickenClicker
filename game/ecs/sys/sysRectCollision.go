package sys

import (
	"fmt"

	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
)

type SystemRectCollision struct {

}

func (sr *SystemRectCollision) Update(world inter.WorldSpace, entities *[]ec.Entity, delta float32) {
	for i := range *entities {
		if (*entities)[i].HasComponent(ec.C_SPATIAL) && (*entities)[i].HasComponent(ec.C_CLICKER) {
			//Get an entity with spatial
			spatialC := (*entities)[i].GetComponentWithID(ec.C_SPATIAL)
			var spatial *ec.ComponentSpatial = spatialC.(*ec.ComponentSpatial)

			x := spatial.X
			y := spatial.Y
			width := spatial.Width
			height := spatial.Height
			parentID := spatial.GetEntityID()

			//Get all entities with spatial
			entitiesWithSpat := make([]*ec.Entity, 0)
			for j := range *entities {
				if (*entities)[j].HasComponent(ec.C_SPATIAL) {
					entitiesWithSpat = append(entitiesWithSpat, &(*entities)[j])
				}
			}

			//Now check if our spatial collides with those spatials
			for k := range entitiesWithSpat {
				c := entitiesWithSpat[k].GetComponentWithID(ec.C_SPATIAL)
				var spatial2 *ec.ComponentSpatial = c.(*ec.ComponentSpatial)

				tx := spatial2.X
				ty := spatial2.Y
				twidth := spatial2.Width
				theight := spatial2.Height
				tparentID := spatial2.GetEntityID()

				//Check if this is not the same entity first
				if parentID != tparentID {
					//Check for collision
					if x < tx + twidth && x + width > tx && y < ty + theight && y + height > ty {
						fmt.Println("Collided! O.O")
					}
				}
			}
			
		}
	}
}