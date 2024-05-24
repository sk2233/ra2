/*
@author: sk
@date: 2023/9/29
*/
package mask

import (
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
)

type soldierMask struct {
}

func (s *soldierMask) InRect(temp any, rect interfaces.IRect) bool {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos)
	return utils.PointCollision(pos, rect)
}

func NewSoldierMask() *soldierMask {
	return &soldierMask{}
}

var (
	soldierRect = &model.Rect{Anchor: model.NewVec2(4, 32), Size: model.NewVec2(8, 32)}
)

func (s *soldierMask) PointIn(temp any, pos model.Vec2) bool {
	unit0 := temp.(*unit.Unit)
	soldierRect.Pos = tool.Camera.World2Screen(unit0.Pos)
	return utils.PointCollision(pos, soldierRect)
}
