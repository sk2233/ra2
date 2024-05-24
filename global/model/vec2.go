/*
@author: sk
@date: 2023/9/9
*/
package model

import "fmt"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) String() string {
	return fmt.Sprintf("X:%v,Y:%v", v.X, v.Y)
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vec2) Gt(other Vec2) bool {
	return v.X > other.X && v.Y > other.Y
}

func (v Vec2) Lt(other Vec2) bool {
	return v.X < other.X && v.Y < other.Y
}

func (v Vec2) Mul(other Vec2) Vec2 {
	return Vec2{X: v.X * other.X, Y: v.Y * other.Y}
}

func (v Vec2) Scale(val float64) Vec2 {
	return Vec2{X: v.X * val, Y: v.Y * val}
}

func NewVec2(x float64, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}
