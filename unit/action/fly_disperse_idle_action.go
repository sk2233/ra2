/*
@author: sk
@date: 2023/10/1
*/
package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
)

func initFlyDisperseIdleAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeFlyDisperseIdle, &flyDisperseIdleActionCreator{})
}

type flyDisperseIdleActionCreator struct {
}

func (n *flyDisperseIdleActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeCreateIdle {
		return nil
	}
	return &flyDisperseIdleAction{data: param, landEnd: true}
}

type flyDisperseIdleAction struct {
	data    *model.UnitActionCreatorExtra
	height  float64 // 要降落到的高度
	landEnd bool    // 是否降落完毕
}

func (n *flyDisperseIdleAction) OnTop(temp any) { // OnBelow 会自动尝试释放位置，这里可以直接占用位置
	unit0 := temp.(*unit.Unit)
	grid := unit.GetGrid(unit0)
	// 回归idle状态后试图锁当前格子，失败寻找附近格子
	if tool.UnitManager.CheckGrid(grid, unit0.Data) {
		pos := tool.UnitManager.GetPos(grid, unit0.Data)
		pos.Z += global.FlyHeight
		unit0.Pos = pos
		tool.UnitManager.AddToGrid(grid, unit0)
		if n.data.NeedLand {
			pos.Z -= global.FlyHeight // 判断是否可以降落
			data := utils.Clone(unit0.Data)
			data.StandExtra.Land = true
			if tool.UnitManager.CheckGrid(utils.Pos2Grid(pos).Down(), data) {
				n.landEnd = false
				n.height = pos.Z
			} else { // 不可以降落
				n.landEnd = true
			}
		}
	} else {
		res := tool.TileManager.GetNearGrids(grid, []*unit.Unit{unit0}).(map[*unit.Unit]model.Grid)
		if len(res) > 0 {
			unit0.Execute(&model.UnitCmd{ // 试图移动到对应位置
				Type: global.UnitCmdTypeMove,
				Grid: res[unit0],
			})
		} else { // 无法处理，单位太多？地图太小？
			utils.LogErr("[flyDisperseIdleAction] GetNearGrids no space")
		}
	}
}

func (n *flyDisperseIdleAction) OnBelow(temp any) {
	unit0 := temp.(*unit.Unit)
	grid := unit.GetGrid(unit0)
	tool.UnitManager.RemoveFromGrid(grid, unit0) // 其他行为不占用格子，尝试释放自己占用的空间
}

func (n *flyDisperseIdleAction) UpdateAction(temp any) {
	if !n.data.NeedLand || n.landEnd {
		return
	}
	unit0 := temp.(*unit.Unit)
	unit0.Pos.Z -= n.data.MoveSpeed // 默认移动速度是负数(下降)
	if unit0.Pos.Z < n.height {
		unit0.Pos.Z = n.height
		n.landEnd = true
	}
}
