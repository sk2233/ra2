/*
@author: sk
@date: 2023/9/29
*/
package model

type Range struct {
	Pos1, Pos2 Vec2
}

func (r *Range) GetMin() Vec2 {
	return NewVec2(Min(r.Pos1.X, r.Pos2.X), Min(r.Pos1.Y, r.Pos2.Y))
}

func (r *Range) GetMax() Vec2 {
	return NewVec2(Max(r.Pos1.X, r.Pos2.X), Max(r.Pos1.Y, r.Pos2.Y))
}
