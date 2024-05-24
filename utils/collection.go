/*
@author: sk
@date: 2023/9/2
*/
package utils

import "ra3/global/collection"

func HasKey[K comparable, V any](data map[K]V, key K) bool {
	_, ok := data[key]
	return ok
}

func MergeMap[K comparable, V any](maps ...map[K]V) map[K]V {
	res := make(map[K]V)
	for _, temp := range maps {
		for k, v := range temp {
			res[k] = v
		}
	}
	return res
}

func RemoveSlice[T comparable](old, sub []T) []T {
	if len(sub) == 0 {
		return old
	}
	set := collection.NewSet[T](sub...)
	res := make([]T, 0)
	for _, t := range old {
		if !set.Has(t) {
			res = append(res, t)
		}
	}
	return res
}

func ReverseSlice[T any](values []T) []T {
	l, r := 0, len(values)-1
	for l < r {
		values[l], values[r] = values[r], values[l]
		l++
		r--
	}
	return values
}

func ContainSlice[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
