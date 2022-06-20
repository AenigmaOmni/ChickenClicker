package fontLoader

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

func NewFontFace(size float64, dpi float64) font.Face {
	tt, err := opentype.Parse(goregular.TTF)

	if err != nil {
		panic(err)
	}

	normalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: size,
		DPI: dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}

	return normalFont
}