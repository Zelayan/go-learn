package main

import "fmt"

// Definition for singly-linked list.

func maxProfit( prices []int ) int {
	// write code here
	min := prices[0]
	max := 0
	for _, d := range prices {
		min = Min(min, d)
		max = Max(max, d - min)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	ints := []int{4,1, 2, 3,4,119}
	fmt.Println(maxProfit(ints))
}


