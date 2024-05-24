/*
@author: sk
@date: 2023/9/9
*/
package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initConstructAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeConstruct, &constructActionCreator{})
}

type constructActionCreator struct {
}

func (c *constructActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeInit {
		return &constructAction{}
	}
	return nil
}

type constructAction struct {
}

func (c *constructAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Anim.PlayAnim(global.AnimConstruct)
	unit0.Anim.SetListener(c)
}

func (c *constructAction) Destroy(unit any) {
}

func (c *constructAction) AnimEnd(temp any, animName string) {
	if animName == global.AnimConstruct {
		unit0 := temp.(*unit.Unit)
		unit0.Anim.PlayAnim(global.AnimIdle)
		unit0.Anim.SetListener(nil)
		unit0.PopAction()
	}
}

func (c *constructAction) UpdateAction(unit any) {
}
