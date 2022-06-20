package ec

import (
	"golang.org/x/image/font"
	"github.com/AenigmaOmni/ChickenClicker/game/loader/fontLoader"
)

type ComponentText struct {
	FontFace font.Face
	Message string
	ID int
	parent Entity
}

func NewTextComponent(size float64, dpi float64, msg string) ComponentText {
	ct := ComponentText{}
	ct.ID = C_TEXT
	ct.FontFace = fontLoader.NewFontFace(size, dpi)
	ct.Message = msg

	return ct
}

func (t *ComponentText) GetID() int {
	return t.ID
}

func (c *ComponentText) GetEntityID() int64 {
	return c.parent.GetID()
}