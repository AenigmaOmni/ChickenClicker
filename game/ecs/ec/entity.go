package ec

type Entity struct {
	tag string
	id int64
	components []Component
}

func (e *Entity) SetTag(t string) {
	e.tag = t
}

func (e *Entity) GetTag() string {
	return e.tag
}

func (e *Entity) AddComponent(comp Component) {
	e.components = append(e.components, comp)
}

func (e *Entity) GetComponents() []Component {
	return e.components
}

func (e *Entity) GetComponentWithID(id int) Component {
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

func (e *Entity) GetID() int64 {
	return e.id
}