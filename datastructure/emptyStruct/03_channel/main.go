package main

import (
	"fmt"
	"time"
)

func worker(ch chan struct{}) {
	<- ch
	fmt.Println("do something")
	close(ch)
}

func main() {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}

	time.Sleep(time.Second)
}
