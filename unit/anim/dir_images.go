/*
@author: sk
@date: 2023/9/9
*/
package anim

import (
	"math"
	"ra3/global"
	"ra3/global/model"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type dirImages struct {
	images    [][]*ebiten.Image // [frame][dir]
	dirUnit   float64
	animSpeed float64
	current   float64
	anchor    model.Vec2
}

func (i *dirImages) DrawAnim(screen *ebiten.Image, pos model.Vec2, dir float64) {
	//i.anchor = utils.GetAnchor() // TODO TEST
	//if len(i.images) > 4 {
	//	i.current = 4 // TODO TEST
	//}
	// 注意，因为图片排序使用的角度坐标系与该系统使用的角度坐标系不一致需要进行一个转换
	dirIndex := utils.SnapAngle(utils.TrimAngle(math.Pi*5/4-dir), i.dirUnit)
	utils.DrawImage(screen, i.images[int(i.current)][dirIndex], pos, i.anchor, global.WindowRegion)
}

func (i *dirImages) UpdateAnim() bool {
	i.current += i.animSpeed
	if int(i.current) < len(i.images) {
		return false
	}
	i.current -= float64(len(i.images))
	return true
}

func NewDirImages(images [][]*ebiten.Image, dirUnit float64, animSpeed float64, anchor model.Vec2) *dirImages {
	return &dirImages{images: images, dirUnit: dirUnit, animSpeed: animSpeed, current: 0, anchor: anchor}
}
