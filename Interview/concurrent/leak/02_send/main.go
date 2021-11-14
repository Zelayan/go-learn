package main

import (
	"fmt"
	"runtime"
)

func query(ch <-chan int) int {
	for i := 0; i < 1000; i++ {
		go func() { <- ch }()
	}
	return 0
}


func main() {
	ch := make(chan int)
	for i := 0; i < 4; i++ {
		query(ch)
		fmt.Printf("goroutines: %d", runtime.NumGoroutine())
	}

}