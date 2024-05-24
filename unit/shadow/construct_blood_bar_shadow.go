/*
@author: sk
@date: 2023/9/10
*/
package shadow

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

type constructBloodBarShadow struct {
}

func NewConstructBloodBarShadow() *constructBloodBarShadow {
	return &constructBloodBarShadow{}
}

func (c *constructBloodBarShadow) DrawShadow(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	if unit0.Select {
		extra := unit0.Data.ConstructExtra
		down := unit0.Pos
		up := down.Add(model.Vec3{Z: float64(extra.Height * global.BlockSize)})
		sDown := tool.Camera.World2Screen(down)
		sUp := tool.Camera.World2Screen(up)
		utils.DrawLine(screen, sDown, sUp, 1, colornames.Black, global.WindowRegion)
	}
}
