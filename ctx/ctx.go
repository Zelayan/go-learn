package main

import (
	"context"
	"fmt"
	"time"
)

// Go 语言中的 context.Context 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用，虽然它也有传值的功能，但是这个功能我们还是很少用到。
//在真正使用传值的功能时 我们也应该非常谨慎，使用 context.Context 传递请求的所有参数一种非常差的设计，比较常见的使用场景是传递请求对应用户的认证令牌以及用于进行分布式追踪的请求 ID


// Tip: 通过 cancel 主动关闭
func ctxCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Time out")
		}
	}(ctx)

	cancel()
}

// Tip: 通过超时，自动触发
func ctxTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	// 主动执行cancel，也会让协程收到消息
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Time out")
		}
	}(ctx)

	time.Sleep(time.Second)
}

// Tip: 通过设置截止时间，触发time out
func ctxDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Millisecond))
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Millisecond * 100):
			fmt.Println("Time out")
		}
	}(ctx)

	time.Sleep(time.Second)
}

// Tip: 用Key/Value传递参数，可以浅浅封装一层，转化为自己想要的结构体
func ctxValue() {
	// ctx := context.WithValue(context.Background(), "user1", "chenzeze") 不会输出
	ctx := context.WithValue(context.Background(), "user", "chenzeze")
	go func(ctx context.Context) {
		v, ok := ctx.Value("user").(string)
		if ok {
			fmt.Println("pass user value", v)
		}
	}(ctx)
	time.Sleep(time.Second)
}
