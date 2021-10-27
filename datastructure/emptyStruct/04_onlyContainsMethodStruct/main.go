package main

import "fmt"

// 在部分场景下，结构体只包含方法，不包含任何的字段

type Door struct {

}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}

