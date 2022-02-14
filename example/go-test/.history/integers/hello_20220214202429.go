package main

import "fmt"

const englishHelloPrefix = "Hello,"

func Hello(name string) string {
	
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world"))
}