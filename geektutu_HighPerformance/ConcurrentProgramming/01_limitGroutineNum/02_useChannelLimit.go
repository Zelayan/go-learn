package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)

	for i :=0; i < 10; i++ {
		ch <- struct{}{}
		go func() {
			err := handleIterm(i, ch, wg)
			if err != nil {
				log.Printf("main error %v", err)
			}
		}()
	}

	wg.Wait()
}

func handleIterm(item int, ch chan struct{}, wg sync.WaitGroup) error {
	wg.Add(1)
	// handle gorouting hanle
	defer func(group sync.WaitGroup) {
		group.Done()
		err := recover()
		if err != nil {
			log.Println("panic error...")
			//return error.Error(fmt.Sprintf(""))
		}
		<- ch
	}(wg)
	time.Sleep(time.Second)
	log.Println("item is : ", item)
	panic("error ...")

}
