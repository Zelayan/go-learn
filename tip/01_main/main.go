package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// http服务
/*func main() {
	// 创建一个 sig 的 channel，捕获系统的信号，传递到sig中
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	// http服务改造成异步
	go http.ListenAndServe(":8080", mux)

	// 程序阻塞在这里，除非收到了interrupt或者kill信号
	fmt.Println(<-sig)
}*/

// onetoone 最直观的解决方案 - 2个channel 互相通知
/*func main() {
	sig := make(chan os.Signal)

	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh, finishChar chan struct{}) {
		for {
			select {
			case <- stopCh:
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
			

		}
	}(stopCh, finishedCh)

	<-sig
	stopCh <- struct{}{}

	<-finishedCh
	fmt.Println("finoshed")
}
*/
// 华丽的解决方案 - channel嵌套channel

/*func main() {
	sig := make(chan os.Signal)
	stopCh := make(chan chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopChh chan chan struct{}) {
		for {
			select {
			case ch := <-stopCh:
				// 结束后，通过ch通知主goroutine
				fmt.Println("stopped")
				ch <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	<-sig
	// ch作为一个channel，传递给子goroutine，待其结束后从中返回
	ch := make(chan struct{})
	stopCh <- ch
	<-ch
	fmt.Println("finished")
}*/

// 标准解决方案 - 引入上下文context

/*func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	gorouting, cancel := 01_context.WithCancel(01_context.Background())
	finishedCh := make(chan struct{})

	go func(gorouting 01_context.Context, finishedCh chan struct{}) {
		for {
			select {
			case <-gorouting.Done():
				// 结束后，通过ch通知主goroutine
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(gorouting, finishedCh)

	<-sig
	cancel()
	<-finishedCh
	fmt.Println("finished")
}*/

// 1对多
//
func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	num := 10

	// 用wg来控制多个子goroutine的生命周期
	wg := sync.WaitGroup{}
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("stopped")
					return
				default:
					time.Sleep(time.Duration(i) * time.Second)
				}
			}
		}(ctx)
	}

	<-sig
	cancel()
	// 等待所有的子goroutine都优雅退出
	wg.Wait()
	fmt.Println("finished")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}