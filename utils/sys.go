/*
@author: sk
@date: 2023/9/2
*/
package utils

import (
	"encoding/json"
	"reflect"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Clone 深拷贝，传入的对象必须是指针，返回的也是对应的指针
func Clone[T any](src T) T {
	tar := reflect.New(reflect.TypeOf(src).Elem())
	bs, err := json.Marshal(src)
	HandleErr(err)
	HandleErr(json.Unmarshal(bs, tar.Interface()))
	return tar.Interface().(T)
}
