package mapLoader

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/AenigmaOmni/ChickenClicker/game/ecs/comps"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type MapData struct {
	TileHeight int `json:"tileheight"`
	TileWidth int `json:"tilewidth"`
	Height int `json:"height"`
	Width int `json:"width"`
	Layers []LayerData `json:"layers"`
	Tileset []TilesetData `json:"tilesets"`
}

type TilesetData struct {
	FirstGID int `json:"firstgid"`
	Columns int `json:"columns"`
	ImageName string `json:"image"`
	ImageHeight int `json:"imageheight"`
	ImageWidth int `json:"imagewidth"`
	TileCount int `json:"tilecount"`
	TileHeight int `json:"tileheight"`
	TileWidth int `json:"tilewidth"`
}

type LayerData struct {
	Height int `json:"height"`
	Width int `json:"width"`
	ID int `json:"id"`
	Name string `json:"name"`
	Opacity int `json:"opacity"`
	Visible bool `json:"visible"`
	Data []int `json:"data"`
}

func LoadJSONTiledMapComponent(path string) comps.ComponentMap {
	jsonF, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer jsonF.Close()

	byteValue, _ := ioutil.ReadAll(jsonF)

	var data MapData

	json.Unmarshal(byteValue, &data)
	
	m := comps.NewComponentMap()

	m.Height = int(data.Height)
	m.Width = int(data.Width)
	m.TileSize = int(data.TileHeight)
	m.LayerGround = data.Layers[0].Data
	m.LayerBottom = data.Layers[1].Data
	m.LayerMiddle = data.Layers[2].Data
	m.LayerTop = data.Layers[3].Data
	m.Columns = data.Tileset[0].Columns
	m.ImageWidth = data.Tileset[0].ImageWidth
	m.ImageHeight = data.Tileset[0].ImageHeight
	m.TileImageCount = data.Tileset[0].TileCount

	
	img, _, err := ebitenutil.NewImageFromFile("res/maps/" + data.Tileset[0].ImageName)
	
	if err != nil {
		panic(err)
	}
	
	m.Image = ebiten.NewImageFromImage(img) 

	return m 

}