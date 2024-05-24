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

func initRallyPointIdleAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeRallyPointIdle, &rallyPointIdleActionCreator{})
}

type rallyPointIdleActionCreator struct {
}

func (r *rallyPointIdleActionCreator) CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeCreateIdle {
		return nil
	}
	return &rallyPointIdleAction{}
}

type rallyPointIdleAction struct {
}

func (r *rallyPointIdleAction) DrawAfter(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	if !unit0.Select {
		return
	}
	src := unit.GetGrid(unit0).Add(unit0.Data.ConstructExtra.Door)
	tar := unit0.Extra.RallyPoint
	if src.Equal(tar) {
		return
	}
	srcPos3 := utils.Grid2PosUpCenter(src)
	tarPos3 := utils.Grid2PosUpCenter(tar)
	srcPos2 := tool.Camera.World2Screen(srcPos3)
	tarPos2 := tool.Camera.World2Screen(tarPos3)
	utils.DrawLine(screen, srcPos2, tarPos2, 2, colornames.Orange, global.WindowRegion)
}

func (r *rallyPointIdleAction) UpdateAction(unit any) {
}
