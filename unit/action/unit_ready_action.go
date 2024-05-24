/*
@author: sk
@date: 2023/10/3
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

func initUnitReadyAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeUnitReady, &unitReadyActionCreator{})
}

type unitReadyActionCreator struct {
}

func (u *unitReadyActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeUnitReady || !utils.ContainSlice(param.ProduceUnits, cmd.UnitName) {
		return nil
	}
	cmd.Abort = true
	unit0 := temp.(*unit.Unit)
	grid := unit.GetGrid(unit0)
	data := tool.DataManager.GetUnitData(cmd.UnitName)
	animName := ""
	if data.StandExtra.Land {
		grid = grid.Add(unit0.Data.ConstructExtra.Door)
		animName = global.AnimOpenDoor
	} else {
		grid = grid.Add(unit0.Data.ConstructExtra.Dome)
		animName = global.AnimOpenDome
	}
	unit1 := tool.UnitManager.BuildUnit(grid, data).(*unit.Unit)
	tool.UnitManager.AddUnit(unit1)
	unit1.Execute(&model.UnitCmd{
		Type: global.UnitCmdTypeMove,
		Grid: unit0.Extra.RallyPoint,
	})
	return &animOnceAction{anim: animName}
}
