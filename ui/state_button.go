/*
@author: sk
@date: 2023/9/16
*/
package ui

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StateButton struct {
	rect    *model.Rect
	text    string
	support interfaces.IStateButtonSupport
	tag     model.ButtonTag
}

func (s *StateButton) HandleCursor(pos model.Vec2) bool {
	if !utils.PointCollision(pos, s.rect) {
		return false
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.support.OnClick(s.tag)
	}
	return true
}

func (s *StateButton) Draw(screen *ebiten.Image) {
	utils.FillRect(screen, s.rect, colornames.Silver, global.MenuRegion)
	utils.StrokeRect(screen, s.rect, 1, colornames.Black, global.MenuRegion)
	textConfig.Color = colornames.Red
	if s.support.GetState(s.tag) {
		textConfig.Color = colornames.Yellow
	}
	tool.TextDraw.DrawText(screen, s.rect.Middle(), s.text, textConfig)
}

var (
	textConfig *model.TextConfig
)

func init() {
	textConfig = &model.TextConfig{
		Anchor: model.NewVec2(0.5, 0.5),
		Color:  colornames.Black,
		Size:   1,
		Region: global.MenuRegion,
	}
}

func NewStateButton(pos, size model.Vec2, tag model.ButtonTag, text string,
	support interfaces.IStateButtonSupport) *StateButton {
	return &StateButton{
		rect: &model.Rect{
			Pos:  pos,
			Size: size,
		},
		tag:     tag,
		text:    text,
		support: support,
	}
}
