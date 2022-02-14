package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return "Bonjour," + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world", ""))
}