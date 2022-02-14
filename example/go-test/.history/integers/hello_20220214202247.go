package main

import "fmt"

func Hello(name string) string {
	return "Hello,"
}

func main() {
	fmt.Printf(Hello("world"))
}