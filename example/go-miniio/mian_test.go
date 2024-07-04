package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	switch 1 {
	case 1:
		fmt.Printf("xx")
	case 2:
		fmt.Printf("xx1")
	default:
		fmt.Printf("123")

	}
}
