package game

import (

	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	entityManager ec.EntityManager
	entities []ec.Entity
	updateSystems []sys.UpdateSystem
	drawSystems []sys.DrawSystem
}

func loadHUD(w *World, screenWidth int, screenHeight int) {
	//fps
	fpsCounter := w.entityManager.Create()
	fpsC := ec.NewTextComponent(32, 32, "FPS: 60")
	fpsCounter.AddComponent(&fpsC)
	fpsT := ec.NewComponentFPSTracker()
	fpsCounter.AddComponent(&fpsT)
	fpsP := ec.NewComponentPosition(float64(screenWidth - 80), 20)
	fpsCounter.AddComponent(&fpsP)
	w.AddEntity(fpsCounter)

	//egg counter
	eggCounter := w.entityManager.Create()
	eggC := ec.NewTextComponent(32, 32, "Eggs: 0")
	eggCounter.AddComponent(&eggC)
	eggP := ec.NewComponentPosition(10, 20)
	eggCounter.AddComponent(&eggP)
	w.AddEntity(eggCounter)
}

func loadSprites(w *World, screenWidth int, screenHeight int) {

	h := w.entityManager.Create()

	chickSprite := ec.NewComponentSprite("res/sprites/perfect_chicken.png")
	chickPos := ec.NewComponentPosition(float64(screenWidth / 2 - chickSprite.Width / 2),
		float64(screenHeight / 2 - chickSprite.Height / 2))
	h.AddComponent(&chickSprite)
	h.AddComponent(&chickPos)

	w.AddEntity(h)
}

func NewWorld(screenWidth int, screenHeight int) World {
	w := World{}
	w.entityManager = ec.NewEntityManager()
	
	loadSprites(&w, screenWidth, screenHeight)
	loadHUD(&w, screenWidth, screenHeight)

	w.AddDrawSystem(&sys.SystemTextRenderer{})
	w.AddUpdateSystem(&sys.SystemFPSTracker{})
	w.AddDrawSystem(&sys.SystemSpriteRender{})

	return w
}

func (w *World) AddDrawSystem(ds sys.DrawSystem) {
	w.drawSystems = append(w.drawSystems, ds)
}

func (w *World) AddUpdateSystem(us sys.UpdateSystem) {
	w.updateSystems = append(w.updateSystems, us)
}

func (w *World) AddEntity(ent ec.Entity) {
	w.entities = append(w.entities, ent)
}

func (w *World) Update(delta float64) {
	for i := 0; i < len(w.updateSystems); i++ {
		w.updateSystems[i].Update(&w.entities, delta)
	}
}

func (w *World) Draw(screen *ebiten.Image) {
	for i := 0; i < len(w.drawSystems); i++ {
		w.drawSystems[i].Draw(&w.entities, screen)
	}
}