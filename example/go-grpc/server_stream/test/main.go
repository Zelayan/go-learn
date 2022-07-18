package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(3 * time.Second)
			getenv := os.Getenv("JAVA_HOME")
			if getenv != "" {
				fmt.Println(getenv)
				wg.Done()
			}
			fmt.Println("1...")
		}
	}()
	wg.Wait()
}
