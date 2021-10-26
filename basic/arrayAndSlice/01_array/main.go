package main

import "fmt"

func main() {
	nums := [3]int{}
	nums[0] = 1

	n := nums[0]
	n = 2

	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("n: %d\n", n)
}
