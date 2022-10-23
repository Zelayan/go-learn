package main

import (
	"errors"
	"fmt"
)

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

// 1. return defer 


func deferFunc() int {
    fmt.Println("defer func called")
    return 0
}

func returnFunc() int {
    fmt.Println("return func called")
    return 0
}

func returnAndDefer() int {
    defer deferFunc()
    return returnFunc()
}


func DeferFunc1(i int) (t int) {
    fmt.Println("t = ", t)
    return 2
}

func returnButDefer() (t int) {  //t初始化0， 并且作用域为该函数全域
    defer func() {
        t = t * 10
    }()

    return 1
}

func defer_call() {
    defer func() { fmt.Println("defer: panic 之前1") }()
    defer func() { fmt.Println("defer: panic 之前2") }()

    panic("异常内容")  //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
}

func defer_call2() {

    defer func() {
        fmt.Println("defer: panic 之前1, 捕获异常")
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

    panic("异常内容")  //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}

func panic2() {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	 }()
 
	 defer func() {
		 panic("defer panic")
	 }()
 
	 panic("panic")
}

func function(index int, value int) int {

    fmt.Println(index)

    return index
}


func deferWithPanic() {
	a, err := panicE()
	if err != nil {
		panic(err)
	}
	defer func ()  {
		fmt.Println("----"+a)	
	}()
}

func panicE() (string, error) {
	return "123", errors.New("xx")
}