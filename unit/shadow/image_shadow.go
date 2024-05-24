/*
@author: sk
@date: 2023/10/1
*/
package shadow

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	shadowImgs    = make(map[model.UnitType]*ebiten.Image)
	shadowAnchors = make(map[model.UnitType]model.Vec2)
)

func initImageShadow() {
	// 32 * 16
	shadowImgs[global.UnitTypeSoldier] = tool.ImageLoader.LoadImage("res/image/shadow/soldier.png")
	shadowAnchors[global.UnitTypeSoldier] = model.NewVec2(16, 8)
	// 64 * 32
	shadowImgs[global.UnitTypeVehicle] = tool.ImageLoader.LoadImage("res/image/shadow/vehicle.png")
	shadowAnchors[global.UnitTypeVehicle] = model.NewVec2(32, 16)
}

type imageShadow struct {
	unitType model.UnitType
}

func NewImageShadow(data *model.UnitData) *imageShadow {
	return &imageShadow{unitType: data.Type}
}

func (i *imageShadow) DrawShadow(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos3 := unit0.Pos
	pos3.Z = tool.TileManager.GetHeight(pos3.X, pos3.Y) // 只要地面tile的高度即可
	pos2 := tool.Camera.World2Screen(pos3)
	utils.DrawImage(screen, shadowImgs[i.unitType], pos2, shadowAnchors[i.unitType], global.WindowRegion)
}
