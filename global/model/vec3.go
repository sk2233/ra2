/*
@author: sk
@date: 2023/9/9
*/
package model

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) String() string {
	return fmt.Sprintf("X:%v,Y:%v,Z:%v", v.X, v.Y, v.Z)
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v Vec3) Unit() Vec3 {
	l := v.Len()
	return Vec3{X: v.X / l, Y: v.Y / l, Z: v.Z / l}
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.Len2())
}

func (v Vec3) Len2() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Scale(speed float64) Vec3 {
	return Vec3{X: v.X * speed, Y: v.Y * speed, Z: v.Z * speed}
}

func NewVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}
