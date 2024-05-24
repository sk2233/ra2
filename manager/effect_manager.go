/*
@author: sk
@date: 2023/10/2
*/
package manager

import (
	"fmt"
	"ra3/effect"
	"ra3/effect/anim"
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"
)

func init() {
	tool.EffectManager = &effectManager{effects: make([]*effect.Effect, 0)}
}

type effectManager struct {
	effects []*effect.Effect
}

func (e *effectManager) CreateEffect(name string, pos model.Vec3) any {
	res := e.BuildEffect(name, pos)
	e.effects = append(e.effects, res)
	tool.DrawManager.AddDraw(res)
	return res
}

func (e *effectManager) Update() {
	// TODO TEST
	//if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	//	pos := tool.Camera.Screen2World(utils.GetCursorPos())
	//	if pos == global.InvalidVec3 {
	//		return
	//	}
	//	tool.EffectManager.CreateEffect("水波", pos)
	//}
	if len(e.effects) == 0 {
		return
	}
	removeEffects := make([]interfaces.IOrderDraw, 0)
	lastEffects := make([]*effect.Effect, 0)
	for _, item := range e.effects {
		if item.Die {
			removeEffects = append(removeEffects, item)
		} else {
			item.Update()
			lastEffects = append(lastEffects, item)
		}
	}
	e.effects = lastEffects
	tool.DrawManager.RemoveDraw(removeEffects...)
}

func (e *effectManager) BuildEffect(name string, pos model.Vec3) *effect.Effect {
	data := tool.DataManager.GetEffectData(name)
	return &effect.Effect{
		Pos:  pos,
		Anim: e.CreateAnim(data),
		// 如果不是显示指定时间的类型，计时器直接设置为-1
		Timer: utils.If(data.Type == global.EffectTypeTimer, data.Timer, -1),
		Die:   false,
	}
}

func (e *effectManager) CreateAnim(data *model.EffectData) interfaces.IEffectAnim {
	switch data.Type {
	case global.EffectTypeAnim:
		return anim.NewSequenceAnim(data)
	case global.EffectTypeTimer:
		return anim.NewImageAnim(data)
	default:
		panic(fmt.Sprintf("CreateAnim err can't handle Type of %v", data.Type))
	}
}
