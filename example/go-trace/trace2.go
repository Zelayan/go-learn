package main

import (
	"fmt"
	"time"
)

func main() {
	count := 5
	for i := 0; i < count; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}