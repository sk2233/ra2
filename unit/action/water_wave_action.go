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

func initWaterWaveAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeWaterWave, &waterWaveActionCreator{})
}

type waterWaveActionCreator struct { // 只管普通动画的切换不管是否在水中
}

func (l *waterWaveActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeMoveToNext {
		unit0 := temp.(*unit.Unit)
		// 只有水面动画(只能在水上走)，或者都有但是当前在水面上
		if param.WaterAnim || unit.IsInWater(unit0) {
			tool.EffectManager.CreateEffect("水波", unit0.Pos)
		}
	}
	return nil
}
