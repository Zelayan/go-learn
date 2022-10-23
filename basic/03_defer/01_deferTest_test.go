package main

import (
	"fmt"
	"testing"
)

// 有名函数返回值，在进入函数的时候初始化，并初始化为0值
func TestDeferFunc1(t *testing.T) {
	fmt.Printf("DeferFunc1(2): %v\n", DeferFunc1(2))
	// 0
	// 2
}

// defer遇上return
// 先执行return，再执行defer
func TestDeferAndReturn(t *testing.T) {
	returnAndDefer()
	// return func called
	// defer func called
}

// 有名函数返回值遇上defer
func TestDeferAndValueReturn(t *testing.T) {
	fmt.Printf("returnButDefer(): %v\n", returnButDefer())
}

// defer遇上panic 不捕获
func Test_defercall(t *testing.T) {
	defer_call()
}

// defer遇上panic 捕获
func Test_defercall2(t *testing.T) {
	defer_call2()
}

// defer 中发生panic
func Test_panic2(t *testing.T) {
	panic2()
	// defer panic
}

// 压zhan
func Test_functionInDefer(t *testing.T) {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}

func Test_deferWithPanic(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"succes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deferWithPanic()
		})
	}
}
