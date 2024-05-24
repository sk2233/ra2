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

type sequenceAnim struct {
	images    []*ebiten.Image
	animSpeed float64
	current   float64
	anchor    model.Vec2
}

func NewSequenceAnim(data *model.EffectData) *sequenceAnim {
	image := tool.ImageLoader.LoadImage(data.Path)
	return &sequenceAnim{
		images:    utils.SplitImage(image, data.Count, 1),
		animSpeed: data.AnimSpeed,
		anchor:    data.Anchor,
	}
}

func (s *sequenceAnim) DrawEffect(screen *ebiten.Image, temp any) {
	effect0 := temp.(*effect.Effect)
	pos := tool.Camera.World2Screen(effect0.Pos)
	utils.DrawImage(screen, s.images[int(s.current)], pos, s.anchor, global.WindowRegion)
}

func (s *sequenceAnim) UpdateAnim() bool {
	s.current += s.animSpeed
	return s.current >= float64(len(s.images))
}
