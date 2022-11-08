package main

import (
	"testing"
)

func Test_initSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"successs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initSlice()
		})
	}
}

func Test_newSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"successs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newSlice()
		})
	}
}

func Test_nilSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "succes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nilSlice()
		})
	}
}
