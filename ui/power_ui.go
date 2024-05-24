/*
@author: sk
@date: 2023/9/16
*/
package ui

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

type PowerUI struct { // 16 * 520   黄电 100 缓冲区
}

func NewPowerUI() *PowerUI {
	return &PowerUI{}
}

var (
	powerRect = &model.Rect{
		Pos:    model.NewVec2(global.WindowW-144, global.WindowH-32), // 点在左下方
		Size:   model.NewVec2(16, 520),
		Anchor: model.NewVec2(0, 520),
	}
)

func (p *PowerUI) Draw(screen *ebiten.Image) {
	powerRect.Size.Y = 520
	powerRect.Anchor.Y = 520
	utils.FillRect(screen, powerRect, colornames.Black, global.MenuRegion)
	power, load := tool.DataManager.GetPower()
	scale := 1.0
	if power > 520 {
		scale = utils.Min(scale, 520/float64(power))
	}
	if load > 520 {
		scale = utils.Min(scale, 520/float64(load))
	}
	greenH := float64(power) * scale
	yellowH := utils.Min(float64(load+100)*scale, greenH)
	redH := float64(load) * scale
	if greenH > yellowH {
		powerRect.Size.Y = greenH - yellowH
		powerRect.Anchor.Y = greenH
		utils.FillRect(screen, powerRect, colornames.Green, global.MenuRegion)
	}
	if yellowH > redH {
		powerRect.Size.Y = yellowH - redH
		powerRect.Anchor.Y = yellowH
		utils.FillRect(screen, powerRect, colornames.Yellow, global.MenuRegion)
	}
	if redH > 0 {
		powerRect.Size.Y = redH
		powerRect.Anchor.Y = redH
		utils.FillRect(screen, powerRect, colornames.Red, global.MenuRegion)
	}
}
