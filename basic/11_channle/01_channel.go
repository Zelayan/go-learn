package _1_channle

import (
	"fmt"
	"time"
)

func ChannleChannle(ch1 chan struct{}, ch2 chan struct{}) {
	ticker := time.NewTicker(3 * time.Second)

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	case <-ticker.C:
		fmt.Println("ticker")
	}

}
