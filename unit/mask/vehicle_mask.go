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

type vehicleMask struct {
}

func (s *vehicleMask) InRect(temp any, rect interfaces.IRect) bool {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos)
	return utils.PointCollision(pos, rect)
}

func NewVehicleMask() *vehicleMask {
	return &vehicleMask{}
}

var (
	vehicleRect = &model.Rect{Anchor: model.NewVec2(32, 8), Size: model.NewVec2(64, 16)}
)

func (s *vehicleMask) PointIn(temp any, pos model.Vec2) bool {
	unit0 := temp.(*unit.Unit)
	vehicleRect.Pos = tool.Camera.World2Screen(unit0.Pos)
	return utils.PointCollision(pos, vehicleRect)
}
