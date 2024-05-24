/*
@author: sk
@date: 2023/10/2
*/
package anim

import (
	"ra3/effect"
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type imageAnim struct {
	image  *ebiten.Image
	anchor model.Vec2
}

func NewImageAnim(data *model.EffectData) *imageAnim {
	return &imageAnim{
		image:  tool.ImageLoader.LoadImage(data.Path),
		anchor: data.Anchor,
	}
}

func (i *imageAnim) DrawEffect(screen *ebiten.Image, temp any) {
	effect0 := temp.(*effect.Effect)
	pos := tool.Camera.World2Screen(effect0.Pos)
	utils.DrawImage(screen, i.image, pos, i.anchor, global.WindowRegion)
}

func (i *imageAnim) UpdateAnim() bool {
	return false
}
