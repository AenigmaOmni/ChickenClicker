package world

import "github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"

type EntityManager struct {
	count int64
	world *World
}

func NewEntityManager(w *World) EntityManager {
	eM := EntityManager{}
	eM.world = w
	eM.count = 0
	return eM
}

func (em *EntityManager) Create() *ec.Entity {
	e := ec.Entity{}
	e.Id = em.count
	e.SetTag("None")
	em.count++

	em.world.AddEntity(&e)

	return &e
}
