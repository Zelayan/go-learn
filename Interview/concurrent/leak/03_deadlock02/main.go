// 串联信道 单一挂起导致死锁
package main

import "fmt"

var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)

func say(s string) {
	fmt.Println(s)
	ch1 <- <- ch2 // ch1 等待 ch2流出的数据
}

func main() {
	go say("hello")
	<- ch1  // 堵塞主线
}
