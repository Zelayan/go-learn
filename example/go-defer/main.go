package main

import (
	"errors"
	"fmt"
)

func Print() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // 3, 3, 3
		}()
	}
}

func Print2() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func Print3() (r int) {

	// 经过编译实际是
	// r = t
	// t = t+5
	// return

	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func Print4() (r int) {

	// r = 1
	// 值传递
	// func(r int) {
	//		r = r + 5
	//	}(r)

	// return
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f1() {

	// err = nil
	// fmt.Println(err)
	// err = errors.New("defer err")
	var err error
	defer fmt.Println(err)
	err = errors.New("defer err")
	return
}

func f2() {

	// err = nil
	// 这里是闭包，err在执行时变成了defer2 err
	var err error
	defer func() {
		fmt.Println(err)
	}()

	err = errors.New("defer2 err")
	return
}

func f3() {
	var err error
	// 这里的err在nil时就已经被传递进去了
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer3 err")
	return
}

func closure() {
	str := "hello"
	foo := func() {
		str = "xxx"
	}
	foo() //xxx
	fmt.Println(str)
}

func closureRemember() {
	acc := Acc(1)
	print(acc()) //2
	print(acc()) // 3
}

func Acc(value int) func() int {
	return func() int {
		value++
		return value
	}
}
