package main

import "fmt"

func main() {
	nums := [3]int{}
	nums[0] = 1
	dnums := nums[:]

	fmt.Printf("dnums: %v", dnums)

}

func initSlice() {
	s := make([]int, 10)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	// [0 0 0 0 0 0 0 0 0 0 1 2 3]
}

func newSlice() {

	list := new([]int)
	// 会报错，
	//list = append(list, 1)
	*list = append(*list, 1)
	fmt.Println(list)
}

func nilSlice() {
	list := []int{}
	list = nil
	// 不可以赋值的，编译都不通过
	list = append(list, 1)
	fmt.Println(list)
}
