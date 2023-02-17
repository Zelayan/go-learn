package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 创建一个接收系统信号的channle
	signals := make(chan os.Signal, 1)

	// 当收到中断信号和终止信号时，通知我们自定义的chanel
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个协程 监听我们自定义的channle
	go func() {
		sig := <-signals
		fmt.Printf("recived: %s\n", sig)
		os.Exit(0)
	}()

	fmt.Println("running")
	select {}
}
