package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SystemClickCollision struct {

}

func (sr *SystemClickCollision) Update(entities *[]ec.Entity, delta float64) {
	//If left mouse button released
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		//Loop over entities
		for i := range *entities {
			//If this entity has a spatial and a clicker component
			if (*entities)[i].HasComponent(ec.C_SPATIAL) && (*entities)[i].HasComponent(ec.C_CLICKER) {
				//Get spatial
				spatialC := (*entities)[i].GetComponentWithID(ec.C_SPATIAL)
				var spatial *ec.ComponentSpatial = spatialC.(*ec.ComponentSpatial)

				//Get specs for spatial collider
				x := spatial.X
				y := spatial.Y
				width := spatial.Width
				height := spatial.Height
				//parentID := spatial.GetEntityID()

				//Check if mouse point is inside rect
				mouseX, mouseY := ebiten.CursorPosition()

				if mouseX > int(x) &&
					mouseY > int(y) &&
					mouseX < int(x + width) &&
					mouseY < int(y + height) {

						clickC := (*entities)[i].GetComponentWithID(ec.C_CLICKER)
						var clicker *ec.ComponentClicker = clickC.(*ec.ComponentClicker)
						clicker.Clicked = true
				}
			}
		}
	}
}