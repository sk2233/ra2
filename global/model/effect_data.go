/*
@author: sk
@date: 2023/10/2
*/
package model

type EffectData struct {
	Name      string
	Type      EffectType
	Path      string
	Count     int
	AnimSpeed float64
	Anchor    Vec2
	Timer     int
}
