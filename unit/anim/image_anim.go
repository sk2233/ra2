/*
@author: sk
@date: 2023/9/10
*/
package anim

import (
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type imageAnim struct {
	images   map[string]*ebiten.Image
	animName string
	anchor   model.Vec2
}

func (i *imageAnim) TryPlayAnim(animName string) bool {
	if !utils.HasKey(i.images, animName) {
		return false
	}
	i.PlayAnim(animName)
	return true
}

func NewImageAnim(param *model.UnitAnimExtra) *imageAnim {
	imgs := make(map[string]*ebiten.Image)
	for name, path := range param.Images {
		imgs[name] = tool.ImageLoader.LoadImage(path)
	}
	return &imageAnim{images: imgs, anchor: param.Anchor, animName: global.AnimIdle}
}

func (i *imageAnim) DrawUnit(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos)
	//i.anchor = utils.GetAnchor()
	utils.DrawImage(screen, i.images[i.animName], pos, i.anchor, global.WindowRegion)
}

func (i *imageAnim) UpdateAnim(unit any) {
}

func (i *imageAnim) SetListener(listener interfaces.IUnitAnimListener) {
}

func (i *imageAnim) PlayAnim(animName string) {
	i.animName = animName
}
