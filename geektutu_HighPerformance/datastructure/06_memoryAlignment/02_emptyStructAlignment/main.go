package main

import (
	"fmt"
	"unsafe"
)

type demo3 struct {
	c int32
	a struct{} // 需要保证安全 填充内存
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {
	fmt.Println(unsafe.Sizeof(demo3{})) // 8
	fmt.Println(unsafe.Sizeof(demo4{})) // 4
}