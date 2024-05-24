/*
@author: sk
@date: 2023/9/24
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

func initTurnNoWaitAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeTurnNoWait, &turnNoWaitActionCreator{})
}

type turnNoWaitActionCreator struct {
}

func (o *turnNoWaitActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeTurn {
		return nil
	}
	return &turnNoWaitAction{angle: cmd.Angle, bindAngle: param.BindAngle, turnMove: cmd.TurnMove}
}

type turnNoWaitAction struct {
	angle     float64
	bindAngle bool
	turnMove  bool
}

func (o *turnNoWaitAction) UpdateAction(temp any) {
	unit0 := temp.(*unit.Unit)
	if o.bindAngle { // 绑定情况下，两个转向是一致的
		unit0.MoveDir = o.angle
		unit0.AttackDir = o.angle
	} else if o.turnMove { // 非绑定情况下移动转向会带动炮塔
		offset := utils.TrimAngle(o.angle - unit0.MoveDir)
		unit0.MoveDir = o.angle
		unit0.AttackDir = utils.TrimAngle(unit0.AttackDir + offset)
	} else { // 非绑定情况下旋转炮塔仅旋转炮塔
		unit0.AttackDir = o.angle
	}
	unit0.PopAction()
}
