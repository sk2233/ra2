/*
@author: sk
@date: 2023/9/9
*/
package anim

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"

	"github.com/hajimehoshi/ebiten/v2"
)

type turnTurretAnim struct {
	anim         *dirImages
	turretAnim   *dirImages
	angle, speed float64
}

func (a *turnTurretAnim) TryPlayAnim(animName string) bool {
	return false
}

func (a *turnTurretAnim) SetListener(listener interfaces.IUnitAnimListener) {
}

func (a *turnTurretAnim) PlayAnim(animName string) {
}

func NewTurnTurretAnim(param *model.UnitAnimExtra) *turnTurretAnim {
	dir32Anim := loadAnim(param.Dir32, 32)
	dir32TurretAnim := loadAnim(param.TurretDir32, 32)
	return &turnTurretAnim{anim: dir32Anim[global.AnimIdle],
		turretAnim: dir32TurretAnim[global.AnimIdle], speed: param.TurnSpeed}
}

func (a *turnTurretAnim) DrawUnit(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos)
	// 绘制身子与机翼
	a.anim.DrawAnim(screen, pos, unit0.MoveDir)
	a.turretAnim.DrawAnim(screen, pos, a.angle)
}

func (a *turnTurretAnim) UpdateAnim(unit any) {
	a.angle += a.speed
}
