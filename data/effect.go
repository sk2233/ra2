/*
@author: sk
@date: 2023/10/2
*/
package data

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/utils"
)

func LoadEffect() []*model.EffectData {
	datas := make([]*model.EffectData, 0)
	datas = append(datas, &model.EffectData{
		Name:      "水波",
		Type:      global.EffectTypeAnim,
		Path:      "res/image/effect/水波.png",
		Count:     15,
		AnimSpeed: utils.Frame2AnimSpeed(10),
		Anchor:    model.NewVec2(28, 17),
	})
	return datas
}
