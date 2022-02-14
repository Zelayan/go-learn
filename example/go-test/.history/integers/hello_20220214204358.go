package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}
	
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world", ""))
}