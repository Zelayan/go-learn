package main

import (
	"fmt"
	"sync"
)

func Print2() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan int, 1)

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("doneA")
			close(ch1)
		}()
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				ch1 <- i
			} else {
				fmt.Printf("A:%d\n", i)
			}
		}
	}()

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("doneB")
		}()
		for i := range ch1 {
			fmt.Printf("b:%d\n", i)
		}
	}()

	wg.Wait()

}

// Print3 循环打印 abc
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

// 使用两个协程交替打印数字和字母， 使用channel进行通信
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func printZimuNumber() {
	letter, number := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				letter <- struct{}{}
			}
		}
	}()

	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wg.Done()
					return
				}
				fmt.Println(string(i))
				i++
				fmt.Println(string(i))
				i++
				number <- struct{}{}

			}
		}
	}()

	number <- struct{}{}
	wg.Wait()

}
