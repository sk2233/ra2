/*
@author: sk
@date: 2023/10/4
*/
package action

import (
	"math"
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
)

func initDeployConstructAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeDeployConstruct, &deployConstructActionCreator{})
}

type deployConstructActionCreator struct {
}

func (d *deployConstructActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeDeploy {
		return nil
	}
	return &deployConstructAction{construct: param.ConstructName}
}

type deployConstructAction struct {
	construct string
}

func (d *deployConstructAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Execute(&model.UnitCmd{
		Type:     global.UnitCmdTypeTurn,
		TurnMove: true,
		Angle:    math.Pi,
	})
}

func (d *deployConstructAction) Destroy(unit any) {
}

func (d *deployConstructAction) UpdateAction(temp any) {
	data := tool.DataManager.GetUnitData(d.construct)
	unit0 := temp.(*unit.Unit) // 部署时尽量让原对象处于中间
	oldGrid := unit.GetGrid(unit0)
	grid := oldGrid.Sub(model.NewGrid(data.ConstructExtra.WNum/2, data.ConstructExtra.HNum/2, 0))
	tool.UnitManager.RemoveFromGrid(oldGrid, unit0) // 必须先移除否则影响基地展开
	if tool.UnitManager.CreateUnit(d.construct, grid) != nil {
		unit0.Pos = model.Vec3{X: -32, Y: -32} // TODO 暂时采用移除处理
		utils.PlaySe("uplace.wav")
		tool.UnitManager.RemoveSelect(unit0)
		tool.UiManager.UpdateIconItems()
		tool.UiManager.UpdateMiniMap()
		tool.DataManager.UpdatePower()
	} else {
		tool.UnitManager.AddToGrid(oldGrid, unit0) // 展开失败
	}
	unit0.PopAction()
}
