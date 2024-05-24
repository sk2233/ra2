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
)

func initEnterUnitAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeEnterUnit, &enterUnitActionCreator{})
}

type enterUnitActionCreator struct {
}

func (e *enterUnitActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeHandleUnit {
		return nil
	}
	unit0 := cmd.Unit.(*unit.Unit)
	if unit0.Data.Type == global.UnitTypeConstruct && unit0.Data.ConstructExtra.Tag != global.ConstructTagEnter {
		return nil
	}
	if unit0.Data.Type == global.UnitTypeVehicle && unit0.Data.VehicleExtra.Tag != global.VehicleTagEnter {
		return nil
	}
	if unit0.Data.Type == global.UnitTypeSoldier { // 对面不能是步兵
		return nil
	}
	capacity := unit.GetCapacity(unit0)
	if len(unit0.Extra.Passengers) >= capacity {
		return nil
	}
	return &enterUnitAction{capacity: capacity, target: unit0}
}

type enterUnitAction struct {
	capacity int
	target   *unit.Unit
}

func (e *enterUnitAction) UpdateAction(temp any) {
	unit0 := temp.(*unit.Unit)
	if len(e.target.Extra.Passengers) >= e.capacity {
		unit0.PopAction() // 来晚了
		return
	}
	l2 := e.target.Pos.Sub(unit0.Pos).Len2()
	minL := global.BlockSize * 3.0 / 2.0
	if l2 < minL*minL { // 可以进入了
		e.target.Execute(&model.UnitCmd{
			Type: global.UnitCmdTypeEnterUnit,
			Unit: unit0,
		})
		unit0.PopAction() // 结束了
	} else {
		unit0.Execute(&model.UnitCmd{ // 继续寻路到达
			Type: global.UnitCmdTypeMove,
			Grid: unit.GetGrid(e.target),
		})
	}
}
