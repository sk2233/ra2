/*
@author: sk
@date: 2023/9/10
*/
package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initSwitchDeployAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeSwitchDeploy, &switchDeployActionCreator{})
}

type switchDeployActionCreator struct {
}

func (i *switchDeployActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeDeploy {
		unit0 := temp.(*unit.Unit)
		if unit0.Extra.Deploy {
			unit0.Extra.Deploy = false
			unit0.Anim.PlayAnim(global.AnimIdle)
		} else {
			unit0.Extra.Deploy = true
			unit0.Anim.PlayAnim(global.AnimDeployIdle)
		}
	}
	return nil
}
