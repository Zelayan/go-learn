package main

import "fmt"

// 定义一个通知行为的接口
type notifier interface {
	notify()
}

// 定义一个接收接口类型参数的函数
func sendNotification(n notifier) {
	n.notify()
}

// 自定义类型
type userInfo struct {
	name string
	email string
}

// 自定义类型
type client struct {
	userInfo // 嵌入类型
	level string
}

// 定义user指针类型调用的方法
func (u *userInfo) notify() {
	// 模拟发邮件的功能
	fmt.Printf("Send email to %s<%s>\n", u.name, u.email)
}


func main() {
	// 创建一个client
	c := client{
		userInfo:userInfo{
			name:"lioney",
			email:"lioney_liu@sina.com",
		},
		level: "normal",
	}
	// 实现接口的内部类型的方法，被提升到了外部类型
	sendNotification(&c)  // Send email to lioney<lioney_liu@sina.com>

	// 多态调用必须遵守方法集规则
	//sendNotification(c)  // client does not implement notifier (notify method has pointer receiver)
}
