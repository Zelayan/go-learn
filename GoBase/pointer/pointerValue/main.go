package main

import "fmt"

func main() {
	//pointerValue
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // pointerValue（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)



	// 	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
}
