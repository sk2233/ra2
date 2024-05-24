/*
@author: sk
@date: 2023/9/29
*/
package data

import (
	"ra3/data/allies"
	"ra3/data/other"
	"ra3/data/soviet"
	"ra3/global/model"
)

func LoadAlliesUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, allies.LoadConstructUnit()...)
	datas = append(datas, allies.LoadDefenseUnit()...)
	datas = append(datas, allies.LoadSoldierUnit()...)
	datas = append(datas, allies.LoadVehicleUnit()...)
	return datas
}

func LoadSovietUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, soviet.LoadConstructUnit()...)
	datas = append(datas, soviet.LoadDefenseUnit()...)
	datas = append(datas, soviet.LoadSoldierUnit()...)
	datas = append(datas, soviet.LoadVehicleUnit()...)
	return datas
}

func LoadOtherUnit() []*model.UnitData {
	return other.LoadVehicleUnit()
}
