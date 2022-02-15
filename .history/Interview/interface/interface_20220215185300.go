package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false

	fmt.Println("-------")
	fmt.Println(IsNil(a))
	fmt.Println(IsNil(b))

}

// 一个接口包括动态类型和动态值。
// ru
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}	

	return false
}
