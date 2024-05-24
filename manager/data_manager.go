/*
@author: sk
@date: 2023/9/9
*/
package manager

import (
	"fmt"
	"ra3/data"
	"ra3/global"
	"ra3/global/collection"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/unit"
	"ra3/utils"
)

func init() {
	temp := &dataManager{
		money:       10000,
		power:       0,
		load:        0,
		powerEnough: false,
	}
	temp.loadUnitData()
	temp.loadIconData()
	temp.loadEffectData()
	tool.DataManager = temp
}

func initDataManager() {
	tool.DataManager.UpdatePower()
}

type dataManager struct {
	unitDatas   map[string]*model.UnitData
	iconDatas   map[string]*model.IconData
	effectDatas map[string]*model.EffectData
	money       int64 // 钱数目
	power, load int64 // 总电量与负载
	powerEnough bool
}

func (d *dataManager) GetEffectData(name string) *model.EffectData {
	return d.effectDatas[name]
}

func (d *dataManager) UpdatePower() {
	temps := tool.UnitManager.FilterUnits(func(temp any) bool {
		return temp.(*unit.Unit).Data.Type == global.UnitTypeConstruct
	}).([]*unit.Unit)
	d.power = 0
	d.load = 0
	for _, item := range temps {
		d.power += item.Data.ConstructExtra.ProductPower
		d.load += item.Data.ConstructExtra.CostPower
	}
	powerEnough := d.power > d.load
	if powerEnough != d.powerEnough {
		tool.UnitManager.BroadcastCmd(func(temp any) bool {
			unit0 := temp.(*unit.Unit)
			return unit0.Data.Type == global.UnitTypeConstruct
		}, &model.UnitCmd{
			Type:        global.UnitCmdTypePowerChange,
			PowerEnough: powerEnough,
		})
		d.powerEnough = powerEnough
		if !d.powerEnough {
			utils.PlayDeputySe("deputy053.wav")
		}
	}
}

func (d *dataManager) ChangeMoney(value int64) bool {
	if d.money+value < 0 {
		return false
	}
	d.money += value
	return true
}

func (d *dataManager) FilterIconDatas(filter func(icon *model.IconData) bool) []*model.IconData {
	res := make([]*model.IconData, 0)
	for _, temp := range d.iconDatas {
		if filter(temp) {
			res = append(res, temp)
		}
	}
	return res
}

func (d *dataManager) GetIconData(name string) *model.IconData {
	return d.iconDatas[name]
}

func (d *dataManager) GetMoney() int64 {
	return d.money
}

func (d *dataManager) GetPower() (int64, int64) {
	return d.power, d.load
}

func (d *dataManager) GetUnitData(name string) *model.UnitData {
	return d.unitDatas[name]
}

func (d *dataManager) loadUnitData() {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, data.LoadAlliesUnit()...)
	datas = append(datas, data.LoadSovietUnit()...)
	datas = append(datas, data.LoadOtherUnit()...)
	d.unitDatas = make(map[string]*model.UnitData, len(datas))
	for _, temp := range datas {
		d.unitDatas[temp.Name] = temp
	}
}

func (d *dataManager) loadEffectData() {
	datas := make([]*model.EffectData, 0)
	datas = append(datas, data.LoadEffect()...)
	d.effectDatas = make(map[string]*model.EffectData, len(datas))
	for _, temp := range datas {
		d.effectDatas[temp.Name] = temp
	}
}

func (d *dataManager) loadIconData() {
	datas := make([]*model.IconData, 0)
	datas = append(datas, data.LoadAlliesIcon()...)
	datas = append(datas, data.LoadSovietIcon()...)
	d.iconDatas = make(map[string]*model.IconData, len(datas))
	conditions := collection.NewSet[string]() // 被当做条件的建筑
	for index, temp := range datas {
		temp.Order = index                    // 确定排序
		conditions.AddAll(temp.Conditions...) // 收集被当做条件的，只有这些才使用bit位存储
		d.iconDatas[temp.Name] = temp
	}
	index := 0 // 只有条件对象才有hash值，普通对象直接使用默认值0，节约hash位置
	for _, temp := range datas {
		if conditions.Has(temp.Name) {
			temp.Hash = 1 << index
			index++
		}
	}
	for _, item := range datas {
		res := 0
		for _, name := range item.Conditions {
			if temp, ok := d.iconDatas[name]; ok {
				res |= temp.Hash
			} else {
				panic(fmt.Sprintf("data : %+v condition : %v invalid", temp, name))
			}
		}
		item.ConditionHash = res // 生成对应的条件hash
	}
}
