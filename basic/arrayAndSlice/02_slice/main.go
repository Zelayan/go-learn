package main

import "fmt"

func main() {
	nums := [3]int{}
	nums[0] = 1
	dnums := nums[:]

	fmt.Printf("dnums: %v", dnums)

}