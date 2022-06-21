package inter

import "github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
import "github.com/hajimehoshi/ebiten/v2"

type WorldSpace interface {
	CreateEntity() *ec.Entity
}

type UpdateSystem interface {
	Update(WorldSpace, *[]*ec.Entity, float64)
}

type DrawSystem interface {
	Draw(WorldSpace, *[]*ec.Entity, *ebiten.Image)
}