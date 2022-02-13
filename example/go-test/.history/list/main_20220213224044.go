package main

import (
    "fmt"
    "time"
)

func main() {
	ch := make(chan int)
	//create a channel to prevent the main program from exiting before the done signal is received
	done := make(chan bool)
	go sendingGoRoutine(ch)
	go receivingGoRoutine(ch,done)
	//This will prevent the program from exiting till a value is sent over the "done" channel, value doesn't matter
	<- done
}

func sendingGoRoutine(ch chan int){
	//start a timer to wait 5 seconds
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