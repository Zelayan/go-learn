package main


import "fmt"

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

	// 直接访问内部类型的方法
	c.userInfo.notify()  // Send email to lioney<lioney_liu@sina.com>

	// 内部类型的方法也被提升到外部类型
	c.notify()  // Send email to lioney<lioney_liu@sina.com>
}
