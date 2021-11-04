package _2_OutTime

import (
	"fmt"
	"time"
)

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}

}

func timeout(f func(chan bool)) error {
	// 会阻塞
	 done := make(chan bool)
	//done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}