/*
@author: sk
@date: 2023/10/2
*/
package manager

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 0; i++ {
		arr[i] = i
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
}
