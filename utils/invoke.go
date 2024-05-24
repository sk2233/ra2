/*
@author: sk
@date: 2023/9/9
*/
package utils

import (
	interfaces "ra3/global/interface"

	"github.com/hajimehoshi/ebiten/v2"
)

func InvokeInit(src any, unit any) {
	if tar, ok := src.(interfaces.IUnitActionInit); ok {
		tar.Init(unit)
	}
}

func InvokeDestroy(src any, unit any) {
	if tar, ok := src.(interfaces.IUnitActionDestroy); ok {
		tar.Destroy(unit)
	}
}

func InvokeOnTop(src any, unit any) {
	if tar, ok := src.(interfaces.IUnitActionOnTop); ok {
		tar.OnTop(unit)
	}
}

func InvokeOnBelow(src any, unit any) {
	if tar, ok := src.(interfaces.IUnitActionOnBelow); ok {
		tar.OnBelow(unit)
	}
}

func InvokeDrawAfter(src any, screen *ebiten.Image) {
	if tar, ok := src.(interfaces.IAfterDraw); ok {
		tar.DrawAfter(screen)
	}
}

func InvokeDrawBefore(src any, screen *ebiten.Image) {
	if tar, ok := src.(interfaces.IBeforeDraw); ok {
		tar.DrawBefore(screen)
	}
}

func InvokeUnitDrawAfter(src any, screen *ebiten.Image, unit any) {
	if tar, ok := src.(interfaces.IUnitAfterDraw); ok {
		tar.DrawAfter(screen, unit)
	}
}
