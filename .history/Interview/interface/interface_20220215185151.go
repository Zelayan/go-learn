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
func IsNil(i interface{}) bool {

	vi := reflect.ValueOf(i)
	return vi.IsNil()
}