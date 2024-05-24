/*
@author: sk
@date: 2023/10/1
*/
package action

import (
	"math"
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initFlyHomeIdleAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeFlyHomeIdle, &flyHomeIdleActionCreator{})
}

type flyHomeIdleActionCreator struct {
}

func (n *flyHomeIdleActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeCreateIdle {
		return nil
	}
	// TODO 这个位置应该由飞机场给出并绑定飞机场的
	unit0 := temp.(*unit.Unit)
	height := tool.TileManager.GetHeight(unit0.Pos.X, unit0.Pos.Y)
	return &flyHomeIdleAction{speed: param.MoveSpeed, landEnd: true, grid: unit.GetGrid(unit0), height: height}
}

type flyHomeIdleAction struct {
	speed   float64    // 下降的速度
	grid    model.Grid // 要回到的位置
	height  float64    // 要下降到的高度，是个定值，由grid计算得到，这里保存起来方便使用
	landEnd bool       // 是否降落完毕
}

func (n *flyHomeIdleAction) OnTop(temp any) { // OnBelow 会自动尝试释放位置，这里可以直接占用位置
	unit0 := temp.(*unit.Unit)
	grid := unit.GetGrid(unit0)
	if grid.X == n.grid.X && grid.Y == n.grid.Y { // 已经到达对应的位置准备降落
		n.landEnd = false
	} else { // 未到达位置，先到达位置并摆正角度
		unit0.Execute(&model.UnitCmd{
			Type:  global.UnitCmdTypeTurn,
			Angle: math.Pi,
		})
		unit0.Execute(&model.UnitCmd{ // 一般这个位置肯定能到达，机场上一般没有飞行器提留
			Type: global.UnitCmdTypeMove,
			Grid: n.grid,
		})
	}
}

func (n *flyHomeIdleAction) OnBelow(temp any) {
	unit0 := temp.(*unit.Unit)
	grid := unit.GetGrid(unit0)
	tool.UnitManager.RemoveFromGrid(grid, unit0) // 其他行为不占用格子，尝试释放自己占用的空间
}

func (n *flyHomeIdleAction) UpdateAction(temp any) {
	if n.landEnd {
		return
	}
	unit0 := temp.(*unit.Unit)
	unit0.Pos.Z -= n.speed // 默认移动速度是负数(下降)
	if unit0.Pos.Z < n.height {
		unit0.Pos.Z = n.height
		n.landEnd = true
	}
}
