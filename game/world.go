package game

import (

	"github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/entity"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	entityManager entity.EntityManager
	entities []entity.Entity
	updateSystems []sys.UpdateSystem
	drawSystems []sys.DrawSystem
}

func loadHUD(w *World) {
	hud := w.entityManager.Create()
	
	//fps
	fpsC := comps.NewTextComponent(32, 32, "FPS: 60")
	hud.AddComponent(&fpsC)
	fpsT := comps.NewComponentFPSTracker()
	hud.AddComponent(&fpsT)
	fpsP := comps.NewComponentPosition()
	fpsP.X = 10
	fpsP.Y = 20
	hud.AddComponent(&fpsP)

	w.AddEntity(hud)
}

func loadSprites(w *World, screenWidth int, screenHeight int) {
	w.entityManager = entity.NewEntityManager()

	h := w.entityManager.Create()

	chickSprite := comps.NewComponentSprite("res/sprites/perfect_chicken.png")
	chickPos := comps.NewComponentPosition()
	chickPos.X = float64(screenWidth / 2 - chickSprite.Width / 2)  
	chickPos.Y = float64(screenHeight / 2 - chickSprite.Height / 2)

	h.AddComponent(&chickSprite)
	h.AddComponent(&chickPos)

	w.AddEntity(h)
}

func NewWorld(screenWidth int, screenHeight int) World {
	w := World{}

	loadSprites(&w, screenWidth, screenHeight)
	loadHUD(&w)

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

func (w *World) AddEntity(ent entity.Entity) {
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