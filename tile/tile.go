/*
@author: sk
@date: 2023/9/2
*/
package tile

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	pos    model.Vec3
	img    *ebiten.Image
	Type   model.TileType
	Enable bool
}

func (t *Tile) GetOrder() float64 { // 锚点保持与坦克一直否则权重太小了 Tile 与其他对象重叠时后显示
	return t.pos.Z*10000 + t.pos.Y + global.BlockSize/2 + t.pos.X + global.BlockSize/2 - 1 // X,Y 等权
}

func NewTile(pos model.Vec3, img *ebiten.Image, tileType model.TileType) *Tile {
	return &Tile{pos: pos, img: img, Enable: true, Type: tileType}
}

var (
	tileAnchor = model.NewVec2(32, 32)
)

func (t *Tile) Draw(screen *ebiten.Image) {
	if !t.Enable {
		return
	}
	pos := tool.Camera.World2Screen(t.pos)
	utils.DrawImage(screen, t.img, pos, tileAnchor, global.WindowRegion)
}

func (t *Tile) GetHeight(x float64, y float64) float64 {
	switch t.Type {
	case global.TileTypeSlope1:
		return t.pos.Z + global.BlockSize - (x - t.pos.X)
	case global.TileTypeSlope2:
		return t.pos.Z + global.BlockSize - (y - t.pos.Y)
	default:
		return t.pos.Z + global.BlockSize
	}
}
