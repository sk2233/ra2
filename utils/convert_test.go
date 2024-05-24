/*
@author: sk
@date: 2023/9/30
*/
package utils

import (
	"fmt"
	"ra3/global/model"
	"testing"
)

func TestToIsometricGridAndToCommonGrid(t *testing.T) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			grid := model.Grid{X: i, Y: j}
			fmt.Println(grid, ToIsometricGrid(ToCommonGrid(grid)))
		}
	}
}

func TestGridOffset(t *testing.T) {
	offsets := [][]int{{-1, -1}, {-1, 1}, {0, -1}, {0, 1}}
	grid := model.Grid{X: -21, Y: 80}
	fmt.Println(grid)
	grid = ToIsometricGrid(grid)
	for _, offset := range offsets {
		temp := grid.Add(model.Grid{X: offset[0], Y: offset[1]})
		fmt.Println(ToCommonGrid(temp))
	}
}
