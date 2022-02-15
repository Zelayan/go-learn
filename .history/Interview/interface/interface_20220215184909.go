package main

import (
	"fmt"
)

func main() {

	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false


}


func IsNil(i interface{}) bool {
	def
}