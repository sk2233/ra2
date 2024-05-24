/*
@author: sk
@date: 2023/9/2
*/
package utils

import (
	"math"
	"ra3/global"
	"ra3/global/model"
)

// 用于装换 地图中对象的位置
func ToWorldPos(x, y float64) (float64, float64) {
	return (y*2 - x + 32) / 2, (y*2 + x - 32) / 2
}

// 最重要的两个装换函数

func ToIsometric(x, y int) (int, int) {
	return (y - x) / 2, x + y
}

func ToCommon(x, y int) (int, int) {
	return y/2 - x, x + (y+1)/2
}

// ToIsometricGrid 笛卡尔坐标系转等距坐标系
func ToIsometricGrid(grid model.Grid) model.Grid {
	return model.NewGrid((grid.Y-grid.X)/2, grid.X+grid.Y, grid.Z)
}

// ToCommonGrid 等距坐标系转笛卡尔坐标系
func ToCommonGrid(grid model.Grid) model.Grid {
	return model.NewGrid(grid.Y/2-grid.X, grid.X+(grid.Y+1)/2, grid.Z)
}

// 尽量对外使用方便理解的笛卡尔坐标系

// Grid2Pos 笛卡尔坐标系下的grid转笛卡尔坐标系下的pos
func Grid2Pos(grid model.Grid) model.Vec3 {
	return model.NewVec3(float64(grid.X)*global.BlockSize, float64(grid.Y)*global.BlockSize, float64(grid.Z)*global.BlockSize)
}

// Grid2PosUpCenter 笛卡尔坐标系下的grid转笛卡尔坐标系下的pos，并增加到中间的偏移量，与到Y轴顶的偏移量
func Grid2PosUpCenter(grid model.Grid) model.Vec3 {
	return model.NewVec3(float64(grid.X)*global.BlockSize+global.BlockSize/2,
		float64(grid.Y)*global.BlockSize+global.BlockSize/2, float64(grid.Z)*global.BlockSize+global.BlockSize)
}

// Pos2Grid 笛卡尔坐标系下的pos转笛卡尔坐标系下的grid
func Pos2Grid(pos model.Vec3) model.Grid {
	return model.NewGrid(GetAdjustIndex(pos.X), GetAdjustIndex(pos.Y), GetAdjustIndex(pos.Z))
}

func GetAdjustIndex(value float64) int { // 防止负数取整问题
	if value >= 0 {
		return int(value / global.BlockSize)
	}
	offset := int(-value/global.BlockSize) + 1
	value += float64(offset * global.BlockSize)
	return int(value/global.BlockSize) - offset
}

func SnapGrid(pos model.Vec3) model.Vec3 {
	xIndex, yIndex, zIndex := GetAdjustIndex(pos.X), GetAdjustIndex(pos.Y), GetAdjustIndex(pos.Z)
	return model.NewVec3(float64(xIndex)*global.BlockSize, float64(yIndex)*global.BlockSize,
		float64(zIndex)*global.BlockSize)
}

func Frame2AnimSpeed(frame int) float64 {
	return float64(frame) / 60
}

// 每秒移动grid个格子，转换为每帧的速度
func ToMoveSpeed(grid float64) float64 {
	return grid * global.BlockSize / 60 / 2 // TODO 默认移速太快了
}

// second 秒转一圈，换算为角速度
func ToTurnSpeed(second float64) float64 {
	return math.Pi * 2 / second / 60
}

// 传入的必须是 等距坐标系下的位置
func MiniMapGrid2Pos(x, y int) (float64, float64) {
	xPos := x * global.MiniMapBlockSize
	if y%2 == 0 {
		xPos -= global.MiniMapBlockSize / 2
	}
	yPos := (y - 1) * global.MiniMapBlockSize / 2 / 2
	return float64(xPos), float64(yPos)
}
