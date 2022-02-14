package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello,"
const 

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return "Hola," + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("world", ""))
}