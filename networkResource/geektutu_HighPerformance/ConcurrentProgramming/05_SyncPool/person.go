package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool {
		New: func()interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

func PersonUsage() {

	initPool()

	person := pool.Get().(*Person)
	fmt.Println("first get person: ", person)

	person.Name = "test"

	fmt.Printf("设置 p.Name = %s\n", person.Name)

	pool.Put(person)

	// 就是上次存入的对象
	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}
