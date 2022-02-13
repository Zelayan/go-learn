package fib

import "testing"

func TestFib(t *testing.T) {
	var (
		in = 7
		excepted = 13
	)

	actual := Fib(7)

	if actual != excepted {
		
	}
}