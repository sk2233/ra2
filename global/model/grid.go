/*
@author: sk
@date: 2023/9/9
*/
package model

import (
	"fmt"
)

type Grid struct {
	X, Y, Z int
}

func (g Grid) String() string {
	return fmt.Sprintf("X:%d,Y:%d,Z:%d", g.X, g.Y, g.Z)
}

func (g Grid) Add(other Grid) Grid {
	return Grid{X: g.X + other.X, Y: g.Y + other.Y, Z: g.Z + other.Z}
}

// 向下一个单位，方便贴到地面
func (g Grid) Down() Grid {
	return Grid{X: g.X, Y: g.Y, Z: g.Z - 1}
}

// 向上一个单位
func (g Grid) Up() Grid {
	return Grid{X: g.X, Y: g.Y, Z: g.Z + 1}
}

func (g Grid) Sub(other Grid) Grid {
	return Grid{g.X - other.X, g.Y - other.Y, g.Z - other.Z}
}

func (g Grid) Equal(other Grid) bool {
	return g.X == other.X && g.Y == other.Y && g.Z == other.Z
}

func (g Grid) Len2() int {
	return g.X*g.X + g.Y*g.Y + g.Z*g.Z
}

func NewGrid(x int, y int, z int) Grid {
	return Grid{X: x, Y: y, Z: z}
}
