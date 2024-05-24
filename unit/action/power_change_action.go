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
	"ra3/utils"
)

func initPowerChangeAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypePowerChange, &powerChangeActionCreator{})
}

type powerChangeActionCreator struct {
}

func (p *powerChangeActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypePowerChange {
		unit0 := temp.(*unit.Unit)
		unit0.Anim.TryPlayAnim(utils.If(cmd.PowerEnough, global.AnimIdle, global.AnimDeActive))
	}
	return nil
}
