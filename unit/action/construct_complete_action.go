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
)

func initConstructCompleteAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeConstructComplete, &constructCompleteActionCreator{})
}

type constructCompleteActionCreator struct {
}

func (l *constructCompleteActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeConstructComplete {
		return nil
	}
	return &animOnceAction{anim: global.AnimBuild}
}
