/*
@author: sk
@date: 2023/9/29
*/
package soviet

import (
	"ra3/global"
	"ra3/global/model"
)

func LoadConstructIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "苏联基地",
		Conditions: []string{}, // 设置条件为空 会设置为 0 使用 & 判断永远为0 不会出现在建造栏里
		// 这里只是为了做其他建筑的前置条件，并不允许建造
	})
	datas = append(datas, &model.IconData{
		Name:       "磁能反应炉",
		Conditions: []string{"苏联基地"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/磁能反应炉图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联兵营",
		Conditions: []string{"苏联基地", "磁能反应炉"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/苏联兵营图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联矿石精炼厂",
		Conditions: []string{"苏联基地", "磁能反应炉"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/苏联矿石精炼厂.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联战争工厂",
		Conditions: []string{"苏联基地", "苏联兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/苏联战争工厂图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联海军船坞",
		Conditions: []string{"苏联基地", "苏联兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/苏联海军船坞图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "雷达",
		Conditions: []string{"苏联基地", "苏联兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/雷达图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联维修工厂",
		Conditions: []string{"苏联基地", "雷达"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/苏联维修工厂图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联作战实验室",
		Conditions: []string{"苏联基地", "雷达"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/苏联作战实验室图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "核子反应炉",
		Conditions: []string{"苏联基地", "苏联作战实验室"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/核子反应炉图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "工业工厂",
		Conditions: []string{"苏联基地", "苏联作战实验室"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/soviet/construct/icon/工业工厂图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}

func LoadDefenseIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "苏联围墙",
		Conditions: []string{"苏联基地"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/苏联围墙图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "战斗碉堡",
		Conditions: []string{"苏联基地"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/战斗碉堡图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "哨戒炮",
		Conditions: []string{"苏联基地", "苏联兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/哨戒炮图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "防空炮",
		Conditions: []string{"苏联基地", "苏联兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/防空炮图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "磁暴线圈",
		Conditions: []string{"苏联基地", "雷达"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/磁暴线圈图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "铁幕",
		Conditions: []string{"苏联基地", "苏联作战实验室"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/铁幕图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "无敌",
		Conditions: []string{"铁幕"},
		Type:       global.IconTypeCountdown,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/无敌图标.png",
		Time:       1200, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "核弹发射井",
		Conditions: []string{"苏联基地", "苏联作战实验室"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/核弹发射井图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "核弹",
		Conditions: []string{"核弹发射井"},
		Type:       global.IconTypeCountdown,
		Group:      global.IconGroupDefense,
		Path:       "res/image/soviet/construct/icon/核弹图标.png",
		Time:       1200, // 帧为单位
	})
	return datas
}

func LoadSoldierIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "动员兵",
		Conditions: []string{"苏联兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/动员兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联军犬",
		Conditions: []string{"苏联兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/苏联军犬图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "防空步兵",
		Conditions: []string{"苏联兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/防空步兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联工程师",
		Conditions: []string{"苏联兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/苏联工程师图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "磁暴步兵",
		Conditions: []string{"苏联兵营", "雷达"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/磁暴步兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "疯狂伊文",
		Conditions: []string{"苏联兵营", "雷达"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/疯狂伊文图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "辐射工兵",
		Conditions: []string{"苏联兵营", "雷达"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/辐射工兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "鲍里斯",
		Conditions: []string{"苏联兵营", "苏联作战实验室"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/soviet/soldier/icon/鲍里斯图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}

func LoadVehicleIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "武装采矿车",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/武装采矿车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "犀牛坦克",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/犀牛坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "半履带车",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/半履带车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "恐怖机器人",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/恐怖机器人图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "V3导弹车",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/V3导弹车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "磁能坦克",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/磁能坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "自爆卡车",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/自爆卡车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "苏联基地车",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/苏联基地车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "武装直升机",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/武装直升机图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "基洛夫飞艇",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/基洛夫飞艇图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "天启坦克",
		Conditions: []string{"苏联战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/天启坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "野牛运输艇",
		Conditions: []string{"苏联海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/野牛运输艇图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "台风潜艇",
		Conditions: []string{"苏联海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/台风潜艇图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "海蝎",
		Conditions: []string{"苏联海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/海蝎图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "巨型乌贼",
		Conditions: []string{"苏联海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/巨型乌贼图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "无畏导弹舰",
		Conditions: []string{"苏联海军船坞", "苏联作战实验室"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/soviet/vehicle/icon/无畏导弹舰图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}
