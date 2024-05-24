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

func initLandMoveAction() {
	tool.UnitManager.RegisterActionCreator(global.UnitActionTypeLandMove, &landMoveActionCreator{})
}

type landMoveActionCreator struct {
}

func (l *landMoveActionCreator) CreateAction(temp any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) interfaces.IUnitAction {
	unit0 := temp.(*unit.Unit)
	if cmd.Type != global.UnitCmdTypeMove || !unit.CanMove(unit0) {
		return nil
	}
	return &landMoveAction{grid: cmd.Grid, data: param}
}

type landMoveAction struct {
	data    *model.UnitActionCreatorExtra
	grid    model.Grid   // 目标位置
	path    []model.Grid // 路径
	index   int
	speed   model.Vec3
	nextPos model.Vec3 // 具体的下一个点，用于判断是否到达了
}

func (l *landMoveAction) DrawAfter(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	if !unit0.Select {
		return
	}
	srcPos := tool.Camera.World2Screen(unit0.Pos)
	tarPos := tool.Camera.World2Screen(utils.Grid2PosUpCenter(l.grid))
	utils.DrawLine(screen, srcPos, tarPos, 1, colornames.Green, global.WindowRegion)
}

func (l *landMoveAction) Init(temp any) {
	unit0 := temp.(*unit.Unit)
	l.moveToGrid(unit0)
}

func (l *landMoveAction) Destroy(temp any) {
}

func (l *landMoveAction) UpdateAction(temp any) {
	unit0 := temp.(*unit.Unit)
	if !unit0.Extra.Move { // 试图找下一个位置取消等待
		l.calculateSpeed(unit0)
		return
	}
	unit0.Pos = unit0.Pos.Add(l.speed)                    // 正常移动
	if unit0.Pos.Sub(l.nextPos).Len2() < l.speed.Len2() { // 到达终点
		unit0.Pos = l.nextPos
		l.index++
		l.calculateSpeed(unit0)
	}
}

func (l *landMoveAction) moveToGrid(temp *unit.Unit) {
	l.path = tool.TileManager.GetLandPath(l.grid, temp)
	if len(l.path) == 0 { // 寻路失败
		temp.PopAction()
		return
	}
	temp.Execute(&model.UnitCmd{
		Type: global.UnitCmdTypeMoveStart,
	})
	l.index = 0
	l.calculateSpeed(temp)
}

func (l *landMoveAction) calculateSpeed(unit0 *unit.Unit) {
	if l.index >= len(l.path) { // 寻路结束了
		unit0.Extra.Move = false
		unit0.PopAction()
		unit0.Execute(&model.UnitCmd{
			Type: global.UnitCmdTypeMoveEnd,
		})
		return
	}
	nextGrid := l.path[l.index]
	if tool.UnitManager.CheckGrid(nextGrid, unit0.Data) { // 可以通行
		l.nextPos = tool.UnitManager.GetPos(nextGrid, unit0.Data) // 计算位置与速度
		l.speed = l.nextPos.Sub(unit0.Pos).Unit().Scale(l.data.MoveSpeed)
		unit0.Extra.Move = true
		grid := unit.GetGrid(unit0) // 解除老的占用，占用新的格子
		tool.UnitManager.RemoveFromGrid(grid, unit0)
		tool.UnitManager.AddToGrid(nextGrid, unit0)
		unit0.Execute(&model.UnitCmd{ // 先进行转向操作
			Type:     global.UnitCmdTypeTurn,
			Angle:    utils.Atan2(float64(nextGrid.Y-grid.Y), float64(nextGrid.X-grid.X)),
			TurnMove: true,
		})
		unit0.Execute(&model.UnitCmd{
			Type:     global.UnitCmdTypeMoveToNext,
			Grid:     grid,
			NextGrid: nextGrid,
		})
	} else { // 锁定失败
		others := tool.UnitManager.GetUnits(nextGrid, true).([]*unit.Unit)
		for _, other := range others { // 目标位置有已经不移动的对象就重新寻路，简单计算有一个不移动的就重新寻路
			if other != nil && !other.Extra.Move {
				l.moveToGrid(unit0)
				return
			}
		}
		unit0.Extra.Move = false // 都在移动中就等一下
		unit0.Execute(&model.UnitCmd{
			Type: global.UnitCmdTypeMovePause,
		})
	}
}
