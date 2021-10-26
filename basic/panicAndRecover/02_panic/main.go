package main

import "fmt"

func main() {
	defer func() {
	if err := recover(); err != nil {
		fmt.Printf("recover: %v", err)
	}
	}()

	panic("panic")
}