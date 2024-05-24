/*
@author: sk
@date: 2023/10/4
*/
package other

import (
	"ra3/global"
	"ra3/global/model"
	"ra3/utils"
)

func LoadVehicleUnit() []*model.UnitData {
	datas := make([]*model.UnitData, 0)
	datas = append(datas, &model.UnitData{
		Name:      "飞天犀牛坦克",
		Hp:        400,
		Type:      global.UnitTypeVehicle,
		StandTile: global.TileTypeGround,
		StandExtra: &model.UnitStandExtra{
			Land: false,
		},
		VehicleExtra: &model.UnitVehicleExtra{},
		MaskType:     global.UnitMaskTypeVehicle,
		AnimType:     global.UnitAnimTypeTurnTurret,
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
						global.AnimIdle: {Start: 0, End: 1},
					},
				},
			},
			TurnSpeed: utils.ToTurnSpeed(0.1),
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
			global.UnitActionTypeFlyDisperseIdle: {
				NeedLand:  true,
				MoveSpeed: utils.ToMoveSpeed(8),
			},
			global.UnitActionTypeTurnSpeed: {
				TurnSpeed: utils.ToTurnSpeed(1),
				BindAngle: false,
			},
			global.UnitActionTypeTurnTurren: {
				UnitName: "犀牛坦克",
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
	return datas
}
