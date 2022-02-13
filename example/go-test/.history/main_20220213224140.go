package main

import (
    "fmt"
    "time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sendingGoRoutine(ch)
	go receivingGoRoutine(ch,done)
	<- done
}

func sendingGoRoutine(ch chan int){
	t := time.NewTimer(time.Second*5)
	<- t.C
	fmt.Println("Sending a value on a channel")
    ch <- 111
}

func receivingGoRoutine(ch chan int, done chan bool){
    v := <- ch
	fmt.Println("Received value ", v)
	done <- true
}