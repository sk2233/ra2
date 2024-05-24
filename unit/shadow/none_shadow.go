/*
@author: sk
@date: 2023/9/9
*/
package shadow

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type noneShadow struct {
}

func NewNoneShadow() *noneShadow {
	return &noneShadow{}
}

func (s *noneShadow) DrawShadow(screen *ebiten.Image, unit any) {
}
