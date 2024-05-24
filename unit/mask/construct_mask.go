/*
@author: sk
@date: 2023/9/9
*/
package mask

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
)

type constructMask struct {
	wNum, hNum int
	height     int
}

func (c *constructMask) InRect(unit any, rect interfaces.IRect) bool {
	return false // 建筑不能框选
}

func (c *constructMask) PointIn(temp any, pos2 model.Vec2) bool {
	unit0 := temp.(*unit.Unit)
	grid := utils.Pos2Grid(unit0.Pos)
	for z := 0; z < c.height; z++ {
		pos3 := tool.Camera.Screen2WorldAtZ(pos2, float64(grid.Z+z)*global.BlockSize)
		hitGrid := utils.Pos2Grid(pos3)
		if hitGrid.X >= grid.X && hitGrid.X < grid.X+c.wNum && hitGrid.Y >= grid.Y && hitGrid.Y < grid.Y+c.hNum {
			return true
		}
	}
	return false
}

func NewConstructMask(param *model.UnitConstructExtra) *constructMask {
	return &constructMask{wNum: param.WNum, hNum: param.HNum, height: param.Height}
}
