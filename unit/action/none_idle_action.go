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
)

func initNoneIdleAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeNoneIdle, &noneIdleActionCreator{})
}

type noneIdleActionCreator struct {
}

func (n *noneIdleActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeCreateIdle {
		return nil
	}
	return &noneIdleAction{}
}

type noneIdleAction struct {
}

func (n *noneIdleAction) UpdateAction(unit any) {
}
