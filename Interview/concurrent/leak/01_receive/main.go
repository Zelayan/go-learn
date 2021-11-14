package main

import (
	"fmt"
	"runtime"
)

func query() int {
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() { ch <- 0 }()
	}
	return <-ch
}


func main() {

	for i := 0; i < 4; i++ {
		query()
		fmt.Printf("goroutines: %d", runtime.NumGoroutine())
	}

}
