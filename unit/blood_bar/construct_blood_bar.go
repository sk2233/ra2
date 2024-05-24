/*
@author: sk
@date: 2023/9/9
*/
package blood_bar

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

func initConstructBloodBar() {
	img := tool.ImageLoader.LoadImage("res/image/blood_bar/construct/big.png")
	constructBloodBars[global.UnitBloodBarSizeBig] = utils.SplitImage(img, 1, 22)
	img = tool.ImageLoader.LoadImage("res/image/blood_bar/construct/middle.png")
	constructBloodBars[global.UnitBloodBarSizeMiddle] = utils.SplitImage(img, 1, 15)
	img = tool.ImageLoader.LoadImage("res/image/blood_bar/construct/small.png")
	constructBloodBars[global.UnitBloodBarSizeSmall] = utils.SplitImage(img, 1, 7)
	constructAnchors[global.UnitBloodBarSizeBig] = model.Vec2{X: 89, Y: -3}
	constructAnchors[global.UnitBloodBarSizeMiddle] = model.Vec2{X: 61, Y: -1}
	constructAnchors[global.UnitBloodBarSizeSmall] = model.Vec2{X: 29, Y: -1}
}

var (
	constructBloodBars = make(map[model.UnitBloodBarSize][]*ebiten.Image)
	constructAnchors   = make(map[model.UnitBloodBarSize]model.Vec2)
)

type constructBloodBar struct {
	images []*ebiten.Image
	anchor model.Vec2
}

func NewConstructBloodBar(param *model.UnitData) *constructBloodBar {
	size := param.BloodBarExtra.BloodBarSize
	height := param.ConstructExtra.Height
	return &constructBloodBar{
		images: constructBloodBars[size],
		anchor: constructAnchors[size].Add(model.Vec2{Y: float64(height * global.BlockSize)}),
	}
}

func (c *constructBloodBar) DrawBloodBar(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	if unit0.Hover || unit0.Select {
		index := utils.Min(unit0.Hp*len(c.images)/unit0.Data.Hp, len(c.images)-1)
		pos := tool.Camera.World2Screen(unit0.Pos)
		utils.DrawImage(screen, c.images[index], pos, c.anchor, global.WindowRegion)
	}
	if unit0.Select {
		extra := unit0.Data.ConstructExtra
		xNum, yNum, zNum := extra.WNum, extra.HNum, extra.Height
		down2 := unit0.Pos.Add(model.Vec3{X: float64(xNum * global.BlockSize)})
		down3 := down2.Add(model.Vec3{Y: float64(yNum * global.BlockSize)})
		down4 := unit0.Pos.Add(model.Vec3{Y: float64(yNum * global.BlockSize)})
		up1 := unit0.Pos.Add(model.Vec3{Z: float64(zNum * global.BlockSize)})
		up2 := down2.Add(model.Vec3{Z: float64(zNum * global.BlockSize)})
		up3 := down3.Add(model.Vec3{Z: float64(zNum * global.BlockSize)})
		up4 := down4.Add(model.Vec3{Z: float64(zNum * global.BlockSize)})
		sDown2 := tool.Camera.World2Screen(down2)
		sDown3 := tool.Camera.World2Screen(down3)
		sDown4 := tool.Camera.World2Screen(down4)
		sUp1 := tool.Camera.World2Screen(up1)
		sUp2 := tool.Camera.World2Screen(up2)
		sUp3 := tool.Camera.World2Screen(up3)
		sUp4 := tool.Camera.World2Screen(up4)
		utils.DrawLine(screen, sDown2, sUp2, 1, colornames.Black, global.WindowRegion)
		utils.DrawLine(screen, sDown3, sUp3, 1, colornames.Black, global.WindowRegion)
		utils.DrawLine(screen, sDown4, sUp4, 1, colornames.Black, global.WindowRegion)
		utils.DrawLine(screen, sUp1, sUp2, 1, colornames.Black, global.WindowRegion)
		utils.DrawLine(screen, sUp2, sUp3, 1, colornames.Black, global.WindowRegion)
		utils.DrawLine(screen, sUp3, sUp4, 1, colornames.Black, global.WindowRegion)
		utils.DrawLine(screen, sUp4, sUp1, 1, colornames.Black, global.WindowRegion)
	}
}
