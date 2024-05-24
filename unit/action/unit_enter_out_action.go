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

func initUnitEnterOutAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeUnitEnterOut, &unitEnterOutActionCreator{})
}

type unitEnterOutActionCreator struct {
}

func (u *unitEnterOutActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type == global.UnitCmdTypeEnterUnit {
		enterUnit(temp, cmd.Unit)
	}
	if cmd.Type == global.UnitCmdTypeDeploy {
		outUnit(temp)
	}
	return nil
}

func outUnit(temp any) {
	unit0 := temp.(*unit.Unit)
	if len(unit0.Extra.Passengers) == 0 || unit.IsInWater(unit0) {
		return
	}
	grid := unit.GetGrid(unit0)
	for _, passenger := range unit0.Extra.Passengers {
		passenger.Pos = unit0.Pos
		passenger.Execute(&model.UnitCmd{
			Type: global.UnitCmdTypeMove,
			Grid: grid,
		})
	}
	unit0.Extra.Passengers = nil
	utils.PlaySe("gexit1a.wav")
	unit0.Anim.TryPlayAnim(global.AnimIdle)
}

func enterUnit(tar any, src any) {
	tarUnit := tar.(*unit.Unit)
	srcUnit := src.(*unit.Unit)
	tool.UnitManager.RemoveFromGrid(unit.GetGrid(srcUnit), srcUnit) // 先从地图中移除
	srcUnit.Pos = model.Vec3{X: -32, Y: -32}
	tarUnit.Extra.Passengers = append(tarUnit.Extra.Passengers, srcUnit)
	utils.PlaySe("genter1a.wav")
	tarUnit.Anim.TryPlayAnim(global.AnimActive)
	tool.UnitManager.RemoveSelect(srcUnit)
}
