package main

import (
	"encoding/json"
	"sync"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func unmarsh() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}

// 声明对象池
var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func Usage01() {
	student := studentPool.Get().(*Student)
	json.Unmarshal(buf, student)
	// Put 是在对象使用完毕后，返回对象池。
	studentPool.Put(student)
}


