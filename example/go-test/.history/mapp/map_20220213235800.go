package mapp

import "fmt"

func MapNil() {
	// 使用字面初始化
	var a = map[int]int{1: 1}

	for k, v := range a {
		fmt.Print(k + v)
	}
}