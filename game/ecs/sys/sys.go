package sys

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"

type UpdateSystem interface {
	Update(*[]ec.Entity, float64)
}

type DrawSystem interface {
	Draw(*[]ec.Entity, *ebiten.Image)
}