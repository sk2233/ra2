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
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type turretAnim struct {
	anim           map[string]*dirImages // 绘制身体 一般只使用idle动画，考虑为后面驱逐舰有多身子处理
	turretAnim     map[string]*dirImages // 炮塔动画，切换动画切换的就是这个
	turretAnimName string
	listener       interfaces.IUnitAnimListener // 监听也是监听的炮塔动画事件
}

func (a *turretAnim) TryPlayAnim(animName string) bool {
	if !utils.HasKey(a.turretAnim, animName) {
		return false
	}
	a.PlayAnim(animName)
	return true
}

func (a *turretAnim) SetListener(listener interfaces.IUnitAnimListener) {
	a.listener = listener
}

func (a *turretAnim) PlayAnim(animName string) {
	a.turretAnimName = animName
}

func NewTurretAnim(param *model.UnitAnimExtra) *turretAnim {
	dir32Anim := loadAnim(param.Dir32, 32)
	dir32TurretAnim := loadAnim(param.TurretDir32, 32)
	return &turretAnim{anim: utils.MergeMap(dir32Anim),
		turretAnim:     utils.MergeMap(dir32TurretAnim),
		turretAnimName: global.AnimTurretIdle}
}

func (a *turretAnim) DrawUnit(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos)
	// 绘制身子，身子是固定的
	a.anim[global.AnimIdle].DrawAnim(screen, pos, unit0.MoveDir)
	if anim, ok := a.turretAnim[a.turretAnimName]; ok {
		anim.DrawAnim(screen, pos, unit0.AttackDir)
	}
}

func (a *turretAnim) UpdateAnim(unit any) {
	// 身子没有动画事件
	a.anim[global.AnimIdle].UpdateAnim()
	if anim, ok := a.turretAnim[a.turretAnimName]; ok {
		animEnd := anim.UpdateAnim()
		if animEnd && a.listener != nil {
			a.listener.AnimEnd(unit, a.turretAnimName)
		}
	}
}
