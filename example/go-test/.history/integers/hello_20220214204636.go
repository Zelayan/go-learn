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

	s
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {

}

func main() {
	fmt.Printf(Hello("world", ""))
}