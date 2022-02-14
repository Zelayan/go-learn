package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == ""
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world", ""))
}