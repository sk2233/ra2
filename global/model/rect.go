/*
@author: sk
@date: 2023/9/2
*/
package model

type Rect struct {
	Pos    Vec2
	Anchor Vec2
	Size   Vec2
}

func (r *Rect) GetMin() Vec2 {
	return r.Pos.Sub(r.Anchor)
}

func (r *Rect) GetMax() Vec2 {
	return r.GetMin().Add(r.Size)
}

func (r *Rect) Middle() Vec2 {
	return r.Pos.Add(r.Size.Scale(0.5))
}
