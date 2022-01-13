range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，那么 for 和 range 的性能几乎是一样，
例如 []int。

但是如果迭代的元素内存占用较高，例如一个包含很多属性的 struct 结构体，那么 for 的性能将显著地高于 range，有时候甚至会有上千倍的性能差异。
对于这种场景，建议使用 for，

如果使用 range，建议只迭代下标，通过下标访问迭代值，这种使用方式和 for 就没有区别了。
如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。

```go
package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint64
}

func main() {
	u := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	n := make([]*user, 0, len(u))
	for _, v := range u {
		n = append(n, &v)
	}
	// 在for range中，变量v是用来保存迭代切片所得的值，因为v只被声明了一次，每次迭代的值都是赋值给v，该变量的内存地址始终未变，是复制了一份，
	fmt.Println(n)
	for _, v := range n {
		fmt.Println(v)
	}
}
```

在for range中，变量v是用来保存迭代切片所得的值，因为v只被声明了一次，每次迭代的值都是赋值给v，该变量的内存地址始终未变，是复制了一份，