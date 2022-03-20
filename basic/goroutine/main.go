package goroutine

import "fmt"

// 循环打印 abc
func Print3() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 2)
	ch3 := make(chan struct{}, 2)
	ch3 <- struct{}{}
	for i := 0; i < 100; i++ {
		go func() {
			<-ch3
			fmt.Println("a")
			ch1 <- struct{}{}
		}()

		go func() {
			<-ch1
			fmt.Println("b")
			ch2 <- struct{}{}
		}()

		go func() {
			<-ch2
			fmt.Println("c")
			ch3 <- struct{}{}
		}()
	}

}
