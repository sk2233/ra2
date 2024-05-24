/*
@author: sk
@date: 2022/10/3
*/
package utils

import (
	"strings"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2/text"
)

// WarpText 或取指定宽度 自动换行后的高度 与对应的字符串
func WarpText(str string, width float64, face font.Face) (float64, string) {
	res := strings.Builder{}
	rs := []rune(str)
	start := 0
	count := 0
	for i := 1; i < len(rs); i++ {
		bound := text.BoundString(face, string(rs[start:i]))
		if bound.Dx() > int(width) {
			res.WriteString(string(rs[start : i-1]))
			res.WriteRune('\n')
			start = i - 1
			count++
		}
	}
	res.WriteString(string(rs[start:]))
	height := text.BoundString(face, str).Dy()
	return float64(height * (count + 1)), res.String()
}
