package go_reflect

import "testing"

func TestIndirectDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IndirectDemo()
		})
	}
}

func Test_typeOfDemo(t *testing.T) {
	typeOfDemo()
}

func Test_valueOfDemo(t *testing.T) {
	valueOfDemo()
}

func TestNewDemo(t *testing.T) {
	NewDemo()
}

func TestNewSliceDemo(t *testing.T) {
	NewSliceDemo()
}

func TestNewMapDemo(t *testing.T) {
	NewMapDemo()
}
