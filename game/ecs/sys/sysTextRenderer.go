package sys

import (
	"image/color"
	comps "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SystemTextRenderer struct {

}

func (sr *SystemTextRenderer) Update(entities []entity.Entity, delta float64) {

}

func (sr *SystemTextRenderer) Draw(entities []entity.Entity, screen *ebiten.Image) {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if e.HasComponent(comps.C_TEXT) {
			comp := e.GetComponentWithID(comps.C_TEXT)
			var textCom *comps.ComponentText = comp.(*comps.ComponentText)
			text.Draw(screen, textCom.Message, textCom.FontFace, 20, 20, color.White)
		}
	}
}