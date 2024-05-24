/*
@author: sk
@date: 2023/9/3
*/
package utils

import (
	"ra3/global/model"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetAxis(min, max ebiten.Key) float64 {
	res := 0.0
	if ebiten.IsKeyPressed(min) {
		res--
	}
	if ebiten.IsKeyPressed(max) {
		res++
	}
	return res
}

func GetCursorPos() model.Vec2 {
	x, y := ebiten.CursorPosition()
	return model.NewVec2(float64(x), float64(y))
}
