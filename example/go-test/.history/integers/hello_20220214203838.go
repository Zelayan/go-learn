package main

import "fmt"

const englishHelloPrefix = "Hello,"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language = ""
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world", ""))
}