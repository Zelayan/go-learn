package main

import (
	"fmt"
	"reflect"
)

func main() {
	of := reflect.ValueOf(fmt.Sprintf("test %v", "test"))
	fmt.Printf(of.String())
}
