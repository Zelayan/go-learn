package fib

import "testing"

func TestFib(t *testing.T) {
	var (
		in       = 7
		excepted = 13
	)

	actual := Fib(7)

	if actual != excepted {
		t.Errorf("Fib(%d) == %d, excepted is %d", in, actual, excepted)
	}
}

func TestFib2(t *testing.T) {
	testCases := []struct {
		in       int
		excepted int
	}{
		{
		},
	}
	for _, tC := range testCases {
		actual := Fib(tC.in)
		if actual != tC.excepted {
			t.Errorf("Fib(%d) == %d, excepted is %d", tC.in, actual, tC.excepted)
		}
	}
}
