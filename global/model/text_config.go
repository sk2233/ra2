/*
@author: sk
@date: 2023/9/16
*/
package model

import (
	"image/color"
)

type TextConfig struct {
	Anchor Vec2
	Color  color.Color
	Size   float64
	Region *Rect
}
