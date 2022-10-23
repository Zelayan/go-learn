package main

import (
	"fmt"
	"reflect"
)

func compareStruct() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	

	sn3 := struct {
		age  int
		name string
	}{age:11, name:"qq"}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sn1 == sn3 {
		fmt.Println("sn1 = sn3")
	}

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Printf("sm1 == sm2")
	}
}