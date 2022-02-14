package main

import "fmt"

const englishHelloPrefix = "Hello,"

func Hello(name string, lang) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world"))
}