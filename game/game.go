package game

import (
	"time"

	"github.com/AenigmaOmni/ChickenClicker/game/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	screenWidth int
	screenHeight int
	title string
	world world.World
	prevTime int64
}

func NewGame() Game {
	g := Game{}

	g.screenWidth = 1024 / 2
	g.screenHeight = 720 / 2
	g.title = "Chicken Clicker v0.2.0"

	g.world = world.NewWorld(g.screenWidth, g.screenHeight)

	g.prevTime = time.Now().UnixMilli()

	ebiten.SetWindowSize(g.screenWidth*2, g.screenHeight*2)
	ebiten.SetWindowTitle(g.title)

	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) Update() error {
	//Calculate delta time
	timeNow := time.Now().UnixNano()
	deltaTime := float64(((timeNow - g.prevTime) / 1000000)) * 0.001
	g.prevTime = timeNow
	
	delta := float32(deltaTime)

	g.world.Update(delta)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}