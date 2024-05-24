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

func initWallAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeWall, &wallActionCreator{})
}

type wallActionCreator struct {
}

var (
	wallOffsets = []model.Grid{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
)

// 即刻执行，不再创建Action 入栈
func (c *wallActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	if cmd.Type != global.UnitCmdTypeInit && cmd.Type != global.UnitCmdTypeTidy {
		return nil
	}
	unit0 := temp.(*unit.Unit)
	grid := unit.GetGrid(unit0)
	walls := make([]*unit.Unit, 0)
	for _, offset := range wallOffsets {
		wall := getWall(grid.Add(offset), unit0.Data)
		walls = append(walls, wall)
	} // init tidy 都会进行tidy
	if walls[0] != nil && walls[1] != nil && walls[2] == nil && walls[3] == nil {
		unit0.Anim.PlayAnim(global.AnimHorizontal)
	} else if walls[0] == nil && walls[1] == nil && walls[2] != nil && walls[3] != nil {
		unit0.Anim.PlayAnim(global.AnimVertical)
	} else {
		unit0.Anim.PlayAnim(global.AnimIdle)
	}
	if cmd.Type == global.UnitCmdTypeInit { // init 额外触发周围方块的 tidy
		for _, wall := range walls {
			if wall != nil {
				wall.Execute(&model.UnitCmd{Type: global.UnitCmdTypeTidy})
			}
		}
	}
	return nil
}

func getWall(grid model.Grid, data *model.UnitData) *unit.Unit {
	units, ok := tool.UnitManager.GetUnits(grid, true).([]*unit.Unit)
	if !ok || units[3] == nil || units[3].Data != data {
		return nil
	}
	return units[3]
}
