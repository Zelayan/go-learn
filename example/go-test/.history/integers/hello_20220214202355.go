package main

import "fmt"

func Hello(name string) string {
	x := "Hello,"
	return x + name
}

func main() {
	fmt.Printf(Hello("world"))
}