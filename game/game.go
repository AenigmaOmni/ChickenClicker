package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	screenWidth int
	screenHeight int
	title string
	world World
	prevUpdateTime time.Time
}

func NewGame() Game {
	g := Game{}

	g.screenWidth = 1024 / 2
	g.screenHeight = 720 / 2
	g.title = "Project Adventure Time v0.0.1"

	g.world = NewWorld(g.screenWidth, g.screenHeight)

	ebiten.SetWindowSize(g.screenWidth*2, g.screenHeight*2)
	ebiten.SetWindowTitle(g.title)

	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) Update() error {
	//Calculate delta time
	delta := float64(time.Since(g.prevUpdateTime))
	g.prevUpdateTime = time.Now()

	g.world.Update(delta)
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}