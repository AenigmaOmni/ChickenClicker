package sys

import (
	"github.com/AenigmaOmni/ChickenClicker/game/inter"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemSpriteRender struct {
}

func (sr *SystemSpriteRender) Draw(world inter.WorldSpace, entities *[]*ec.Entity, screen *ebiten.Image) {

	//Draw layer 1

	//Loop through all entities and get entities with layer 1
	layer1 := getLayerEntities(entities, 1)

	//Loop through entities of layer1
	for i := 0; i < len(layer1); i++ {
		//Get entity
		e := layer1[i]
		drawSprite(screen, e)
	}

	//Draw layer 2

	//Loop through all entities and get entities with layer 1
	layer2 := getLayerEntities(entities, 2)

	//Loop through entities of layer1
	for i := 0; i < len(layer2); i++ {
		//Get entity
		e := layer2[i]
		drawSprite(screen, e)
	}

	//Draw layer 3

	//Loop through all entities and get entities with layer 1
	layer3 := getLayerEntities(entities, 3)

	//Loop through entities of layer1
	for i := 0; i < len(layer3); i++ {
		//Get entity
		e := layer3[i]
		drawSprite(screen, e)
	}

	//Draw layer 4

	//Loop through all entities and get entities with layer 1
	layer4 := getLayerEntities(entities, 4)

	//Loop through entities of layer1
	for i := 0; i < len(layer4); i++ {
		//Get entity
		e := layer4[i]
		drawSprite(screen, e)
	}
}

func getLayerEntities(entities *[]*ec.Entity, layerID int) []*ec.Entity {
	slice := make([]*ec.Entity, 0)
	for i := 0; i < len(*entities); i++ {
		e := (*entities)[i]

		if e.HasComponent(ec.C_POSITION) && e.HasComponent(ec.C_SPRITE) {
			c := e.GetComponentWithID(ec.C_SPRITE)
			var spr *ec.ComponentSprite = c.(*ec.ComponentSprite)
			if spr.Layer == layerID {
				slice = append(slice, (*entities)[i])
			}
		} else if e.HasComponent(ec.C_SPATIAL) && e.HasComponent(ec.C_SPRITE) {
			c := e.GetComponentWithID(ec.C_SPRITE)
			var spr *ec.ComponentSprite = c.(*ec.ComponentSprite)
			if spr.Layer == layerID {
				slice = append(slice, (*entities)[i])
			}
		}
	}

	return slice
}

func drawSprite(screen *ebiten.Image, e *ec.Entity) {
	//Check if this entity has a position comp and sprite comp
	if e.HasComponent(ec.C_POSITION) && e.HasComponent(ec.C_SPRITE) {
		posComp := e.GetComponentWithID(ec.C_POSITION)
		var pos *ec.ComponentPosition = posComp.(*ec.ComponentPosition)
		x := pos.X
		y := pos.Y
		
		spriteComp := e.GetComponentWithID(ec.C_SPRITE)
		var sprite *ec.ComponentSprite = spriteComp.(*ec.ComponentSprite)
		img := sprite.Image

		drawOp := &ebiten.DrawImageOptions{}
		drawOp.GeoM.Translate(x, y)
		screen.DrawImage(img, drawOp)
		//Check instead if this entity has a spatial and sprite comp
	} else if e.HasComponent(ec.C_SPATIAL) && e.HasComponent(ec.C_SPRITE) {
		sComp := e.GetComponentWithID(ec.C_SPATIAL)
		var spat *ec.ComponentSpatial = sComp.(*ec.ComponentSpatial)
		x := spat.X
		y := spat.Y
				
		spriteComp := e.GetComponentWithID(ec.C_SPRITE)
		var sprite *ec.ComponentSprite = spriteComp.(*ec.ComponentSprite)
		img := sprite.Image

		drawOp := &ebiten.DrawImageOptions{}
		drawOp.GeoM.Translate(x, y)
		screen.DrawImage(img, drawOp)
	}	
}