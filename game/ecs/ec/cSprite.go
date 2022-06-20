package ec

import (
	_ "image/png"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type ComponentSprite struct {
	ID int
	Image *ebiten.Image
	Width int
	Height int
	parent Entity
}


func NewComponentSprite(path string) ComponentSprite {
	s := ComponentSprite{}
	s.ID = C_SPRITE

	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		panic(err)
	}

	s.Image = ebiten.NewImageFromImage(img)
	s.Width, s.Height = s.Image.Size()

	return s
}

func (s *ComponentSprite) GetID() int {
	return s.ID
}

func (c *ComponentSprite) GetEntityID() int64 {
	return c.parent.GetID()
}