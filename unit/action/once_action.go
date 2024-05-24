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
)

func initOnceAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeAnimOnce, &animOnceActionCreator{})
}

type animOnceActionCreator struct {
}

func (o *animOnceActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeAnimOnce {
		return nil
	}
	return &animOnceAction{anim: cmd.Anim}
}

type animOnceAction struct {
	anim string
}

func (o *animOnceAction) AnimEnd(temp any, animName string) {
	if animName == o.anim {
		unit0 := temp.(*unit.Unit)
		unit0.PopAction()
	}
}

func (o *animOnceAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	if !unit0.Anim.TryPlayAnim(o.anim) {
		unit0.PopAction()
		return
	}
	unit0.Anim.SetListener(o)
}

func (o *animOnceAction) Destroy(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Anim.PlayAnim(global.AnimIdle)
	unit0.Anim.SetListener(nil)
}

func (o *animOnceAction) UpdateAction(unit any) {
}
