package sys

import (
	"image/color"

	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SystemTextRenderer struct {

}

func (sr *SystemTextRenderer) Draw(world inter.WorldSpace, entities *[]*ec.Entity, screen *ebiten.Image) {
	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		if e.HasComponent(ec.C_TEXT) && e.HasComponent(ec.C_POSITION) {
			comp := e.GetComponentWithID(ec.C_TEXT)
			var textCom *ec.ComponentText = comp.(*ec.ComponentText)

			coPos := e.GetComponentWithID(ec.C_POSITION)
			var pos *ec.ComponentPosition = coPos.(*ec.ComponentPosition)
			text.Draw(screen, textCom.Message, textCom.FontFace, int(pos.X), int(pos.Y), color.White)
		}
	}
}