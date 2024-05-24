/*
@author: sk
@date: 2023/9/9
*/
package utils

import "math"

func Max[T int | float64](value1, value2 T) T {
	if value1 > value2 {
		return value1
	}
	return value2
}

func Min[T int | float64](value1, value2 T) T {
	if value1 < value2 {
		return value1
	}
	return value2
}

func Abs[T int | float64](value T) T {
	if value > 0 {
		return value
	}
	return -value
}

func Sign[T int | float64](value T) T {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

// 找angle在以unit为单位刻度的圆上，最接近的刻度值   angle 必须在 0~2PI
func SnapAngle(angle, unit float64) int {
	base := int(angle / unit)
	if angle-float64(base)*unit > unit/2 { // 如果差值大于一半就过度到下一个角度
		base++
	}
	return base % int(math.Pi*2/unit)
}

func TrimAngle(angle float64) float64 { // 约束到  [0 2PI)
	for angle < 0 {
		angle += math.Pi * 2
	}
	for angle >= math.Pi*2 {
		angle -= math.Pi * 2
	}
	return angle
}

// 从angle1到angle2最近的转向
func AngleDir(src, tar float64) float64 {
	if tar < src { // 确保 目标大于当前
		tar += math.Pi * 2
	}
	if tar > src+math.Pi {
		return -1
	}
	return 1
}

// 纠正到 0~2PI
func Atan2(offsetY, offsetX float64) float64 {
	return TrimAngle(math.Atan2(offsetY, offsetX))
}
