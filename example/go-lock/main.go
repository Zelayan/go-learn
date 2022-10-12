package main

import (
	"fmt"
	"sync"
	"time"
)

type Session struct {
	Mux sync.Mutex
	a   int
}

func main() {
	maxActive := 10
	//minIdle := 2
	maxIdle := 5
	use := 0
	unUse := 5
	fmt.Printf("use: %d, unUse: %d\n", use, unUse)
	ticker := time.NewTicker(time.Second * 3)
	for range ticker.C {
		if maxActive == use {
			fmt.Println("等于创建的最大数量了 结束创建")
			fmt.Printf("use: %d, unUse: %d\n", use, unUse)
			return
		}
		fmt.Println("未使用减1")
		unUse = unUse - 1
		fmt.Println("分配一个一个")
		use = use + 1
		fmt.Printf("use: %d, unUse: %d\n", use, unUse)

		// check
		if unUse < maxIdle {
			if unUse+use < maxActive {
				fmt.Println("小于最大需要补充")
				unUse = unUse + 1
				fmt.Printf("use: %d, unUse: %d\n", use, unUse)
			}
		}
	}
}
