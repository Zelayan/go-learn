package mapp

import "fmt"

func MapNil() {
	var a = map[int]int{1 1}

	a[1] = 1
	for k, v := range a {
		fmt.Print(k + v)
	}
}