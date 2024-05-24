/*
@author: sk
@date: 2023/9/30
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

func initAirMoveAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeAirMove, &airMoveActionCreator{})
}

type airMoveActionCreator struct {
}

func (l *airMoveActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	unit0 := temp.(*unit.Unit)
	if cmd.Type != global.UnitCmdTypeMove || !unit.CanMove(unit0) {
		return nil
	}
	return &airMoveAction{grid: cmd.Grid, data: param}
}

type airMoveAction struct {
	data    *model.UnitActionCreatorExtra
	grid    model.Grid   // 目标位置
	path    []model.Grid // 路径
	index   int
	speed   model.Vec3
	nextPos model.Vec3 // 具体的下一个点，用于判断是否到达了
}

func (l *airMoveAction) DrawAfter(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	if !unit0.Select {
		return
	}
	srcPos := tool.Camera.World2Screen(unit0.Pos)
	tarPos := tool.Camera.World2Screen(utils.Grid2PosUpCenter(l.grid))
	utils.DrawLine(screen, srcPos, tarPos, 1, colornames.Green, global.WindowRegion)
}

func (l *airMoveAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	l.path = getAirPath(unit.GetGrid(unit0), l.grid)
	if len(l.path) == 0 { // 无需行动，起点与终点重合了
		unit0.PopAction()
		return
	}
	unit0.Execute(&model.UnitCmd{
		Type: global.UnitCmdTypeMoveStart,
	})
	l.index = 0
	l.calculateSpeed(unit0)
}

func getAirPath(src model.Grid, tar model.Grid) []model.Grid { // 先斜线，后直线到达，无需寻路，只看x,y
	res := make([]model.Grid, 0)
	offX := utils.Sign(tar.X - src.X)
	offY := utils.Sign(tar.Y - src.Y)
	for src.X != tar.X || src.Y != tar.Y {
		if src.X != tar.X {
			src.X += offX
		}
		if src.Y != tar.Y {
			src.Y += offY
		}
		res = append(res, src)
	}
	return res
}

func (l *airMoveAction) Destroy(temp any) {
}

func (l *airMoveAction) UpdateAction(temp any) {
	unit0 := temp.(*unit.Unit)
	unit0.Pos = unit0.Pos.Add(l.speed)
	if unit0.Pos.Sub(l.nextPos).Len2() < l.speed.Len2() { // 到达终点
		unit0.Pos = l.nextPos
		l.index++
		l.calculateSpeed(unit0)
	}
}

func (l *airMoveAction) calculateSpeed(unit0 *unit.Unit) {
	if l.index >= len(l.path) { // 寻路结束了
		unit0.PopAction()
		return
	}
	nextGrid := l.path[l.index]
	l.nextPos = utils.Grid2Pos(nextGrid).Add(model.Vec3{X: 16, Y: 16})
	l.nextPos.Z = tool.UnitManager.GetHeight(l.nextPos.X, l.nextPos.Y) + global.FlyHeight
	// 这里速度只看非Z轴的速度，Z轴变化由单位爬升实现 机场的战斗机 不需要这样爬升可以直接起飞
	offset := l.nextPos.Sub(unit0.Pos)
	offset.Z = 0
	l.speed = offset.Unit().Scale(l.data.MoveSpeed)
	// 进行转向操作
	grid := unit.GetGrid(unit0)
	unit0.Execute(&model.UnitCmd{
		Type:     global.UnitCmdTypeTurn,
		Angle:    utils.Atan2(float64(nextGrid.Y-grid.Y), float64(nextGrid.X-grid.X)),
		TurnMove: true,
	})
	unit0.Execute(&model.UnitCmd{
		Type:     global.UnitCmdTypeMoveToNext,
		Grid:     grid,
		NextGrid: nextGrid,
	})
	// 进行目标点Z轴对齐操作 后进先执行，先执行Z轴对齐
	unit0.Execute(&model.UnitCmd{
		Type:   global.UnitCmdTypeLift,
		Height: l.nextPos.Z,
	})
}
