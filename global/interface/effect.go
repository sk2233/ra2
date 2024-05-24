/*
@author: sk
@date: 2023/10/2
*/
package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type IEffectAnim interface {
	DrawEffect(screen *ebiten.Image, effect any)
	UpdateAnim() bool
}
