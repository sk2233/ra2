/*
@author: sk
@date: 2023/10/2
*/
package utils

import (
	"fmt"
	"ra3/global/model"
	"testing"
)

func TestClone(t *testing.T) {
	data := &model.UnitData{}
	temp := Clone(data)
	data.Name = "sdfdsfsdf"
	fmt.Println(temp, data)
}
