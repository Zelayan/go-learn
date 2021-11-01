package main

import "fmt"

func main()  {
	fmt.Println("f1 result: ", f1())
	fmt.Println("f2 result: ", f2())
}

//原因就是return会将返回值先保存起来，对于无名返回值来说，
//保存在一个临时对象中，defer是看不到这个临时对象的；
//而对于有名返回值来说，就保存在已命名的变量中。
func f1() int {
	var i int
	defer func() {
		i++
		fmt.Println("f11: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f12: ", i)
	}()

	i = 1000
	return i // 1000
}

func f2() (i int) {
	defer func() {
		i++
		fmt.Println("f21: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f22: ", i)
	}()

	i = 1000
	return i //1002
}
