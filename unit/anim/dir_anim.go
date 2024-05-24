/*
@author: sk
@date: 2023/9/9
*/
package anim

import (
	"math"
	"ra3/global"
	interfaces "ra3/global/interface"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type dirAnim struct {
	anim       map[string]*dirImages
	afterAnim  map[string]*dirImages // 可选动画上不能附加事件
	beforeAnim map[string]*dirImages // 可选动画上不能附加事件
	animName   string
	listener   interfaces.IUnitAnimListener
}

func (a *dirAnim) TryPlayAnim(animName string) bool {
	if !utils.HasKey(a.anim, animName) {
		return false
	}
	a.PlayAnim(animName)
	return true
}

func (a *dirAnim) PlayAnim(animName string) {
	a.animName = animName
}

func (a *dirAnim) SetListener(listener interfaces.IUnitAnimListener) {
	a.listener = listener
}

func NewDirAnim(param *model.UnitAnimExtra) *dirAnim {
	dir1Anims := loadAnim(param.Dir1, 1)
	dir8Anims := loadAnim(param.Dir8, 8)
	dir32Anims := loadAnim(param.Dir32, 32)
	dir1AfterAnims := loadAnim(param.AfterDir1, 1)
	dir1BeforeAnims := loadAnim(param.BeforeDir1, 1)
	return &dirAnim{anim: utils.MergeMap(dir1Anims, dir8Anims, dir32Anims),
		afterAnim:  utils.MergeMap(dir1AfterAnims),
		beforeAnim: utils.MergeMap(dir1BeforeAnims),
		animName:   global.AnimIdle}
}

func loadAnim(param map[string]*model.ImageExtra, dirNum int) map[string]*dirImages {
	res := make(map[string]*dirImages)
	for path, extra := range param {
		img := tool.ImageLoader.LoadImage(path)
		imgs := utils.SplitImage(img, extra.Count, dirNum)
		for name, range0 := range extra.Frames {
			l := range0.End - range0.Start
			temp := make([][]*ebiten.Image, l)
			for i := 0; i < l; i++ {
				temp[i] = make([]*ebiten.Image, dirNum)
				for y := 0; y < dirNum; y++ {
					x := i + range0.Start
					temp[i][y] = imgs[y*extra.Count+x]
				}
			}
			res[name] = NewDirImages(temp, math.Pi*2/float64(dirNum), extra.AnimSpeed, extra.Anchor)
		}
	}
	return res
}

func (a *dirAnim) DrawUnit(screen *ebiten.Image, temp any) {
	unit0 := temp.(*unit.Unit)
	pos := tool.Camera.World2Screen(unit0.Pos)
	//pos = pos.Sub(utils.GetAnchor())
	//if unit0.Data.ConstructExtra != nil { // 只对建筑有效
	//	utils.DrawGrid(screen, utils.Pos2Grid(unit0.Pos), unit0.Data.ConstructExtra) // TODO DEBUG
	//}
	if beforeAnim, ok := a.beforeAnim[a.animName]; ok {
		beforeAnim.DrawAnim(screen, pos, unit0.MoveDir)
	}
	a.anim[a.animName].DrawAnim(screen, pos, unit0.MoveDir)
	if afterAnim, ok := a.afterAnim[a.animName]; ok {
		afterAnim.DrawAnim(screen, pos, unit0.MoveDir)
	}
}

func (a *dirAnim) UpdateAnim(unit any) {
	if beforeAnim, ok := a.beforeAnim[a.animName]; ok {
		beforeAnim.UpdateAnim()
	}
	animEnd := a.anim[a.animName].UpdateAnim()
	if animEnd && a.listener != nil {
		a.listener.AnimEnd(unit, a.animName)
	}
	if afterAnim, ok := a.afterAnim[a.animName]; ok {
		afterAnim.UpdateAnim()
	}
}
