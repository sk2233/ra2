/*
@author: sk
@date: 2023/9/30
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

func initHyperMoveAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeHyperMove, &hyperMoveActionCreator{})
}

type hyperMoveActionCreator struct {
}

func (l *hyperMoveActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	unit0 := temp.(*unit.Unit)
	if cmd.Type != global.UnitCmdTypeMove || !unit.CanMove(unit0) {
		return nil
	}
	return &hyperMoveAction{grid: cmd.Grid, data: param}
}

type hyperMoveAction struct {
	data      *model.UnitActionCreatorExtra
	grid      model.Grid // 目标位置
	coolTimer int
}

func (l *hyperMoveAction) Destroy(unit any) {
}

func (l *hyperMoveAction) UpdateAction(temp any) {
	if l.coolTimer > 0 {
		l.coolTimer--
	} else {
		unit0 := temp.(*unit.Unit)
		unit0.Extra.InHyper = false
		unit0.PopAction()
		unit0.Execute(&model.UnitCmd{
			Type: global.UnitCmdTypeMoveEnd,
		})
	}
}

func (l *hyperMoveAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	if !tool.UnitManager.CheckGrid(l.grid, unit0.Data) { // 超时空是最快的，不应该会发生被抢占
		res := tool.TileManager.GetNearGrids(l.grid, []*unit.Unit{unit0}).(map[*unit.Unit]model.Grid)
		if len(res) == 0 {
			unit0.PopAction() // 无法移动直接结束
			return
		}
		l.grid = res[unit0]
	}
	unit0.Execute(&model.UnitCmd{
		Type: global.UnitCmdTypeMoveStart,
	})
	// 先设置到对应的位置
	grid := unit.GetGrid(unit0)
	unit0.Pos = tool.UnitManager.GetPos(l.grid, unit0.Data)
	tool.UnitManager.RemoveFromGrid(grid, unit0)
	tool.UnitManager.AddToGrid(l.grid, unit0)
	// 进行转向
	unit0.Execute(&model.UnitCmd{
		Type:     global.UnitCmdTypeTurn,
		Angle:    utils.Atan2(float64(l.grid.Y-grid.Y), float64(l.grid.X-grid.X)),
		TurnMove: true,
	})
	unit0.Execute(&model.UnitCmd{
		Type:     global.UnitCmdTypeMoveToNext,
		Grid:     grid,
		NextGrid: l.grid,
	})
	unit0.Extra.InHyper = true
	l.coolTimer = int(utils.Grid2Pos(l.grid.Sub(grid)).Len() / 8)
}
