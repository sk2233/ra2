/*
@author: sk
@date: 2023/9/2
*/
package manager

import (
	interfaces "ra3/global/interface"
	"ra3/global/tool"
	"ra3/utils"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	tool.DrawManager = &drawManager{draws: make([]interfaces.IOrderDraw, 0)}
}

type drawManager struct {
	draws []interfaces.IOrderDraw
}

func (d *drawManager) RemoveDraw(draws ...interfaces.IOrderDraw) {
	if len(draws) == 0 {
		return
	} // 使用泛型set对类型校验太早，interfaces.IOrderDraw没有实现比较接口，但实现都是指针是可以比较的，只能重复写一遍
	set := make(map[interfaces.IOrderDraw]struct{}, len(draws))
	for _, draw := range draws {
		set[draw] = struct{}{}
	}
	res := make([]interfaces.IOrderDraw, 0, len(d.draws)-len(draws))
	for _, draw := range d.draws {
		if _, ok := set[draw]; !ok {
			res = append(res, draw)
		}
	}
	d.draws = res
}

func (d *drawManager) AddDraw(draw interfaces.IOrderDraw) {
	d.draws = append(d.draws, draw)
}

func (d *drawManager) DrawBefore(screen *ebiten.Image) {
	sort.Slice(d.draws, func(i, j int) bool {
		return d.draws[i].GetOrder() < d.draws[j].GetOrder()
	})
	for _, draw := range d.draws {
		utils.InvokeDrawBefore(draw, screen)
	}
}

func (d *drawManager) Draw(screen *ebiten.Image) {
	for _, draw := range d.draws {
		draw.Draw(screen)
	}
}

func (d *drawManager) DrawAfter(screen *ebiten.Image) {
	for _, draw := range d.draws {
		utils.InvokeDrawAfter(draw, screen)
	}
}
