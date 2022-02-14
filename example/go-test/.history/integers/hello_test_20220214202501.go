package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello,Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc	string
		
	}{
		{
			desc: "",
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			
		})
	}
}