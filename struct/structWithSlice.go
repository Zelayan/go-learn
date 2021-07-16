package main

import "fmt"

type Person1 struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person1) SetDreams(dreams []string) {
	p.dreams = dreams
}

func main() {
	p1 := Person1{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  // ?
}