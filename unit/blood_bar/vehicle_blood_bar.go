/*
@author: sk
@date: 2023/9/29
*/
package blood_bar

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

// vehicle_blood_bar.go
var (
	vehicleSelectBloodBars []*ebiten.Image
	vehicleTouchBloodBars  []*ebiten.Image
)

func initVehicleBloodBar() {
	image := tool.ImageLoader.LoadImage("res/image/blood_bar/vehicle/select.png")
	vehicleSelectBloodBars = utils.SplitImage(image, 1, 31) // 每个 64 * 5
	image = tool.ImageLoader.LoadImage("res/image/blood_bar/vehicle/hover.png")
	vehicleTouchBloodBars = utils.SplitImage(image, 1, 31) // 每个 64 * 5
}

type vehicleBloodBar struct {
}

func NewVehicleBloodBar() *vehicleBloodBar {
	return &vehicleBloodBar{}
}

var (
	vehicleAnchor = model.NewVec2(32, 5)
)

func (s *vehicleBloodBar) DrawBloodBar(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos).Add(model.NewVec2(0, -32))
	index := utils.Min(unit0.Hp*31/unit0.Data.Hp, 30)
	if unit0.Select {
		utils.DrawImage(screen, vehicleSelectBloodBars[index], pos, vehicleAnchor, global.WindowRegion)
	} else if unit0.Hover {
		utils.DrawImage(screen, vehicleTouchBloodBars[index], pos, vehicleAnchor, global.WindowRegion)
	}
}
