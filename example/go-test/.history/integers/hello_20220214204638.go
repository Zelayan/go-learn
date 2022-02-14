package main

import (
	"fmt"

	"golang.org/x/text/language"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix	
	}
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	
}

func main() {
	fmt.Printf(Hello("world", ""))
}