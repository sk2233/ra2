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

var (
	soldierSelectBloodBars []*ebiten.Image
	soldierTouchBloodBars  []*ebiten.Image
)

func initSoldierBloodBar() {
	image := tool.ImageLoader.LoadImage("res/image/blood_bar/soldier/select.png")
	soldierSelectBloodBars = utils.SplitImage(image, 1, 15) // 每个 32 * 5
	image = tool.ImageLoader.LoadImage("res/image/blood_bar/soldier/hover.png")
	soldierTouchBloodBars = utils.SplitImage(image, 1, 15) // 每个 32 * 5
}

type soldierBloodBar struct {
}

func NewSoldierBloodBar() *soldierBloodBar {
	return &soldierBloodBar{}
}

var (
	soldierAnchor = model.NewVec2(16, 5)
)

func (s *soldierBloodBar) DrawBloodBar(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos).Add(model.NewVec2(0, -32))
	index := utils.Min(unit0.Hp*15/unit0.Data.Hp, 14)
	if unit0.Select {
		utils.DrawImage(screen, soldierSelectBloodBars[index], pos, soldierAnchor, global.WindowRegion)
	} else if unit0.Hover {
		utils.DrawImage(screen, soldierTouchBloodBars[index], pos, soldierAnchor, global.WindowRegion)
	}
}
