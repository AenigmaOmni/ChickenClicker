package sys

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/AenigmaOmni/ChickenClicker/game/ecs/entity"

type UpdateSystem interface {
	Update(*[]entity.Entity, float64)
}

type DrawSystem interface {
	Draw(*[]entity.Entity, *ebiten.Image)
}