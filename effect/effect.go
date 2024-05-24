/*
@author: sk
@date: 2023/10/2
*/
package effect

import (
	interfaces "ra3/global/interface"
	"ra3/global/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type Effect struct {
	Pos   model.Vec3
	Anim  interfaces.IEffectAnim
	Timer int
	Die   bool // 有两种结束方式，一种看时间，一种看动画是否播放完毕
}

func (t *Effect) Update() {
	if t.Die {
		return
	}
	t.Timer--
	if t.Anim.UpdateAnim() || t.Timer == 0 {
		t.Die = true
	}
}

func (t *Effect) GetOrder() float64 {
	return t.Pos.Z*10000 + t.Pos.Y + t.Pos.X // X,Y 等权
}

func (t *Effect) Draw(screen *ebiten.Image) {
	if t.Die {
		return
	}
	t.Anim.DrawEffect(screen, t)
}
