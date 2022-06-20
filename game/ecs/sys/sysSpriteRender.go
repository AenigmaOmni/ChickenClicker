package sys

import "github.com/hajimehoshi/ebiten/v2"
import comps "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
import "github.com/AenigmaOmni/ChickenClicker/game/ecs/entity"

type SystemSpriteRender struct {
}

func (sr *SystemSpriteRender) Update(entities []entity.Entity, delta float64) {

}

func (sr *SystemSpriteRender) Draw(entities []entity.Entity, screen *ebiten.Image) {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if e.HasComponent(comps.C_POSITION) && e.HasComponent(comps.C_SPRITE) {
			posComp := e.GetComponentWithID(comps.C_POSITION)
			var pos *comps.ComponentPosition = posComp.(*comps.ComponentPosition)
			x := pos.X
			y := pos.Y
			
			spriteComp := e.GetComponentWithID(comps.C_SPRITE)
			var sprite *comps.ComponentSprite = spriteComp.(*comps.ComponentSprite)
			img := sprite.Image

			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(x, y)
			screen.DrawImage(img, drawOp)
		}
	}
}