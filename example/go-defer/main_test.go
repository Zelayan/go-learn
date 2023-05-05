package main

import (
	"fmt"
	"testing"
)

func TestPrint2(t *testing.T) {
	Print2() // 2 1 0
	fmt.Println("")
	Print() // 3 3 3
}

func TestPrint3(t *testing.T) {
	print(Print3()) // 5
}

func TestPrint4(t *testing.T) {
	print(Print4()) // 1
}

func TestF3(t *testing.T) {
	f1() // <nil>
}

func TestClosure(t *testing.T) {
	closure()
}

func TestAcc(t *testing.T) {
	closureRemember()
}
