/*
@author: sk
@date: 2023/10/4
*/
package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initTurnTurrenAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeTurnTurren, &turnTurrenActionCreator{})
}

type turnTurrenActionCreator struct {
}

func (t *turnTurrenActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	unit0 := temp.(*unit.Unit) // 必须状态平稳时再进行转换
	if cmd.Type == global.UnitCmdTypeDeploy && unit0.Actions.IsEmpty() {
		grid := unit.GetGrid(unit0)
		if unit0.Extra.OtherUnit == nil {
			data := tool.DataManager.GetUnitData(param.UnitName)
			unit1 := tool.UnitManager.BuildUnit(grid, data).(*unit.Unit)
			tool.UnitManager.AddUnit(unit1)
			unit0.Extra.OtherUnit = unit1
			unit1.Extra.OtherUnit = unit0
		}
		unit1 := unit0.Extra.OtherUnit
		tool.UnitManager.RemoveFromGrid(grid, unit0)
		tool.UnitManager.AddToGrid(grid, unit1)
		unit1.Pos = unit0.Pos
		unit1.MoveDir = unit0.MoveDir
		unit0.Pos = model.NewVec3(-32, -32, 0)
		tool.UnitManager.SetSelectUnit([]*unit.Unit{unit1})
	}
	return nil
}
