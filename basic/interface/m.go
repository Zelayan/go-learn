package _interface

import "fmt"

type My interface {
	SetName1(name string)
}

type MyStruct struct {
	Name string
}

func (s MyStruct) SetName1(name string) {
	s.Name = name
}

func (s *MyStruct) SetName2(name string) {
	s.Name = name
}

func pointStruct() {
	m := MyStruct{Name: "m"}
	m.SetName1("1231")
	fmt.Println(m)
}
