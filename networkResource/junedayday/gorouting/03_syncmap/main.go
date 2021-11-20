package main

import (
	"fmt"
	"sync"
)

func main() {

	var m sync.Map
	m.Store("a",1)
	fmt.Println(m.Load("a"))
	fmt.Println("-------")
	fmt.Println(m.LoadOrStore("test", 1))
	fmt.Println(m.LoadOrStore("test",2))
	fmt.Println(m.LoadOrStore("test2",3))
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v, %v",key, value)
		// tip:false means stop
		return false
	})
}
