/*
@author: sk
@date: 2023/9/29
*/
package model

// 因为不能循环引用就直接在这里冗余实现了

func Max[T int | float64](value1, value2 T) T {
	if value1 > value2 {
		return value1
	}
	return value2
}

func Min[T int | float64](value1, value2 T) T {
	if value1 < value2 {
		return value1
	}
	return value2
}
