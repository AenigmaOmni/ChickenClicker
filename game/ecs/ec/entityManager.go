package ec

type EntityManager struct {
	count int64
}

func NewEntityManager() EntityManager {
	eM := EntityManager{}
	eM.count = 0
	return eM
}

func (em *EntityManager) Create() Entity {
	e := Entity{}

	e.id = em.count
	em.count++

	return e
}