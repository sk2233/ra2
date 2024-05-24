/*
@author: sk
@date: 2023/9/2
*/
package main

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/manager"
	"ra3/other"
	"ra3/ui"
	"ra3/unit/action"
	"ra3/unit/blood_bar"
	"ra3/unit/shadow"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(global.WindowW, global.WindowH)
	err := ebiten.RunGame(NewMainGame())
	utils.HandleErr(err)
}

type MainGame struct {
	test *Test
}

func NewMainGame() *MainGame {
	// Init() 方法里的对象都是必须依赖其他对象的初始化，无依赖的直接在 对应包内的 init() 里初始化
	other.Init()
	action.Init()
	blood_bar.Init()
	shadow.Init()
	manager.Init()
	utils.Init()
	ui.Init()
	tool.Camera.SetPos(model.NewVec2(2975, 605)) // TODO TEST 默认视角挪到基地
	tool.AudioManager.PlayBgm("res/audio/bg.ogg", 3*60+44)
	return &MainGame{test: NewTest()}
}

func (m *MainGame) Update() error {
	tool.Camera.Update()
	tool.UnitManager.Update()
	tool.UiManager.Update()
	tool.CursorManager.Update()
	tool.EffectManager.Update()
	m.test.Update()
	utils.UpdateAnchor()
	return nil
}

func (m *MainGame) Draw(screen *ebiten.Image) {
	tool.DrawManager.DrawBefore(screen)
	tool.DrawManager.Draw(screen)
	tool.FogManager.Draw(screen)
	tool.DrawManager.DrawAfter(screen)
	tool.UnitManager.DrawAfter(screen)
	tool.CursorManager.DrawAfter(screen)
	tool.UiManager.DrawAfter(screen)
	m.test.Draw(screen)
}

func (m *MainGame) Layout(w, h int) (int, int) {
	return w, h
}
