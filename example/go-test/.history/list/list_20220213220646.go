package main

import (
    "fmt"
)

func main() {
    myInts := []int{ 8, 6, 7, 5, 3, 0, 9 }
    fmt.Printf("Ints %v reversed: %v\n", myInts, reverseInts(myInts))
}

func reverseInts(input []int) []int {
    if len(input) == 0 {
        return input
    }
    return append(reverseInts(input[1:]), input[0]) 
}