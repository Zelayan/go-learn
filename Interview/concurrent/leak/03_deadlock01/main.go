// 单一goroutine 死锁
package main

func main() {
	ch := make(chan int)

	// 无缓冲的信道在收消息和发消息的时候，goroutine都处于挂起状态。除非另一端准备好，否则goroutine无法继续往下执行。
	// main goroutine 已经挂起
	ch <- 1

	<-ch
}
