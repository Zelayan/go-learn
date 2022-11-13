package _interface

import "testing"

func Test_pointStruct(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointStruct()
		})
	}
}
