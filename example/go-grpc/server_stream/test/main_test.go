package main

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ticker := time.NewTicker(time.Second * 5)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		select {
		case <-ticker.C:
			wg.Done()
			os.Setenv("IS_MASTER", "true")
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			getenv := os.Getenv("IS_MASTER")
			if getenv != "" {
				fmt.Println(getenv)
				wg.Done()
			}
		}
	}()
	wg.Wait()

}
