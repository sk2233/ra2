/*
@author: sk
@date: 2023/9/29
*/
package soviet

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/utils"
)

func LoadConstructUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:       "苏联基地",
		Hp:         1000,
		Vision:     9,
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
				"res/image/soviet/construct/苏联基地.png": {
					Count:     26,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(139, 105),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 15}, // 是转变过来的单位 不是建造动画(转变动画)
						global.AnimIdle:      {Start: 15, End: 16},
						global.AnimBuild:     {Start: 16, End: 26},
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
			global.UnitActionTypeNoneIdle:          {},
			global.UnitActionTypeConstructComplete: {},
			global.UnitActionTypeConstruct:         {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "磁能反应炉",
		Hp:         750,
		Vision:     5,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:         2,
			HNum:         3,
			Height:       2,
			ProductPower: 150,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/磁能反应炉.png": {
					Count:     16,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(101, 67),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 10},
						global.AnimIdle:      {Start: 10, End: 16},
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
		Audios: map[string][]*model.AudioRate{
			global.AudioDeath: {
				{Path: "res/audio/effect/bpowdiea.wav", Rate: 1},
				{Path: "res/audio/effect/bpowdieb.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "苏联兵营",
		Hp:         500,
		Vision:     5,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      2,
			HNum:      2,
			Height:    4,
			CostPower: 10,
			Tag:       global.ConstructTagProduce,
			Door:      model.NewGrid(1, 1, 0),
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联兵营.png": {
					Count:     9,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(127, 111),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 9},
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
			global.UnitActionTypeConstruct:      {}, // 建造完到idle
			global.UnitActionTypeRallyPointIdle: {},
			global.UnitActionTypeUnitReady: {
				ProduceUnits: []string{"动员兵", "苏联军犬", "防空步兵", "苏联工程师", "磁暴步兵", "疯狂伊文", "辐射工兵", "鲍里斯"},
			},
			global.UnitActionTypeRallyPoint: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "苏联矿石精炼厂",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      4,
			Height:    3,
			CostPower: 50,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联矿石精炼厂.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(98, 86),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 10},
						global.AnimIdle:      {Start: 10, End: 11}, // TODO 倒矿动画 11 ~ 13
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
		Name:       "苏联战争工厂",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      5,
			Height:    3,
			CostPower: 25,
			Tag:       global.ConstructTagProduce,
			Door:      model.NewGrid(1, 4, 0),
			Dome:      model.NewGrid(1, 2, 0),
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联战争工厂.png": {
					Count:     14,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(138, 106),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 9},
						global.AnimOpenDome:  {Start: 9, End: 14},
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
			global.UnitActionTypeConstruct:      {}, // 建造完到idle
			global.UnitActionTypeRallyPointIdle: {},
			global.UnitActionTypeUnitReady: {
				ProduceUnits: []string{"武装采矿车", "犀牛坦克", "半履带车", "恐怖机器人", "V3导弹车", "磁能坦克", "自爆卡车", "苏联基地车", "武装直升机", "基洛夫飞艇", "天启坦克"},
			},
			global.UnitActionTypeRallyPoint: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "苏联海军船坞",
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
			Tag:       global.ConstructTagProduce,
			Door:      model.NewGrid(3, 3, 0),
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联海军船坞.png": {
					Count:     29,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(201, 110),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 12},
						global.AnimIdle:      {Start: 12, End: 19},
						global.AnimOpenDoor:  {Start: 19, End: 29},
					},
				},
			},
			BeforeDir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联海军船坞_effect.png": {
					Count:     15,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(201, 110),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:     {Start: 0, End: 15},
						global.AnimOpenDoor: {Start: 0, End: 15},
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
			global.UnitActionTypeConstruct:      {}, // 建造完到idle
			global.UnitActionTypeRallyPointIdle: {},
			global.UnitActionTypeUnitReady: {
				ProduceUnits: []string{"野牛运输艇", "台风潜艇", "海蝎", "巨型乌贼", "无畏导弹舰"},
			},
			global.UnitActionTypeRallyPoint: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioRepair: {
				{Path: "res/audio/effect/vifvrepa.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "雷达",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      2,
			HNum:      2,
			Height:    3,
			CostPower: 50,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/雷达.png": {
					Count:     38,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(112, 74),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 15},
						global.AnimIdle:      {Start: 15, End: 38},
						global.AnimDeActive:  {Start: 15, End: 16},
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
			global.UnitActionTypeConstruct:   {}, // 建造完到idle
			global.UnitActionTypeNoneIdle:    {},
			global.UnitActionTypePowerChange: {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "苏联维修工厂",
		Hp:         1200,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      4,
			Height:    3,
			CostPower: 25,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联维修工厂.png": {
					Count:     23,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(98, 90),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 9},
						global.AnimIdle:      {Start: 9, End: 10}, // TODO 维修动画
					},
				},
			},
			BeforeDir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联维修工厂_effect.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(98, 90),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
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
		Audios: map[string][]*model.AudioRate{
			global.AudioRepair: {
				{Path: "res/audio/effect/vifvrepa.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "苏联作战实验室",
		Hp:         500,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      3,
			Height:    4,
			CostPower: 100,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/苏联作战实验室.png": {
					Count:     9,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(119, 85),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 9},
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
		Name:       "核子反应炉",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:         4,
			HNum:         4,
			Height:       2,
			ProductPower: 2000,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/核子反应炉.png": {
					Count:     15,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(144, 86),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 9},
						global.AnimIdle:      {Start: 9, End: 15},
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
		Audios: map[string][]*model.AudioRate{
			global.AudioDeath: {
				{Path: "res/audio/effect/snukexpl.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "工业工厂",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:      3,
			HNum:      3,
			Height:    4,
			CostPower: 100,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/工业工厂.png": {
					Count:     14,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(139, 105),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 10},
						global.AnimIdle:      {Start: 10, End: 14},
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
	return datas
}

func LoadDefenseUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:       "苏联围墙",
		Hp:         300,
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
				global.AnimIdle:       "res/image/soviet/wall/苏墙墩.png",
				global.AnimHorizontal: "res/image/soviet/wall/苏墙撇.png",
				global.AnimVertical:   "res/image/soviet/wall/苏墙那.png",
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
		Name:       "哨戒炮",
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
				"res/image/soviet/construct/稍戒炮1.png": {
					Count:     12,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(37, 17),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 12},
					},
				},
			},
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/construct/稍戒炮2.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(35, 25),
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttack: {
				{Path: "bsenatta.wav", Rate: 1},
				{Path: "bsenattb.wav", Rate: 1},
				{Path: "bsenattc.wav", Rate: 1},
				{Path: "bsenattd.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "战斗碉堡",
		Hp:         600,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:     2,
			HNum:     2,
			Height:   2,
			Tag:      global.ConstructTagEnter,
			Capacity: 5,
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/战斗碉堡.png": {
					Count:     10,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(64, 56),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
						global.AnimIdle:      {Start: 8, End: 9},
						global.AnimActive:    {Start: 9, End: 10},
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
			global.UnitActionTypeConstruct:          {}, // 建造完到idle
			global.UnitActionTypePassengerCountIdle: {},
			global.UnitActionTypeUnitEnterOut:       {},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "防空炮",
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
				"res/image/soviet/construct/防空炮1.png": {
					Count:     8,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(60, 44),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 8},
					},
				},
			},
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/construct/防空炮2.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(52, 51),
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttack: {
				{Path: "bflaatta.wav", Rate: 1},
				{Path: "bflaattb.wav", Rate: 1},
				{Path: "bflaattc.wav", Rate: 1},
				{Path: "bflaattd.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "磁暴线圈",
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
				"res/image/soviet/construct/磁暴线圈.png": {
					Count:     32,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(62, 62),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct: {Start: 0, End: 12},
						global.AnimIdle:      {Start: 12, End: 22}, // 攻击特效
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
		Audios: map[string][]*model.AudioRate{
			global.AudioPrepareAttack: {
				{Path: "btespow.wav", Rate: 1},
			},
			global.AudioAttack: {
				{Path: "btesat1a.wav", Rate: 1},
			},
			global.AudioPowerAttack: {
				{Path: "btesat2a.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "铁幕",
		Hp:         750,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   3,
			HNum:   3,
			Height: 2,
			Effect: "无敌",
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/铁幕1.png": {
					Count:     21,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(82, 77),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct:   {Start: 0, End: 9},
						global.AnimIdle:        {Start: 9, End: 15},
						global.AnimIdle2Active: {Start: 15, End: 21},
					},
				},
				"res/image/soviet/construct/铁幕2.png": {
					Count:     11,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(82, 77),
					Frames: map[string]*model.FrameRange{
						global.AnimActive:      {Start: 0, End: 5},
						global.AnimActive2Idle: {Start: 5, End: 11},
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
		Audios: map[string][]*model.AudioRate{
			global.AudioActive: {
				{Path: "res/audio/effect/siroread.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:       "核弹发射井",
		Hp:         1000,
		Type:       global.UnitTypeConstruct,
		StandTile:  global.TileTypeLand,
		StandExtra: &model.UnitStandExtra{},
		MaskType:   global.UnitMaskTypeConstruct,
		ConstructExtra: &model.UnitConstructExtra{
			WNum:   3,
			HNum:   3,
			Height: 3,
			Effect: "核弹",
		},
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/construct/核弹发射井1.png": {
					Count:     40,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(119, 96),
					Frames: map[string]*model.FrameRange{
						global.AnimConstruct:   {Start: 0, End: 25},
						global.AnimIdle:        {Start: 25, End: 26},
						global.AnimIdle2Active: {Start: 26, End: 40},
						global.AnimActive:      {Start: 39, End: 40},
					},
				},
				"res/image/soviet/construct/核弹发射井2.png": {
					Count:     17,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(119, 96),
					Frames: map[string]*model.FrameRange{
						global.AnimActive2Idle: {Start: 1, End: 17},
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
		Audios: map[string][]*model.AudioRate{
			global.AudioActive: {
				{Path: "res/audio/effect/snukread.wav", Rate: 1},
			},
		},
	})
	return datas
}

func LoadSoldierUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:      "动员兵",
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
				"res/image/soviet/soldier/动员兵.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(33, 33),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
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
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeMoveAnim:  {},
			global.UnitActionTypeEnterUnit: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/iconata.wav", Rate: 1},
				{Path: "res/audio/effect/iconatb.wav", Rate: 1},
				{Path: "res/audio/effect/iconatc.wav", Rate: 1},
				{Path: "res/audio/effect/iconatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/iconmoa.wav", Rate: 1},
				{Path: "res/audio/effect/iconmob.wav", Rate: 1},
				{Path: "res/audio/effect/iconmoc.wav", Rate: 1},
				{Path: "res/audio/effect/iconmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/iconsea.wav", Rate: 1},
				{Path: "res/audio/effect/iconseb.wav", Rate: 1},
				{Path: "res/audio/effect/iconsec.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/icondia.wav", Rate: 1},
				{Path: "res/audio/effect/icondib.wav", Rate: 1},
				{Path: "res/audio/effect/icondic.wav", Rate: 1},
				{Path: "res/audio/effect/icondid.wav", Rate: 1},
				{Path: "res/audio/effect/icondie.wav", Rate: 1},
			},
			global.AnimAttack: { // 攻击声音加在武器上好一点
				{Path: "res/audio/effect/iconatta.wav", Rate: 1},
				{Path: "res/audio/effect/iconattb.wav", Rate: 1},
				// 入住攻击 baocatta.wav  baocattb.wav   baocattc.wav
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "苏联军犬",
		Hp:        400,
		Vision:    9,
		Type:      global.UnitTypeSoldier,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		MaskType: global.UnitMaskTypeSoldier,
		AnimType: global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/soviet/soldier/苏联军犬.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
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
			global.UnitActionTypeMoveAnim: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/idogatca.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/idogmova.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/idogsela.wav", Rate: 1},
				{Path: "res/audio/effect/idogselb.wav", Rate: 1},
				{Path: "res/audio/effect/idogselc.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/idogdiea.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "防空步兵",
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
				"res/image/soviet/soldier/防空步兵.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(49, 47),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
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
			global.UnitActionTypeMoveAnim: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/iflaata.wav", Rate: 1},
				{Path: "res/audio/effect/iflaatb.wav", Rate: 1},
				{Path: "res/audio/effect/iflaatc.wav", Rate: 1},
				{Path: "res/audio/effect/iflaatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/iflamoa.wav", Rate: 1},
				{Path: "res/audio/effect/iflamob.wav", Rate: 1},
				{Path: "res/audio/effect/iflamoc.wav", Rate: 1},
				{Path: "res/audio/effect/iflamod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/iflasea.wav", Rate: 1},
				{Path: "res/audio/effect/iflaseb.wav", Rate: 1},
				{Path: "res/audio/effect/iflasec.wav", Rate: 1},
				{Path: "res/audio/effect/iflased.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/ifladia.wav", Rate: 1},
				{Path: "res/audio/effect/ifladib.wav", Rate: 1},
				{Path: "res/audio/effect/ifladic.wav", Rate: 1},
				{Path: "res/audio/effect/ifladid.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "苏联工程师",
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
				"res/image/soviet/soldier/苏联工程师.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
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
			global.UnitActionTypeMoveAnim: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/iensata.wav", Rate: 1},
				{Path: "res/audio/effect/iensatb.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/iensmoa.wav", Rate: 1},
				{Path: "res/audio/effect/iensmob.wav", Rate: 1},
				{Path: "res/audio/effect/iensmoc.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/ienssea.wav", Rate: 1},
				{Path: "res/audio/effect/iensseb.wav", Rate: 1},
				{Path: "res/audio/effect/ienssec.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/iensdia.wav", Rate: 1},
				{Path: "res/audio/effect/iensdib.wav", Rate: 1},
				{Path: "res/audio/effect/iensdic.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "磁暴步兵",
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
				"res/image/soviet/soldier/磁暴步兵.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(29.5, 34),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
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
			global.UnitActionTypeMoveAnim: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/itesata.wav", Rate: 1},
				{Path: "res/audio/effect/itesatb.wav", Rate: 1},
				{Path: "res/audio/effect/itesatc.wav", Rate: 1},
				{Path: "res/audio/effect/itesatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/itesmoa.wav", Rate: 1},
				{Path: "res/audio/effect/itesmob.wav", Rate: 1},
				{Path: "res/audio/effect/itesmoc.wav", Rate: 1},
				{Path: "res/audio/effect/itesmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/itessea.wav", Rate: 1},
				{Path: "res/audio/effect/itesseb.wav", Rate: 1},
				{Path: "res/audio/effect/itessec.wav", Rate: 1},
				{Path: "res/audio/effect/itessed.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/itesdia.wav", Rate: 1},
				{Path: "res/audio/effect/itesdib.wav", Rate: 1},
				{Path: "res/audio/effect/itesdic.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "疯狂伊文",
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
				"res/image/soviet/soldier/疯狂伊文.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
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
			global.UnitActionTypeMoveAnim: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/icraata.wav", Rate: 1},
				{Path: "res/audio/effect/icraatb.wav", Rate: 1},
				{Path: "res/audio/effect/icraatc.wav", Rate: 1},
				{Path: "res/audio/effect/icraatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/icramoa.wav", Rate: 1},
				{Path: "res/audio/effect/icramob.wav", Rate: 1},
				{Path: "res/audio/effect/icramoc.wav", Rate: 1},
				{Path: "res/audio/effect/icramod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/icrasea.wav", Rate: 1},
				{Path: "res/audio/effect/icraseb.wav", Rate: 1},
				{Path: "res/audio/effect/icrasec.wav", Rate: 1},
				{Path: "res/audio/effect/icrased.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/icradia.wav", Rate: 1},
				{Path: "res/audio/effect/icradib.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "辐射工兵",
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
				"res/image/soviet/soldier/辐射工兵1.png": {
					Count:     13,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(32, 37),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13},
					},
				},
			},
			Dir1: map[string]*model.ImageExtra{
				"res/image/soviet/soldier/辐射工兵2.png": {
					Count:     22,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(32, 37),
					Frames: map[string]*model.FrameRange{
						global.AnimDeployIdle: {Start: 12, End: 22}, // 有蹲下的动作  0 ~  12
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
			global.UnitActionTypeMoveAnim:     {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/idesata.wav", Rate: 1},
				{Path: "res/audio/effect/idesatb.wav", Rate: 1},
				{Path: "res/audio/effect/idesatc.wav", Rate: 1},
				{Path: "res/audio/effect/idesatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/idesmoa.wav", Rate: 1},
				{Path: "res/audio/effect/idesmob.wav", Rate: 1},
				{Path: "res/audio/effect/idesmoc.wav", Rate: 1},
				{Path: "res/audio/effect/idesmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/idessea.wav", Rate: 1},
				{Path: "res/audio/effect/idesseb.wav", Rate: 1},
				{Path: "res/audio/effect/idessec.wav", Rate: 1},
				{Path: "res/audio/effect/idesmod.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/idesdia.wav", Rate: 1},
				{Path: "res/audio/effect/idesdib.wav", Rate: 1},
				{Path: "res/audio/effect/idesdic.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "鲍里斯",
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
				"res/image/soviet/soldier/鲍里斯.png": {
					Count:     14,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(30, 30),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 13}, // 13  ~  14 特殊攻击
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
			global.UnitActionTypeNoneIdle:        {},
			global.UnitActionTypeMoveAnim:        {},
			global.UnitActionTypePlayCreateAudio: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/iborata.wav", Rate: 1},
				{Path: "res/audio/effect/iboratb.wav", Rate: 1},
				{Path: "res/audio/effect/iboratc.wav", Rate: 1},
				{Path: "res/audio/effect/iboratd.wav", Rate: 1},
				{Path: "res/audio/effect/iborate.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/ibormoa.wav", Rate: 1},
				{Path: "res/audio/effect/ibormob.wav", Rate: 1},
				{Path: "res/audio/effect/ibormoc.wav", Rate: 1},
				{Path: "res/audio/effect/ibormod.wav", Rate: 1},
				{Path: "res/audio/effect/ibormoe.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/iborsea.wav", Rate: 1},
				{Path: "res/audio/effect/iborseb.wav", Rate: 1},
				{Path: "res/audio/effect/iborsec.wav", Rate: 1},
				{Path: "res/audio/effect/iborsed.wav", Rate: 1},
				{Path: "res/audio/effect/iborsee.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/ibordia.wav", Rate: 1},
				{Path: "res/audio/effect/ibordib.wav", Rate: 1},
				{Path: "res/audio/effect/ibordic.wav", Rate: 1},
				{Path: "res/audio/effect/ibordid.wav", Rate: 1},
				{Path: "res/audio/effect/ibordie.wav", Rate: 1},
			},
			global.AudioCreate: {
				{Path: "res/audio/effect/iborsea.wav", Rate: 1},
			},
		},
	})
	return datas
}

func LoadVehicleUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:      "武装采矿车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/武装采矿车.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/武装采矿车_turret.png": {
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vwarata.wav", Rate: 1},
				{Path: "res/audio/effect/vwaratb.wav", Rate: 1},
				{Path: "res/audio/effect/vwaratc.wav", Rate: 1},
				{Path: "res/audio/effect/vwaratd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vwarmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vwarmob.wav", Rate: 1},
				{Path: "res/audio/effect/vwarmoc.wav", Rate: 1},
				{Path: "res/audio/effect/vwarmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vwarsea.wav", Rate: 1},
				{Path: "res/audio/effect/vwarseb.wav", Rate: 1},
				{Path: "res/audio/effect/vwarsec.wav", Rate: 1},
				{Path: "res/audio/effect/vwarsed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "犀牛坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/犀牛坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/犀牛坦克_turret.png": {
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
			global.UnitActionTypeTurnTurren: {
				UnitName: "飞天犀牛坦克",
			},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vgrsata.wav", Rate: 1},
				{Path: "res/audio/effect/vgrsatb.wav", Rate: 1},
				{Path: "res/audio/effect/vgrsatc.wav", Rate: 1},
				{Path: "res/audio/effect/vgrsatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vgrsmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vgrsmob.wav", Rate: 1},
				{Path: "res/audio/effect/vgrsmoc.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vgrssea.wav", Rate: 1},
				{Path: "res/audio/effect/vgrsseb.wav", Rate: 1},
				{Path: "res/audio/effect/vgrssec.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "半履带车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{
			Tag:      global.VehicleTagEnter,
			Capacity: 5,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/半履带车.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(63, 63),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/半履带车_turret.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(65, 65),
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
			global.UnitActionTypePassengerCountIdle: {},
			global.UnitActionTypeUnitEnterOut:       {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vflaata.wav", Rate: 1},
				{Path: "res/audio/effect/vflaatb.wav", Rate: 1},
				{Path: "res/audio/effect/vflaatc.wav", Rate: 1},
				{Path: "res/audio/effect/vflaatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vflamoa.wav", Rate: 1},
				{Path: "res/audio/effect/vflamob.wav", Rate: 1},
				{Path: "res/audio/effect/vflamoc.wav", Rate: 1},
				{Path: "res/audio/effect/vflamod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vflasea.wav", Rate: 1},
				{Path: "res/audio/effect/vflaseb.wav", Rate: 1},
				{Path: "res/audio/effect/vflasec.wav", Rate: 1},
				{Path: "res/audio/effect/vflased.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "恐怖机器人",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/恐怖机器人.png": {
					Count:     11,
					AnimSpeed: utils.Frame2AnimSpeed(10),
					Anchor:    model.NewVec2(18.5, 18.5),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 1, End: 7},
						global.AnimAttack: {Start: 7, End: 11},
					},
				},
			},
		},
		ShadowType:   global.UnitShadowTypeNone,
		BloodBarType: global.UnitBloodBarTypeVehicle,
		ActionCreators: map[model.UnitActionType]*model.UnitActionCreatorExtra{
			global.UnitActionTypeLandMove: {
				MoveSpeed: utils.ToMoveSpeed(10),
			},
			global.UnitActionTypeTurnNoWait: {
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
			global.UnitActionTypeMoveAnim: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vtermova.wav", Rate: 1},
				{Path: "res/audio/effect/vtermovb.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vtermova.wav", Rate: 1},
				{Path: "res/audio/effect/vtermovb.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vtersela.wav", Rate: 1},
				{Path: "res/audio/effect/vterselb.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/vterdiea.wav", Rate: 1},
				{Path: "res/audio/effect/vterdieb.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "V3导弹车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/V3导弹车.png": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1}, // TODO 还有发射过的样子
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vv3lata.wav", Rate: 1},
				{Path: "res/audio/effect/vv3latb.wav", Rate: 1},
				{Path: "res/audio/effect/vv3latc.wav", Rate: 1},
				{Path: "res/audio/effect/vv3latd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vv3lmob.wav", Rate: 1},
				{Path: "res/audio/effect/vv3lmoc.wav", Rate: 1},
				{Path: "res/audio/effect/vv3lmod.wav", Rate: 1},
				{Path: "res/audio/effect/vv3lmoe.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vv3lsea.wav", Rate: 1},
				{Path: "res/audio/effect/vv3lseb.wav", Rate: 1},
				{Path: "res/audio/effect/vv3lsec.wav", Rate: 1},
				{Path: "res/audio/effect/vv3lsed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "磁能坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/磁能坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/磁能坦克_turret.png": {
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vtesata.wav", Rate: 1},
				{Path: "res/audio/effect/vtesatb.wav", Rate: 1},
				{Path: "res/audio/effect/vtesatd.wav", Rate: 1},
				{Path: "res/audio/effect/vtesatc.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vtesmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vtesmob.wav", Rate: 1},
				{Path: "res/audio/effect/vtesmoc.wav", Rate: 1},
				{Path: "res/audio/effect/vtesmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vtessea.wav", Rate: 1},
				{Path: "res/audio/effect/vtesseb.wav", Rate: 1},
				{Path: "res/audio/effect/vtessec.wav", Rate: 1},
				{Path: "res/audio/effect/vtessed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "自爆卡车",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/自爆卡车.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(60, 60),
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vdemata.wav", Rate: 1},
				{Path: "res/audio/effect/vdematb.wav", Rate: 1},
				{Path: "res/audio/effect/vdematc.wav", Rate: 1},
				{Path: "res/audio/effect/vdematd.wav", Rate: 1},
				{Path: "res/audio/effect/vdematc.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vdemmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vdemmob.wav", Rate: 1},
				{Path: "res/audio/effect/vdemmoc.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vdemsea.wav", Rate: 1},
				{Path: "res/audio/effect/vdemseb.wav", Rate: 1},
				{Path: "res/audio/effect/vdemsec.wav", Rate: 1},
				{Path: "res/audio/effect/vdemsed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "苏联基地车",
		Hp:        400,
		Vision:    5,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/苏联基地车.png": {
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
			global.UnitActionTypeDeployConstruct: {
				ConstructName: "苏联基地",
			},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vmcsmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vmcsmob.wav", Rate: 1},
				{Path: "res/audio/effect/vmcsmoc.wav", Rate: 1},
				{Path: "res/audio/effect/vmcsmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vmcssea.wav", Rate: 1},
				{Path: "res/audio/effect/vmcsseb.wav", Rate: 1},
				{Path: "res/audio/effect/vmcssec.wav", Rate: 1},
				{Path: "res/audio/effect/vmcssed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "武装直升机",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround, // 只是放置时只能放在地面上
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/武装直升机1.png": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(8),
					Anchor:    model.NewVec2(75, 75),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 2},
					},
				},
			}, // 还有部署与炮塔动画
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
			global.UnitActionTypeFlyDisperseIdle: {},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vchoa1a.wav", Rate: 1},
				{Path: "res/audio/effect/vchoa1b.wav", Rate: 1},
				{Path: "res/audio/effect/vchoa1c.wav", Rate: 1},
				{Path: "res/audio/effect/vchoa1d.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vchomoa.wav", Rate: 1},
				{Path: "res/audio/effect/vchomob.wav", Rate: 1},
				{Path: "res/audio/effect/vchomoc.wav", Rate: 1},
				{Path: "res/audio/effect/vchomod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vchosea.wav", Rate: 1},
				{Path: "res/audio/effect/vchoseb.wav", Rate: 1},
				{Path: "res/audio/effect/vchosec.wav", Rate: 1},
				{Path: "res/audio/effect/vchosed.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/vchodia.wav", Rate: 1},
				{Path: "res/audio/effect/vchodib.wav", Rate: 1},
				{Path: "res/audio/effect/vchodic.wav", Rate: 1},
				{Path: "res/audio/effect/vchodid.wav", Rate: 1},
			}, // 还有部署，取消部署的音效，在部署时，攻击指令语音不一样
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "基洛夫飞艇",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround, // 只是放置时只能放在地面上
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/基洛夫空艇.png": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(94, 71),
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
			global.UnitActionTypeFlyDisperseIdle: {},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: true,
			},
			global.UnitActionTypePlayCreateAudio: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vkirata.wav", Rate: 1},
				{Path: "res/audio/effect/vkiratb.wav", Rate: 1},
				{Path: "res/audio/effect/vkiratc.wav", Rate: 1},
				{Path: "res/audio/effect/vkiratd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vkirmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vkirmob.wav", Rate: 1},
				{Path: "res/audio/effect/vkirmoc.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vkirsea.wav", Rate: 1},
				{Path: "res/audio/effect/vkirseb.wav", Rate: 1},
				{Path: "res/audio/effect/vkirsec.wav", Rate: 1},
				{Path: "res/audio/effect/vkirsed.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/vkirdia.wav", Rate: 1},
				{Path: "res/audio/effect/vkirdib.wav", Rate: 1},
				{Path: "res/audio/effect/vkirdic.wav", Rate: 1},
				{Path: "res/audio/effect/vkirdid.wav", Rate: 1},
			},
			global.AudioCreate: {
				{Path: "res/audio/effect/vkirsea.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "天启坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/天启坦克.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(75, 75),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurretDir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/天启坦克_turret.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(75, 75),
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
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vapoata.wav", Rate: 1},
				{Path: "res/audio/effect/vapoatb.wav", Rate: 1},
				{Path: "res/audio/effect/vapoatc.wav", Rate: 1},
				{Path: "res/audio/effect/vapoatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vapomoa.wav", Rate: 1},
				{Path: "res/audio/effect/vapomob.wav", Rate: 1},
				{Path: "res/audio/effect/vapomoc.wav", Rate: 1},
				{Path: "res/audio/effect/vapomod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vaposea.wav", Rate: 1},
				{Path: "res/audio/effect/vaposeb.wav", Rate: 1},
				{Path: "res/audio/effect/vaposec.wav", Rate: 1},
				{Path: "res/audio/effect/vaposed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "野牛运输艇",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround | global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{
			Tag:      global.VehicleTagEnter,
			Capacity: 12,
		},
		MaskType: global.UnitMaskTypeVehicle,
		AnimType: global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/野牛运输艇.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(90, 90),
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
			global.UnitActionTypePassengerCountIdle: {},
			global.UnitActionTypeWaterWave:          {},
			global.UnitActionTypeUnitEnterOut:       {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vhosmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vhosmob.wav", Rate: 1},
				{Path: "res/audio/effect/vhosmod.wav", Rate: 1},
				{Path: "res/audio/effect/vhosmoe.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vhossea.wav", Rate: 1},
				{Path: "res/audio/effect/vhosseb.wav", Rate: 1},
				{Path: "res/audio/effect/vhossec.wav", Rate: 1},
				{Path: "res/audio/effect/vhossed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "台风潜艇",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/台风潜艇.png": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(60, 60),
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
				BindAngle: true,
			},
			global.UnitActionTypeNoneIdle: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vsubata.wav", Rate: 1},
				{Path: "res/audio/effect/vsubatb.wav", Rate: 1},
				{Path: "res/audio/effect/vsubatc.wav", Rate: 1},
				{Path: "res/audio/effect/vsubatd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vsubmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vsubmob.wav", Rate: 1},
				{Path: "res/audio/effect/vsubmoc.wav", Rate: 1},
				{Path: "res/audio/effect/vsubmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vsubsea.wav", Rate: 1},
				{Path: "res/audio/effect/vsubseb.wav", Rate: 1},
				{Path: "res/audio/effect/vsubsec.wav", Rate: 1},
				{Path: "res/audio/effect/vsubsed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "海蝎",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/海蝎.png": {
					Count:     1,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(56, 56),
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
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vscorata.wav", Rate: 1},
				{Path: "res/audio/effect/vscoratb.wav", Rate: 1},
				{Path: "res/audio/effect/vscoratc.wav", Rate: 1},
				{Path: "res/audio/effect/vscoratd.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vscormoa.wav", Rate: 1},
				{Path: "res/audio/effect/vscormob.wav", Rate: 1},
				{Path: "res/audio/effect/vscormoc.wav", Rate: 1},
				{Path: "res/audio/effect/vscormod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vscorsea.wav", Rate: 1},
				{Path: "res/audio/effect/vscorseb.wav", Rate: 1},
				{Path: "res/audio/effect/vscorsec.wav", Rate: 1},
				{Path: "res/audio/effect/vscorsed.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "巨型乌贼",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeDir,
		AnimExtra: &model.UnitAnimExtra{
			Dir8: map[string]*model.ImageExtra{
				"res/image/soviet/vehicle/巨型乌贼.png": {
					Count:     18,
					AnimSpeed: utils.Frame2AnimSpeed(6),
					Anchor:    model.NewVec2(90, 90),
					Frames: map[string]*model.FrameRange{
						global.AnimIdle:   {Start: 0, End: 1},
						global.AnimMove:   {Start: 0, End: 10},
						global.AnimAttack: {Start: 10, End: 18},
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
		Audios: map[string][]*model.AudioRate{
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vsqumova.wav", Rate: 1},
				{Path: "res/audio/effect/vsqumovb.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vsqusela.wav", Rate: 1},
			},
			global.AudioDeath: {
				{Path: "res/audio/effect/vsqudiea.wav", Rate: 1},
			},
		},
	})
	datas = append(datas, &model.UnitData{
		Name:      "无畏导弹舰",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeWater,
		StandExtra: &model.UnitStandExtra{
			Land: true,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurret,
		AnimExtra: &model.UnitAnimExtra{
			Dir32: map[string]*model.ImageExtra{ // 区分有无导弹
				"res/image/soviet/vehicle/无畏导弹舰.png": {
					Count:     2,
					AnimSpeed: utils.Frame2AnimSpeed(0),
					Anchor:    model.NewVec2(112.5, 112.5),
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
			global.UnitActionTypeNoneIdle:  {},
			global.UnitActionTypeWaterWave: {},
		},
		Audios: map[string][]*model.AudioRate{
			global.AudioAttackCmd: {
				{Path: "res/audio/effect/vwasata.wav", Rate: 1},
				{Path: "res/audio/effect/vwasatb.wav", Rate: 1},
				{Path: "res/audio/effect/vwasatc.wav", Rate: 1},
			},
			global.AudioMoveCmd: {
				{Path: "res/audio/effect/vwasmoa.wav", Rate: 1},
				{Path: "res/audio/effect/vwasmob.wav", Rate: 1},
				{Path: "res/audio/effect/vwasmoc.wav", Rate: 1},
				{Path: "res/audio/effect/vwasmod.wav", Rate: 1},
			},
			global.AudioSelectCmd: {
				{Path: "res/audio/effect/vwassea.wav", Rate: 1},
				{Path: "res/audio/effect/vwasseb.wav", Rate: 1},
				{Path: "res/audio/effect/vwassec.wav", Rate: 1},
			},
		},
	})
	return datas
}
