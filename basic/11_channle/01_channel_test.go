package _1_channle

import (
	"testing"
	"time"
)

func TestWithoutBufferChannle(t *testing.T) {
	t.Run("without buffer channel", func(t *testing.T) {
		// WithoutBufferChannle 没用缓冲的channel
		// 发送者会阻塞 直到接受者接受
		ch1 := make(chan struct{})
		ch2 := make(chan struct{})
		ChannleChannle(ch1, ch2)
		// ticker
	})

	t.Run("channel", func(t *testing.T) {

		ch1 := make(chan struct{})
		ch2 := make(chan struct{})
		go ChannleChannle(ch1, ch2)

		go func(chan struct{}) {
			time.Sleep(time.Second)
			ch1 <- struct{}{}
			t.Log("send done")
		}(ch1)

		time.Sleep(time.Second * 10)
	})
}
