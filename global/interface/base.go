/*
@author: sk
@date: 2023/9/9
*/
package interfaces

import (
	"ra3/global/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type IDraw interface {
	Draw(screen *ebiten.Image)
}

type IBeforeDraw interface {
	DrawBefore(screen *ebiten.Image)
}

type IAfterDraw interface {
	DrawAfter(screen *ebiten.Image)
}

type IOrder interface {
	GetOrder() float64
}

type IUpdate interface {
	Update()
}

type IOrderDraw interface {
	IOrder
	IDraw
}

type IRect interface {
	GetMin() model.Vec2
	GetMax() model.Vec2
}
