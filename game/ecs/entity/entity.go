package entity

import "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"

type Entity struct {
	id int64
	components []comps.Component
}

func (e *Entity) AddComponent(comp comps.Component) {
	e.components = append(e.components, comp)
}

func (e *Entity) GetComponents() []comps.Component {
	return e.components
}

func (e *Entity) GetComponentWithID(id int) comps.Component {
	for i := 0; i < len(e.components); i++ {
		c := e.components[i]
		if c.GetID() == id {
			return c
		}
	}
	panic("Entity couldn't find that component to return")
}

func (e *Entity) HasComponent(id int) bool {
	for i := 0; i < len(e.components); i++ {
		c := e.components[i]
		if c.GetID() == id {
			return true
		}
	}
	return false
}