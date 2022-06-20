package ec

const (
	C_POSITION = 0
	C_VELOCITY = 1
	C_SPRITE = 2
	C_ANIMATEDSPRITE = 3
	C_MAP = 4
	C_TEXT = 5
	C_FPSTRACKER = 6
	C_CLICKER = 7
	C_SPATIAL = 8
)

type Component interface {
	GetID() int
	GetEntityID() int64
}