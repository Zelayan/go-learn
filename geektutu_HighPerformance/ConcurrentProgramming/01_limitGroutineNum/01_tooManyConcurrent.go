package main

import (
	"math"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	// panic: too many concurrent operations on a single file or socket (max 1048575)
}