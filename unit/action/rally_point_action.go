package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initRallyPointAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeRallyPoint, &rallyPointActionCreator{})
}

type rallyPointActionCreator struct {
}

func (r *rallyPointActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeRallyPoint {
		unit0 := temp.(*unit.Unit)
		unit0.Extra.RallyPoint = cmd.Grid
	}
	return nil
}
