package go_select

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_select(t *testing.T) {

	c1 := make(chan int, 1)
	c2 := make(chan int, 5)
	c1 <- 10
	var wg sync.WaitGroup
	wg.Add(1)
	ticker := time.NewTicker(time.Second).C
	i := 1

	go func() {
		for range ticker {
			c2 <- i
			i++
			if i == 10 {
				wg.Done()
			}
			if i == 5 {
				c1 <- 5
			}
		}
	}()

	go func() {
		for {
			select {
			case x := <-c1:
				fmt.Printf("c1: %v\n", x)
				time.Sleep(time.Second * 10)
				fmt.Printf(" done c1: %v\n", x)
			case x2 := <-c2:
				fmt.Printf("c2: %v\n", x2)
			}
		}
	}()

	wg.Wait()
}

func Test_rangemap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "aa"
	m["b"] = "bb"
	m["c"] = "cc"
	for k := range m {
		go func(k string) {
			fmt.Println(k)
		}(k)
	}
}
