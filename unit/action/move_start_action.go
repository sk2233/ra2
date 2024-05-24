/*
@author: sk
@date: 2023/10/2
*/
package action

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
)

func initMoveStartAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeMoveStart, &moveStartActionCreator{})
}

type moveStartActionCreator struct { // 只管普通动画的切换不管是否在水中
}

func (l *moveStartActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	unit0 := temp.(*unit.Unit)
	unit.PlaySe(unit0, global.AudioMove)
	return nil
}
