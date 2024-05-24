/*
@author: sk
@date: 2023/9/9
*/
package main

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Test struct {
	img *ebiten.Image
}

var (
	unit0 *unit.Unit
)

func (t *Test) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		pos := tool.Camera.Screen2World(utils.GetCursorPos())
		if pos != global.InvalidVec3 {
			temp := tool.UnitManager.CreateUnit("围墙", utils.Pos2Grid(pos))
			unit0, _ = temp.(*unit.Unit)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) && unit0 != nil {
		unit0.Execute(&model.UnitCmd{Type: global.UnitCmdTypeActive2Idle})
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) && unit0 != nil {
		unit0.Execute(&model.UnitCmd{Type: global.UnitCmdTypeIdle2Active})
	}
}

var (
	testAnchor = model.Vec2{X: 32}
	testText   = &model.TextConfig{
		Anchor: model.NewVec2(0, 0),
		Color:  colornames.Black,
		Size:   2,
	}
	testImg *ebiten.Image
)

func (t *Test) Draw(screen *ebiten.Image) {
	//pos := tool.Camera.Screen2World(utils.GetCursorPos())
	//if pos == global.InvalidVec3 {
	//	return
	//}
	//temp := tool.Camera.World2Screen(utils.SnapGrid(pos))
	//utils.DrawImage(screen, t.img, temp, testAnchor)
	////tool.TextDraw.DrawText(screen, utils.GetCursorPos(), "Text", testText)
	//utils.DrawImage(screen, testImg, utils.GetCursorPos(), model.Vec2{})
}

func NewTest() *Test {
	testImg = tool.ImageLoader.LoadImage("res/image/allies/construct/icon/发电厂图标.png")
	img := tool.ImageLoader.LoadImage("res/image/grid.png")
	imgs := utils.SplitImage(img, 2, 2)
	return &Test{img: imgs[0]}
}
