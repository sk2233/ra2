/*
@author: sk
@date: 2023/9/9
*/
package utils

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func initDebug() {
	img := tool.ImageLoader.LoadImage("res/image/grid.png")
	gridImg = SplitImage(img, 2, 2)[1]
}

var (
	gridImg    *ebiten.Image
	gridAnchor = model.Vec2{X: 32}
)

func DrawGrid(screen *ebiten.Image, grid model.Grid, param *model.UnitConstructExtra) {
	for x := 0; x < param.WNum; x++ {
		for y := 0; y < param.HNum; y++ {
			temp := grid.Add(model.NewGrid(x, y, 0))
			pos := tool.Camera.World2Screen(Grid2Pos(temp))
			DrawImage(screen, gridImg, pos, gridAnchor, global.WindowRegion)
		}
	}
}

var (
	anchor = model.Vec2{}
)

func UpdateAnchor() {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		anchor.X++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		anchor.X--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		anchor.Y++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		anchor.Y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		anchor.X += GetAxis(ebiten.KeyRight, ebiten.KeyLeft)
		anchor.Y += GetAxis(ebiten.KeyDown, ebiten.KeyUp)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		LogInfo("anchor:%s", anchor)
	}
}

func GetAnchor() model.Vec2 {
	return anchor
}
