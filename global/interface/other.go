/*
@author: sk
@date: 2023/9/9
*/
package interfaces

import (
	"ra3/global/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type IImageLoader interface {
	LoadImage(path string) *ebiten.Image
}

type ICamera interface {
	World2Screen(pos model.Vec3) model.Vec2
	Screen2WorldAtZ(pos model.Vec2, z float64) model.Vec3
	Update()
	Screen2World(pos model.Vec2) model.Vec3
	SetPos(pos model.Vec2)
	GetPos() model.Vec2
}

type ITextDraw interface {
	DrawText(screen *ebiten.Image, pos model.Vec2, str string, config *model.TextConfig)
}
