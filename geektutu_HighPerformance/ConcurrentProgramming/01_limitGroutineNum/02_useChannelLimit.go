package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
    errorChan := make(chan error)
	for i :=0; i < 10; i++ {
		ch <- struct{}{}
		go func(i int) {
			handleIterm(i, ch, errorChan, &wg)
		}(i)
	}

	// must use goroutine to run this method
	go func() {
		for {
			select {
			case c, ok := <- errorChan:
				if ok {
					fmt.Println(c.Error())
				} else {
					break
				}
			}
		}
	}()

	time.Sleep(time.Second * 2)

	wg.Wait()
}

func handleIterm(item int, ch chan struct{}, errorChan chan error, wg *sync.WaitGroup) (i int, err error) {
	wg.Add(1)
	// handle gorouting hanle
	defer func(group *sync.WaitGroup) {
		group.Done()
		if p := recover(); p != nil {
			errs := fmt.Errorf("go routine error %v", p)
            errorChan <- errs
		}

	}(wg)
	time.Sleep(time.Second)
	log.Println("item is : ", item)
	<- ch
	panic("error ...")

	return 1, err
}
