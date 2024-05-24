/*
@author: sk
@date: 2023/10/1
*/
package action

import (
	"fmt"
	"ra3/global/model"
	"testing"
)

func Test_getAirPath(t *testing.T) {
	path := getAirPath(model.Grid{X: 0, Y: 0}, model.Grid{X: 10, Y: 5})
	fmt.Println(path)
}
