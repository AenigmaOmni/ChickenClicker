package game

import (
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	entityManager ec.EntityManager
	entities      []ec.Entity
	updateSystems []sys.UpdateSystem
	drawSystems   []sys.DrawSystem
}

//Load the player data structure
func loadPlayer(w *World) {
	p := w.entityManager.Create()
	playerC := ec.NewComponentPlayer()
	p.AddComponent(&playerC)
	p.SetTag("Player")
	w.AddEntity(p)
}

//Load the buy hud
func loadBuyHUD(w *World, screenWidth int, screenHeight int) {
	handBuy := w.entityManager.Create()
	handBuy.SetTag("Buy Hand")
	handSprite := ec.NewComponentSprite("res/sprites/button.png", 2)
	handSpat := ec.NewComponentSpatial(330, 35, float64(handSprite.Width), float64(handSprite.Height))
	handClicker := ec.NewComponentClicker()
	handBuy.AddComponent(&handClicker)
	handBuy.AddComponent(&handSprite)
	handBuy.AddComponent(&handSpat)
	w.AddEntity(handBuy)

	handBuyText := w.entityManager.Create()
	handBuyText.SetTag("Hand Buy Text")
	handBuyTextC := ec.NewTextComponent(32, 32, "Buy Petter")
	handBuyTextPos := ec.NewComponentPosition(345, 65)
	handBuyText.AddComponent(&handBuyTextC)
	handBuyText.AddComponent(&handBuyTextPos)
	w.AddEntity(handBuyText)
}

//Load interface elements
func loadHUD(w *World, screenWidth int, screenHeight int) {
	//hud background
	hudBG := w.entityManager.Create()
	spriteC := ec.NewComponentSprite("res/sprites/hud.png", 1)
	hudBG.AddComponent(&spriteC)
	hudBGPos := ec.NewComponentPosition(0, 0)
	hudBG.AddComponent(&hudBGPos)
	w.AddEntity(hudBG)

	//fps
	fpsCounter := w.entityManager.Create()
	fpsC := ec.NewTextComponent(32, 32, "FPS: 60")
	fpsCounter.AddComponent(&fpsC)
	fpsT := ec.NewComponentFPSTracker()
	fpsCounter.AddComponent(&fpsT)
	fpsP := ec.NewComponentPosition(float64(screenWidth-80), 20)
	fpsCounter.AddComponent(&fpsP)
	w.AddEntity(fpsCounter)

	//egg counter
	eggCounter := w.entityManager.Create()
	eggCounter.SetTag("Egg Counter")
	eggC := ec.NewTextComponent(32, 32, "Eggs: 0")
	eggCounter.AddComponent(&eggC)
	eggP := ec.NewComponentPosition(10, 20)
	eggCounter.AddComponent(&eggP)
	w.AddEntity(eggCounter)

	//pet text
	petStr := w.entityManager.Create()
	petStr.SetTag("Pet Text")
	petTC := ec.NewTextComponent(40, 40, "Pet the Chicken for Eggs!")
	petStr.AddComponent(&petTC)
	petPos := ec.NewComponentPosition(55, float64(screenHeight/2-100))
	petStr.AddComponent(&petPos)
	w.AddEntity(petStr)
}

//Load sprites
func loadSprites(w *World, screenWidth int, screenHeight int) {
	//Load chicken
	h := w.entityManager.Create()
	h.SetTag("Chicken")
	chickSprite := ec.NewComponentSprite("res/sprites/perfect_chicken.png", 2)
	chickPos := ec.NewComponentSpatial(145,
		float64(screenHeight/2-chickSprite.Height/2), float64(chickSprite.Width), float64(chickSprite.Height))
	chickClicker := ec.NewComponentClicker()
	h.AddComponent(&chickClicker)
	h.AddComponent(&chickSprite)
	h.AddComponent(&chickPos)

	w.AddEntity(h)
}

//Load systems
func loadSystems(w *World) {
	
	//Load update systems
	w.AddUpdateSystem(&sys.SystemFPSTracker{})
	w.AddUpdateSystem(&sys.SystemClickCollision{})
	w.AddUpdateSystem(sys.NewSystemClickerEgg())
	w.AddUpdateSystem(&sys.SystemBuyUpgrades{})

	//Load draw systems
	w.AddDrawSystem(&sys.SystemSpriteRender{})
	//Add text renderer after srite renderer
	w.AddDrawSystem(&sys.SystemTextRenderer{})
}

//Create new world
func NewWorld(screenWidth int, screenHeight int) World {
	w := World{}
	w.entityManager = ec.NewEntityManager()

	loadPlayer(&w)
	loadSprites(&w, screenWidth, screenHeight)
	loadHUD(&w, screenWidth, screenHeight)
	loadBuyHUD(&w, screenWidth, screenHeight)

	loadSystems(&w)
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
