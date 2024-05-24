/*
@author: sk
@date: 2023/9/9
*/
package interfaces

import (
	"ra3/global/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type IUnitActionCreator interface {
	CreateAction(unit any, param *model.UnitActionCreatorExtra, cmd *model.UnitCmd) IUnitAction
}

type IUnitAction interface {
	UpdateAction(unit any)
}

type IUnitActionInit interface { //  监听 IUnitAction 的创建
	Init(unit any)
}

type IUnitActionDestroy interface { //  监听 IUnitAction 的销毁
	Destroy(unit any)
}

type IUnitActionOnTop interface { // 监听 IUnitAction 到达顶端
	OnTop(unit any)
}

type IUnitActionOnBelow interface { // 监听 IUnitAction 退出顶端
	OnBelow(unit any)
}

type IUnitAnim interface {
	DrawUnit(screen *ebiten.Image, unit any)
	UpdateAnim(unit any)
	SetListener(listener IUnitAnimListener)
	PlayAnim(animName string)         // 有替补的动画选择，从前向后看，存在一个就使用
	TryPlayAnim(animName string) bool // 尝试播放，会返回是否成功
}

type IUnitAnimListener interface {
	AnimEnd(unit any, animName string)
}

type IUnitShadow interface {
	DrawShadow(screen *ebiten.Image, unit any)
}

type IUnitMask interface {
	PointIn(unit any, pos model.Vec2) bool
	InRect(unit any, rect IRect) bool
}

type IUnitBloodBar interface {
	DrawBloodBar(screen *ebiten.Image, unit any)
}

type IUnitAfterDraw interface {
	DrawAfter(screen *ebiten.Image, unit any)
}
