/*
@author: sk
@date: 2023/9/2
*/
package utils

import (
	interfaces "ra3/global/interface"
	"ra3/global/model"
)

func RectCollision(rect1, rect2 interfaces.IRect) bool {
	return rect1.GetMax().Gt(rect2.GetMin()) && rect2.GetMax().Gt(rect1.GetMin())
}

func PointCollision(pos model.Vec2, rect interfaces.IRect) bool {
	return rect.GetMax().Gt(pos) && rect.GetMin().Lt(pos)
}
