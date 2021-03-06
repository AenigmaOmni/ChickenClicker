package ec

import "github.com/hajimehoshi/ebiten/v2"

type ComponentMap struct {
	ID          int
	Width       int
	Height      int
	TileSize    int
	Image       *ebiten.Image
	LayerGround []int
	LayerBottom []int
	LayerMiddle []int
	LayerTop    []int
	FirstGID int
	Columns int
	ImageHeight int
	ImageWidth int
	TileImageCount int
	parent Entity
}

func NewComponentMap() ComponentMap {
	c := ComponentMap{}
	c.ID = C_MAP

	return c
}

func (c *ComponentMap) GetID() int {
	return c.ID
}

func (c *ComponentMap) GetEntityID() int64 {
	return c.parent.GetID()
}