package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemSpriteRender struct {
}

func (sr *SystemSpriteRender) Update(entities *[]ec.Entity, delta float64) {

}

func (sr *SystemSpriteRender) Draw(entities *[]ec.Entity, screen *ebiten.Image) {
	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		if e.HasComponent(ec.C_POSITION) && e.HasComponent(ec.C_SPRITE) {
			posComp := e.GetComponentWithID(ec.C_POSITION)
			var pos *ec.ComponentPosition = posComp.(*ec.ComponentPosition)
			x := pos.X
			y := pos.Y
			
			spriteComp := e.GetComponentWithID(ec.C_SPRITE)
			var sprite *ec.ComponentSprite = spriteComp.(*ec.ComponentSprite)
			img := sprite.Image

			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(x, y)
			screen.DrawImage(img, drawOp)
		} else if e.HasComponent(ec.C_SPATIAL) && e.HasComponent(ec.C_SPRITE) {
			sComp := e.GetComponentWithID(ec.C_SPATIAL)
			var spat *ec.ComponentSpatial = sComp.(*ec.ComponentSpatial)
			x := spat.X
			y := spat.Y
					
			spriteComp := e.GetComponentWithID(ec.C_SPRITE)
			var sprite *ec.ComponentSprite = spriteComp.(*ec.ComponentSprite)
			img := sprite.Image

			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(x, y)
			screen.DrawImage(img, drawOp)
		}	
	}
}