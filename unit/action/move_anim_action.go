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
	"ra3/utils"
)

func initMoveAnimAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeMoveAnim, &moveAnimActionCreator{})
}

type moveAnimActionCreator struct { // 只管普通动画的切换不管是否在水中
}

func (l *moveAnimActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	unit0 := temp.(*unit.Unit)
	inWater := param.WaterAnim && unit.IsInWater(unit0)
	switch cmd.Type {
	case global.UnitCmdTypeMoveEnd, global.UnitCmdTypeMovePause:
		unit0.Anim.PlayAnim(utils.If(inWater, global.AnimWaterIdle, global.AnimIdle))
	case global.UnitCmdTypeMoveToNext:
		unit0.Anim.PlayAnim(utils.If(inWater, global.AnimWaterMove, global.AnimMove))
	}
	return nil
}
