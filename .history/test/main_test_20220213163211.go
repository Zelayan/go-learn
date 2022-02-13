package main

import (
	"testing"
)

testing.func Test(t *testing.T) {
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

func Test_CloseChan(t *testing.T) {
	CloseChan()
}
