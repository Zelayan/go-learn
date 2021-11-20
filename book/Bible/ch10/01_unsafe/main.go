package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// unsafe.Sizeof函数返回操作数在内存中的字节大小
	fmt.Println(unsafe.Sizeof(true)) // 1
	fmt.Println(unsafe.Sizeof(int32(0))) // 32/8 = 4
	fmt.Println(unsafe.Sizeof(uint32(0))) // 32/8 = 4
	fmt.Println(unsafe.Sizeof(float64(0))) // 64/8 = 8
	fmt.Println(unsafe.Sizeof(complex64(0))) // 64/8 = 8
	fmt.Println(unsafe.Sizeof(&[]int32{})) // 64/8 = 8
	fmt.Println(unsafe.Sizeof("123")) // 16


	fmt.Println("-------------")
	//unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数。和 Sizeof 类似，
	//Alignof 也是返回一个常量表达式，对应一个常量。
	// 通常情况下布尔和数字类型需要对齐到它们本身的大小（最多8个字节），其它的类型对齐到机器字大小。
	fmt.Println(unsafe.Alignof("123")) // 8
	fmt.Println(unsafe.Alignof(true)) // 1
}
