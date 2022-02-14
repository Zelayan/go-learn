package main

import "fmt"

func Hello(name string) string {
	return "Hello, world"
}

func main() {
	fmt.Printf(Hello("world"))
}