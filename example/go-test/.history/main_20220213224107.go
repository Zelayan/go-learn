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
    //this goroutine will wait till another goroutine received the value
    ch <- 45
}

func receivingGoRoutine(ch chan int, done chan bool){
	//this gourtine will wait till the channel received a value
    v := <- ch
	fmt.Println("Received value ", v)
	done <- true
}