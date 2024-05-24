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
	"strconv"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

type MoneyUI struct {
	rect       *model.Rect
	textPos    model.Vec2
	textConfig *model.TextConfig
}

func NewMoneyUI() *MoneyUI { // 144 * 16
	pos := model.NewVec2(global.WindowW-144, 0)
	size := model.NewVec2(144, 16)
	return &MoneyUI{
		rect: &model.Rect{
			Pos:  pos,
			Size: size,
		},
		textPos: pos.Add(size.Scale(0.5)),
		textConfig: &model.TextConfig{
			Anchor: model.NewVec2(0.5, 0.5),
			Color:  colornames.Gold,
			Size:   1,
			Region: global.MenuRegion,
		},
	}
}

func (m *MoneyUI) Draw(screen *ebiten.Image) {
	utils.FillRect(screen, m.rect, colornames.Black, global.MenuRegion)
	money := tool.DataManager.GetMoney()
	tool.TextDraw.DrawText(screen, m.textPos, strconv.FormatInt(money, 10), m.textConfig)
}
