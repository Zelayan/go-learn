package main

import "fmt"

func init() {
	fmt.Println("init")
	// 不能捕获，包中 init 函数引发的 panic 只能在 init 函数中捕获，在 main 中无法被捕获
	panic("init")
}
func main() {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Printf("revocer err:", err)
		}
	}()

	// 可以捕获
	//panic("err") 
    
	// 可以捕获
	//p()
}

func p() {
	// 可以捕获
	panic("err") 
}