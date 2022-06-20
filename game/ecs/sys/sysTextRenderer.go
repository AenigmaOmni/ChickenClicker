package sys

import (
	"image/color"

	comps "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SystemTextRenderer struct {

}

func (sr *SystemTextRenderer) Update(entities *[]ec.Entity, delta float64) {

}

func (sr *SystemTextRenderer) Draw(entities *[]ec.Entity, screen *ebiten.Image) {
	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]
		if e.HasComponent(comps.C_TEXT) && e.HasComponent(comps.C_POSITION) {
			comp := e.GetComponentWithID(comps.C_TEXT)
			var textCom *comps.ComponentText = comp.(*comps.ComponentText)

			coPos := e.GetComponentWithID(comps.C_POSITION)
			var pos *comps.ComponentPosition = coPos.(*comps.ComponentPosition)
			text.Draw(screen, textCom.Message, textCom.FontFace, int(pos.X), int(pos.Y), color.White)
		}
	}
}