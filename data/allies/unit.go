/*
@author: sk
@date: 2023/9/29
*/
package allies

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/utils"
)

func LoadConstructUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:       "盟军基地",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		ConstructExtra: &model.UnitConstructExtra{
			WNum:         4,
			HNum:         4,
			Height:       2,
			CostPower:    0,
			ProductPower: 60,
			Tag:          global.ConstructTagBase,
		},
		MaskType: global.UnitMaskTypeConstruct,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军基地.png": {
					Count:     25,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(138, 104),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 14}, // 是转变过来的单位 不是建造动画(转变动画)
						global.AnimIdle:      {Start: 14, End: 15},
						global.AnimBuild:     {Start: 15, End: 25},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeAnimOnce:  {},
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeConstruct: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "发电厂",
		Hp:         750,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:         2,
			HNum:         2,
			Height:       2,
			CostPower:    0,
			ProductPower: 200,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军发电厂.png": {
					Count:     16,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(70, 52),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 16},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeMiddle,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "盟军兵营",
		Hp:         500,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      2,
			HNum:      3,
			Height:    2,
			CostPower: 10,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军兵营.png": {
					Count:     10,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(123, 70),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 9},
						global.AnimIdle:      {Start: 9, End: 10},
					},
				},
			},
			AfterDir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军兵营_effect.png": {
					Count:     8,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(123, 70),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 8},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeMiddle,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "盟军矿石精炼厂",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      4,
			Height:    2,
			CostPower: 50,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军矿石精炼厂.png": {
					Count:     12,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(102, 93),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 9},
						global.AnimIdle:      {Start: 9, End: 10}, // TODO 倒矿动画 10 ~ 12
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "盟军战争工厂",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      5,
			Height:    2,
			CostPower: 25,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军战争工厂.png": {
					Count:     14,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(128, 106),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 9},
						global.AnimIdle:      {Start: 9, End: 10}, // TODO 生产飞行单位动画
					},
				},
			},
			AfterDir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军战争工厂_effect.png": {
					Count:     8,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(128, 106),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 8},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "空军指挥部",
		Hp:         600,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      5,
			HNum:      3,
			Height:    3,
			CostPower: 50,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/空军指挥部.png": {
					Count:     20,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(146, 125),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 12},
						global.AnimIdle:      {Start: 12, End: 20},
					},
				},
			},
			AfterDir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/空军指挥部_effect.png": {
					Count:     6,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(146, 125),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 6},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "盟军海军船坞",
		Hp:         1500,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      4,
			HNum:      4,
			Height:    4,
			CostPower: 25,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军海军船坞.png": {
					Count:     36,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(213, 131),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 16}, // TODO 生产动画
					},
				},
			},
			BeforeDir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军海军船坞_effect.png": {
					Count:     15,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(111, 93),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 15},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "盟军维修厂",
		Hp:         1200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      3,
			Height:    1,
			CostPower: 25,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/盟军维修厂.png": {
					Count:     24,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(85, 63),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 9}, // TODO 维修动画
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "科技中心",
		Hp:         500,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      2,
			HNum:      3,
			Height:    6,
			CostPower: 100,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/科技中心.png": {
					Count:     9,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(171, 146),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 9},
					},
				},
			},
			AfterDir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/科技中心_effect.png": {
					Count:     8,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(171, 146),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 8},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeMiddle,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "矿石精炼器",
		Hp:         900,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      3,
			Height:    2,
			CostPower: 200,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/矿石精炼器.png": {
					Count:     16,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(75, 79),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 16},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "控制中心",
		Hp:         600,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      2,
			HNum:      2,
			Height:    2,
			CostPower: 100,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/控制中心.png": {
					Count:     20,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(99, 55),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 10},
						global.AnimIdle:      {Start: 10, End: 20},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeMiddle,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	return datas
}

func LoadDefenseUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:       "盟军围墙",
		Hp:         200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		MaskType:   global.UnitMaskTypeConstruct,
		StandExtra: &model.UnitStandExtra{},
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   1, // 围墙的属性只是设置了单个的，多个连续围墙，是连续创建该对象得到的
			HNum:   1,
			Height: 1,
			Tag:    global.ConstructTagExtend, // 自动向四周扩展最多4格 用于单格建筑，主要用于围墙
		},
		AnimType: global.UnitAnimTypeImage,
		AnimExtra: &model.UnitAnimExtra{
			Images: map[string]string{
				global.AnimIdle:       "res/image/allies/wall/围墙墩.png",
				global.AnimHorizontal: "res/image/allies/wall/围墙撇.png",
				global.AnimVertical:   "res/image/allies/wall/围墙那.png",
			},
			Anchor: model.NewVec2(39, 32),
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeSmall,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeWall:     {}, // 根据周围墙体情况进行自适应图片展示
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "机枪碉堡",
		Hp:         400,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   1,
			HNum:   1,
			Height: 1,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/机枪碉堡.png": {
					Count:     8,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(33, 3),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 7},
						global.AnimIdle:      {Start: 7, End: 8},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeSmall,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "爱国者导弹",
		Hp:         900,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		MaskType:   global.UnitMaskTypeConstruct,
		StandExtra: &model.UnitStandExtra{},
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      1,
			HNum:      1,
			Height:    1,
			CostPower: 50,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/爱国者导弹1.png": {
					Count:     8,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(48, 32),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
					},
				},
			},
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/construct/爱国者导弹2.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(52, 49),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeSmall,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "光棱塔",
		Hp:         600,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      1,
			HNum:      1,
			Height:    3,
			CostPower: 75,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/光棱塔.png": {
					Count:     27,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(86, 77),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 18, End: 27}, // 攻击特效
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeSmall,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "巨炮",
		Hp:         200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		MaskType:   global.UnitMaskTypeConstruct,
		StandExtra: &model.UnitStandExtra{},
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   2,
			HNum:   2,
			Height: 2,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/巨炮1.png": {
					Count:     12,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(116, 81),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 12},
					},
				},
			},
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/construct/巨炮2.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(108, 105),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeMiddle,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "卫星中心",
		Hp:         200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   2,
			HNum:   2,
			Height: 2,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/卫星中心.png": {
					Count:     29,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(91, 48),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 13},
						global.AnimIdle:      {Start: 13, End: 29},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeMiddle,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct: {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:  {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "超时空传送仪",
		Hp:         200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   3,
			HNum:   4,
			Height: 2,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/超时空传送仪1.png": {
					Count:     15,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(103, 91),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct:   {Start: 0, End: 8},
						global.AnimIdle:        {Start: 8, End: 9},
						global.AnimIdle2Active: {Start: 9, End: 15},
					},
				},
				"res/image/allies/construct/超时空传送仪2.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(103, 91),
					Frames: map[string]*model.FrameRange{
						global.AnimActive:      {Start: 0, End: 6},
						global.AnimActive2Idle: {Start: 6, End: 13},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct:  {}, // 建造完到idle
			global.UnitActionTypeIdleActive: {},
			global.UnitActionTypeNoneIdle:   {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "天气控制器",
		Hp:         200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   3,
			HNum:   3,
			Height: 2,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/allies/construct/天气控制器1.png": {
					Count:     15,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(77, 73),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct:   {Start: 0, End: 8},
						global.AnimIdle:        {Start: 8, End: 9},
						global.AnimIdle2Active: {Start: 9, End: 15},
					},
				},
				"res/image/allies/construct/天气控制器2.png": {
					Count:     12,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(77, 73),
					Frames: map[string]*model.FrameRange{
						global.AnimActive:      {Start: 0, End: 6},
						global.AnimActive2Idle: {Start: 6, End: 12},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeConstructBloodBar,
		BloodBarType: global.UnitBloodBarTypeConstruct,
		BloodBarExtra: &model.UnitBloodBarExtra{
			BloodBarSize: global.UnitBloodBarSizeBig,
		},
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeConstruct:  {}, // 建造完到idle
			global.UnitActionTypeIdleActive: {},
			global.UnitActionTypeNoneIdle:   {},
		},
	})
	return datas
}

func LoadSoldierUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:      "美国大兵",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/美国大兵1.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
					},
				},
				"res/image/allies/soldier/美国大兵2.png": {
					Count:     7,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimDeployIdle:   {Start: 0, End: 1},
						global.AnimDeployAttack: {Start: 1, End: 7},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(4),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeSwitchDeploy: {},
			global.UnitActionTypeNoneIdle:     {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "重装大兵",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/重装大兵1.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(29, 29),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
					},
				},
				"res/image/allies/soldier/重装大兵2.png": {
					Count:     6,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(39, 39),
					Frames: map[string]*model.FrameRange{
						global.AnimDeployIdle:   {Start: 0, End: 1},
						global.AnimDeployAttack: {Start: 1, End: 6},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(4),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeSwitchDeploy: {},
			global.UnitActionTypeNoneIdle:     {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "盟军工程师",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/盟军工程师.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(29, 29),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
						global.AnimMove: {Start: 1, End: 7}, // TODO 还有一个拆弹动画
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(4),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "盟军军犬",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/盟军军犬.png": {
					Count:     14,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(35, 35),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 14},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(4),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "火箭飞行兵",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround, // 只是放置时只能放在地面上
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/火箭飞行兵.png": {
					Count:     12,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(28, 28),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 6},
						global.AnimAttack: {Start: 6, End: 12},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeImage,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeAirMove: {
				MoveSpeed: utils.ToMoveSpeed(8),
			},
			global.UnitActionTypeAirLift: {
				MoveSpeed: utils.ToMoveSpeed(8),
			},
			global.UnitActionTypeFlyDisperseIdle: {},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "狙击手",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/狙击手.png": {
					Count:     10,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(28, 29),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 10},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(4),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "海豹部队",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround | global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/海豹部队1.png": {
					Count:     19,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13}, // 还有一个使用C4的动画
					},
				},
				"res/image/allies/soldier/海豹部队2.png": { // 水里的动画
					Count:     24,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimWaterIdle:   {Start: 0, End: 6},
						global.AnimWaterMove:   {Start: 6, End: 12},
						global.AnimWaterAttack: {Start: 12, End: 19}, // 还有一个使用C4的动画
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(5),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "间谍",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir, // TODO 间谍需要特别的动画播放器，可以直接引用其他步兵单位的动画
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/间谍.png": {
					Count:     9,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(43, 43),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
						global.AnimMove: {Start: 1, End: 7},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(4),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "超时空军团兵",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/超时空军团兵.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(35, 35),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7}, // 就是传送完毕，等待冷却完毕的样子
						global.AnimAttack: {Start: 7, End: 13},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeHyperMove: {},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "谭雅",
		Hp:        400,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround | global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/soldier/谭雅1.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
					},
				},
				"res/image/allies/soldier/谭雅2.png": { // 水里的动画
					Count:     18,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimWaterIdle:   {Start: 0, End: 6},
						global.AnimWaterMove:   {Start: 6, End: 12},
						global.AnimWaterAttack: {Start: 12, End: 18},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeSoldier,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(5),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
			global.UnitActionTypeMoveAnim: {
				WaterAnim: true,
			},
		},
	})
	return datas
}

func LoadVehicleUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:      "盟军矿车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/盟军矿车.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(55, 55),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "灰熊坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/灰熊坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/灰熊坦克_turret.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "步兵战车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/步兵战车.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/步兵战车_turret1.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretIdle: {Start: 0, End: 1},
					},
				},
				"res/image/allies/vehicle/步兵战车_turret2.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretGun: {Start: 0, End: 1},
					},
				},
				"res/image/allies/vehicle/步兵战车_turret3.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretRepair: {Start: 0, End: 1},
					},
				},
				"res/image/allies/vehicle/步兵战车_turret4.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretOther: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "遥控坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround | global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/遥控坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(40, 40),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/遥控坦克_turret.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(40, 40),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeImage,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "坦克杀手",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/坦克杀手.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(70, 70),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "海马",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/海马.png": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(105, 105),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 2},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeImage,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeAirMove: {
				MoveSpeed: utils.ToMoveSpeed(8),
			},
			global.UnitActionTypeAirLift: {
				MoveSpeed: utils.ToMoveSpeed(8),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeFlyDisperseIdle: {
				NeedLand:  true,
				MoveSpeed: utils.ToMoveSpeed(8),
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "两栖运输船",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround | global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/两栖运输船.jpg": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 2},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "驱逐舰",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/驱逐舰1.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(75, 75),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			}, // TODO 还要区分是否有舰载机
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "神盾巡洋舰",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/神盾巡洋舰.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(70, 70),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "海豚",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/海豚.png": {
					Count:     12,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(35, 35),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 0, End: 6},
						global.AnimAttack: {Start: 6, End: 12},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnNoWait: {},
			global.UnitActionTypeNoneIdle:   {},
			global.UnitActionTypeMoveAnim:   {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "盟军基地车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/盟军基地车.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "入侵者战机",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/入侵者战机.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeImage,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeAirMove: {
				MoveSpeed: utils.ToMoveSpeed(14),
			},
			global.UnitActionTypeAirLift: {
				MoveSpeed: utils.ToMoveSpeed(14),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeFlyHomeIdle: {
				MoveSpeed: utils.ToMoveSpeed(14),
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "黑鹰战机",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/黑鹰战机.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeImage,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeAirMove: {
				MoveSpeed: utils.ToMoveSpeed(14),
			},
			global.UnitActionTypeAirLift: {
				MoveSpeed: utils.ToMoveSpeed(14),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeFlyHomeIdle: {
				MoveSpeed: utils.ToMoveSpeed(14),
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "光棱坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/光棱坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/光棱坦克_turret.png": {
					Count:     5,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(50, 50),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretIdle:   {Start: 4, End: 5},
						global.AnimTurretAttack: {Start: 0, End: 5},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "幻影坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/幻影坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/幻影坦克_turret.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimTurretIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "战斗要塞",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/战斗要塞.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(70, 70),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "航空母舰",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/allies/vehicle/航空母舰.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(100, 100),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(7),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
	})
	return datas
}
