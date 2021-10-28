// context使用场景
package main

import (
	"context"
	"fmt"
)

// 1. 请求链路传旨

func func1(ctx context.Context) {
	ctx = context.WithValue(ctx, "k1", "v1")
	func2(ctx)
}
func func2(ctx context.Context) {
	fmt.Println(ctx.Value("k1").(string))
}

func main() {
	ctx := context.Background()
	func1(ctx)
}

// 2.取消耗时操作，及时释放资源