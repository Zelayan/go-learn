package main

import "fmt"

func main() {
	a := make([]int, 5, 5)
	
	b := make([]int, 0, 5)
	b = append(b, []int{1,2,3}...)

	copy(a, b)
	for _, v := range b {
		fmt.Print(v)
	}
	fmt.Println("x")
	for _, v := range a {
		fmt.Print(v)
	}
}
