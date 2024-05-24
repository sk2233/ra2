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

func initAirLiftAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeAirLift, &airLiftActionCreator{})
}

type airLiftActionCreator struct {
}

func (l *airLiftActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeLift {
		return nil
	}
	return &airLiftAction{height: cmd.Height, data: param}
}

type airLiftAction struct {
	data   *model.UnitActionCreatorExtra
	height float64
	speed  float64
}

func (l *airLiftAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	if utils.Abs(l.height-unit0.Pos.Z) < l.data.MoveSpeed { // 无需爬升，直接结束
		unit0.Pos.Z = l.height
		unit0.PopAction()
		return
	}
	l.speed = l.data.MoveSpeed * utils.Sign(l.height-unit0.Pos.Z)
}

func (l *airLiftAction) Destroy(unit any) {
}

func (l *airLiftAction) UpdateAction(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Pos.Z += l.speed
	if utils.Abs(l.height-unit0.Pos.Z) < l.data.MoveSpeed {
		unit0.Pos.Z = l.height
		unit0.PopAction()
	}
}
