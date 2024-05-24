/*
@author: sk
@date: 2023/9/29
*/
package allies

import (
	"ra3/global"
	"ra3/global/model"
)

func LoadConstructIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "盟军基地",
		Conditions: []string{}, // 设置条件为空 会设置为 0 使用 & 判断永远为0 不会出现在建造栏里
		// 这里只是为了做其他建筑的前置条件，并不允许建造
	})
	datas = append(datas, &model.IconData{
		Name:       "发电厂",
		Conditions: []string{"盟军基地"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/发电厂图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军兵营",
		Conditions: []string{"盟军基地", "发电厂"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/盟军兵营图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军矿石精炼厂",
		Conditions: []string{"盟军基地", "发电厂"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/盟军矿石精炼厂图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军战争工厂",
		Conditions: []string{"盟军基地", "盟军兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/盟军战争工厂图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "空军指挥部",
		Conditions: []string{"盟军基地", "盟军兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/空军指挥部图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军海军船坞",
		Conditions: []string{"盟军基地", "盟军兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/盟军海军船坞图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军维修厂",
		Conditions: []string{"盟军基地", "空军指挥部"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/盟军维修厂图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "科技中心",
		Conditions: []string{"盟军基地", "空军指挥部"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/科技中心图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "矿石精炼器",
		Conditions: []string{"盟军基地", "科技中心"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/矿石精炼器图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "控制中心",
		Conditions: []string{"盟军基地", "科技中心"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupConstruct,
		Path:       "res/image/allies/construct/icon/控制中心图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}

func LoadDefenseIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "盟军围墙",
		Conditions: []string{"盟军基地"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/盟军围墙图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "机枪碉堡",
		Conditions: []string{"盟军基地", "盟军兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/机枪碉堡图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "爱国者导弹",
		Conditions: []string{"盟军基地", "盟军兵营"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/爱国者导弹图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "光棱塔",
		Conditions: []string{"盟军基地", "空军指挥部"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/光棱塔图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "巨炮",
		Conditions: []string{"盟军基地", "空军指挥部"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/巨炮图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "卫星中心",
		Conditions: []string{"盟军基地", "科技中心"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/卫星中心图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "超时空传送仪",
		Conditions: []string{"盟军基地", "科技中心"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/超时空传送仪图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "超时空传送",
		Conditions: []string{"超时空传送仪"},
		Type:       global.IconTypeCountdown,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/超时空传送图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "天气控制器",
		Conditions: []string{"盟军基地", "科技中心"},
		Type:       global.IconTypeSingle,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/天气控制器图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "闪电风暴",
		Conditions: []string{"天气控制器"},
		Type:       global.IconTypeCountdown,
		Group:      global.IconGroupDefense,
		Path:       "res/image/allies/construct/icon/闪电风暴图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}

func LoadSoldierIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "美国大兵",
		Conditions: []string{"盟军兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/美国大兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "重装大兵",
		Conditions: []string{"盟军兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/重装大兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军工程师",
		Conditions: []string{"盟军兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/盟军工程师图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军军犬",
		Conditions: []string{"盟军兵营"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/盟军军犬图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "火箭飞行兵",
		Conditions: []string{"盟军兵营", "空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/火箭飞行兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "狙击手",
		Conditions: []string{"盟军兵营", "空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/狙击手图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "海豹部队",
		Conditions: []string{"盟军兵营", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/海豹部队图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "间谍",
		Conditions: []string{"盟军兵营", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/间谍图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "超时空军团兵",
		Conditions: []string{"盟军兵营", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/超时空军团兵图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "谭雅",
		Conditions: []string{"盟军兵营", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupSoldier,
		Path:       "res/image/allies/soldier/icon/谭雅图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}

func LoadVehicleIcon() []*model.IconData {
	datas := make([]*model.IconData, 0)
	datas = append(datas, &model.IconData{
		Name:       "超时空矿车",
		Conditions: []string{"盟军战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/超时空矿车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "灰熊坦克",
		Conditions: []string{"盟军战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/灰熊坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "步兵战车",
		Conditions: []string{"盟军战争工厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/步兵战车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "遥控坦克",
		Conditions: []string{"盟军战争工厂", "控制中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/遥控坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "坦克杀手",
		Conditions: []string{"盟军战争工厂", "空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/坦克杀手图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "海马运输机",
		Conditions: []string{"盟军战争工厂", "空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/海马运输机图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "两栖运输艇",
		Conditions: []string{"盟军海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/两栖运输艇图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "驱逐舰",
		Conditions: []string{"盟军海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/驱逐舰图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "神盾巡洋舰",
		Conditions: []string{"盟军海军船坞"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/神盾巡洋舰图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "海豚",
		Conditions: []string{"盟军海军船坞", "空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/海豚图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "盟军基地车",
		Conditions: []string{"盟军战争工厂", "盟军维修厂"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/盟军基地车图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "入侵者战机",
		Conditions: []string{"空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/入侵者战机图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "黑鹰战机",
		Conditions: []string{"空军指挥部"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/黑鹰战机图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "光棱坦克",
		Conditions: []string{"盟军战争工厂", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/光棱坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "幻影坦克",
		Conditions: []string{"盟军战争工厂", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/幻影坦克图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "战斗要塞",
		Conditions: []string{"盟军战争工厂", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/战斗要塞图标.png",
		Time:       120, // 帧为单位
	})
	datas = append(datas, &model.IconData{
		Name:       "航空母舰",
		Conditions: []string{"盟军海军船坞", "科技中心"},
		Type:       global.IconTypeQueue,
		Group:      global.IconGroupVehicle,
		Path:       "res/image/allies/vehicle/icon/航空母舰图标.png",
		Time:       120, // 帧为单位
	})
	return datas
}
