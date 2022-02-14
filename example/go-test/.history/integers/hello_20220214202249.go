package main

import "fmt"

func Hello(name string) string {
	return "Hello," + name
}

func main() {
	fmt.Printf(Hello("world"))
}