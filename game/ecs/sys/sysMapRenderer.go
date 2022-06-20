package sys

import (
	"image"

	comps "github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/AenigmaOmni/ChickenClicker/game/ecs/ec"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemMapRenderer struct {
}

func (s *SystemMapRenderer) Update(entities []ec.Entity, delta float64) {

}

func (s *SystemMapRenderer) Draw(entities []ec.Entity, screen *ebiten.Image) {
	for i := range entities {
		if entities[i].HasComponent(comps.C_MAP) {
			mc := entities[i].GetComponentWithID(comps.C_MAP)
			var m *comps.ComponentMap = mc.(*comps.ComponentMap)

			tileSize := int(m.TileSize)
			cols := m.Columns

			//Get GIDs and match with coordinates on image then store in struct
			coordData := make([]GID, 0)
			gd := 1
			tx := 0 - tileSize
			ty := 0
			for j := 1; j < m.TileImageCount; j++ {
				gidData := GID{}
				gidData.X = tx
				gidData.Y = ty
				gidData.ID = gd
				gd++
				tx += tileSize
				if tx >= cols*tileSize {
					tx = 0
					ty += tileSize
				}
				coordData = append(coordData, gidData)
			}

			//declare some helper stuff
			img := m.Image
			var targetX float64 = 0.0
			var targetY float64 = 0.0
			width := m.Width

			//draw ground layer
			for j := range m.LayerGround {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(targetX, targetY)
				gid := m.LayerGround[j]
				sx := coordData[gid].X
				sy := coordData[gid].Y
				screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
				targetX += float64(tileSize)
				if targetX >= float64(width*tileSize) {
					targetX = 0
					targetY += float64(tileSize)
				}
			}

			//draw bottom layer
			targetX = 0.0
			targetY = 0.0
			for j := range m.LayerBottom {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(targetX, targetY)
				gid := m.LayerBottom[j]
				sx := coordData[gid].X
				sy := coordData[gid].Y
				screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
				targetX += float64(tileSize)
				if targetX >= float64(width*tileSize) {
					targetX = 0
					targetY += float64(tileSize)
				}
			}

			//draw medium layer
			targetX = 0.0
			targetY = 0.0
			for j := range m.LayerMiddle {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(targetX, targetY)
				gid := m.LayerMiddle[j]
				sx := coordData[gid].X
				sy := coordData[gid].Y
				screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
				targetX += float64(tileSize)
				if targetX >= float64(width*tileSize) {
					targetX = 0
					targetY += float64(tileSize)
				}
			}

			//draw top layer
			targetX = 0.0
			targetY = 0.0
			for j := range m.LayerTop {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(targetX, targetY)
				gid := m.LayerTop[j]
				sx := coordData[gid].X
				sy := coordData[gid].Y
				screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
				targetX += float64(tileSize)
				if targetX >= float64(width*tileSize) {
					targetX = 0
					targetY += float64(tileSize)
				}
			}
		}
	}
}

type GID struct {
	X  int
	Y  int
	ID int
}
