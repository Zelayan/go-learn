package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint64
}

func Range() {
	u := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	n := make([]*user, 0, len(u))
	for _, v := range u {
		n = append(n, &v)
	}
	// 在for range中，变量v是用来保存迭代切片所得的值，因为v只被声明了一次，每次迭代的值都是赋值给v，该变量的内存地址始终未变，是复制了一份，
	fmt.Println(n)
	for _, v := range n {
		fmt.Println(v)
	}
}

func CloseChan() {

}
