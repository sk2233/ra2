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

func initIdleActiveAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeIdleActive, &idleActiveActionCreator{})
}

type idleActiveActionCreator struct {
}

func (i *idleActiveActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeIdle2Active {
		unit0 := temp.(*unit.Unit)
		unit.PlaySe(unit0, global.AudioActive)
		return &idleActiveAction{switchAnim: global.AnimIdle2Active, endAnim: global.AnimActive}
	} else if cmd.Type == global.UnitCmdTypeActive2Idle {
		return &idleActiveAction{switchAnim: global.AnimActive2Idle, endAnim: global.AnimIdle}
	}
	return nil
}

type idleActiveAction struct {
	switchAnim string
	endAnim    string
}

func (i *idleActiveAction) AnimEnd(temp any, animName string) {
	if animName == i.switchAnim {
		unit0 := temp.(*unit.Unit)
		unit0.PopAction()
	}
}

func (i *idleActiveAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Anim.PlayAnim(i.switchAnim)
	unit0.Anim.SetListener(i)
}

func (i *idleActiveAction) Destroy(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Anim.PlayAnim(i.endAnim)
	unit0.Anim.SetListener(nil)
}

func (i *idleActiveAction) UpdateAction(unit any) {
}
