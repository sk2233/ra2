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

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

func initPassengerCountIdleAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypePassengerCountIdle, &passengerCountIdleActionCreator{})
}

type passengerCountIdleActionCreator struct {
}

func (p *passengerCountIdleActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeCreateIdle {
		return nil
	}
	return &passengerCountIdleAction{isConstruct: param.IsConstruct}
}

type passengerCountIdleAction struct {
	isConstruct bool // 建筑与载具应该是不一样的，但是没有建筑的素材
}

var (
	passengerCountIdleRect = &model.Rect{Size: model.NewVec2(6, 6)}
)

func (p *passengerCountIdleAction) DrawAfter(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	if !unit0.Select && !unit0.Hover {
		return
	}
	pos := tool.Camera.World2Screen(unit0.Pos)
	capacity := unit.GetCapacity(unit0)
	load := len(unit0.Extra.Passengers)
	for i := 0; i < capacity; i++ {
		passengerCountIdleRect.Pos = pos
		pos.X += 8
		if i < load {
			utils.FillRect(screen, passengerCountIdleRect, colornames.Blue, global.WindowRegion)
		}
		utils.StrokeRect(screen, passengerCountIdleRect, 1, colornames.Black, global.WindowRegion)
	}
}

func (p *passengerCountIdleAction) UpdateAction(temp any) {
}
