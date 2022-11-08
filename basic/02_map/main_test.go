package main

import "testing"

func Test_foreachMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foreachMap()
		})
	}
}

func Test_nilMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nilMap()
		})
	}
}
