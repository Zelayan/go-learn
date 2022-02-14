package main

import "fmt"

const englishHelloPrefix = x := "Hello,"

func Hello(name string) string {
	
	return x + name
}

func main() {
	fmt.Printf(Hello("world"))
}