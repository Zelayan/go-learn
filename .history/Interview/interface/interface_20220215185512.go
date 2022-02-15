package main

import (
	"fmt"
	"reflect"
)

// var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
// 只有pointer, channel, func, interface, map, or slice 这些类型的值才可以是nil.

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
// 如果一个接口的动态类型和动态值都为空，则这个接口为空的。
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}	

	return false
}
