/*
@author: sk
@date: 2023/9/30
*/
package unit

import (
	"math/rand"
	"ra3/global"
	"ra3/global/model"
	"ra3/global/tool"
	"ra3/tile"
	"ra3/utils"
)

// CanStand 仅针对载具与步兵对象
func CanStand(unit *Unit, tileType model.TileType, capacity int) bool {
	return (unit.Data.StandTile&tileType) > 0 && capacity >= GetSize(unit)
}

// GetSize 仅针对载具与步兵对象
func GetSize(unit *Unit) int {
	if unit.Data.Type == global.UnitTypeSoldier {
		return 1
	}
	return 3
}

var (
	gridOffset = model.Vec3{Z: -1} // 陷入地下一点，主要考虑到斜坡，虽然物体处于方块上，但是计算时计算到方块内
)

func GetGrid(unit *Unit) model.Grid {
	return utils.Pos2Grid(unit.Pos.Add(gridOffset))
}

func CanMove(unit *Unit) bool {
	return !unit.Extra.Deploy && !unit.Extra.InHyper
}

func IsInWater(unit *Unit) bool {
	grid := GetGrid(unit)
	temp, ok := tool.TileManager.GetTile(grid).(*tile.Tile)
	return ok && temp != nil && temp.Type == global.TileTypeWater
}

// 播放任意一个单位的对应音效
func PlayRandSe(units []*Unit, name string) {
	if len(units) == 0 {
		return
	}
	PlaySe(units[rand.Intn(len(units))], name)
}

func PlaySe(unit *Unit, name string) {
	audios := unit.Data.Audios[name]
	if audios == nil {
		return
	}
	max := 0
	for _, audio := range audios {
		max += audio.Rate
	}
	res := rand.Intn(max)
	for _, audio := range audios {
		if res >= audio.Rate {
			res -= audio.Rate
		} else {
			tool.AudioManager.PlaySe(audio.Path)
			return
		}
	}
}

func GetCapacity(unit *Unit) int {
	if unit.Data.Type == global.UnitTypeConstruct {
		return unit.Data.ConstructExtra.Capacity
	}
	if unit.Data.Type == global.UnitTypeVehicle {
		return unit.Data.VehicleExtra.Capacity
	}
	return 0
}
