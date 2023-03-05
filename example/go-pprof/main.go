package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	readChan := make(chan int, 1)

	go func() {
		defer close(readChan)
		for i := 0; i < 100; i++ {
			readChan <- i
		}

	}()

	for i := 0; i < 100; i++ {
		go func() {
			for i2 := range readChan {
				fmt.Print(i2)
			}
		}()
	}

	http.ListenAndServe("0.0.0.0:6060", nil)
}
