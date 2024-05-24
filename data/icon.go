/*
@author: sk
@date: 2023/9/29
*/
package data

import (
	"ra3/data/allies"
	"ra3/data/soviet"
	"ra3/global/model"
)

func LoadAlliesIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, allies.LoadConstructIcon()...)
	datas = append(datas, allies.LoadDefenseIcon()...)
	datas = append(datas, allies.LoadSoldierIcon()...)
	datas = append(datas, allies.LoadVehicleIcon()...)
	return datas
}

func LoadSovietIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, soviet.LoadConstructIcon()...)
	datas = append(datas, soviet.LoadDefenseIcon()...)
	datas = append(datas, soviet.LoadSoldierIcon()...)
	datas = append(datas, soviet.LoadVehicleIcon()...)
	return datas
}
