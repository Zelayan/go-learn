package structs

import "testing"

func TestPerimeter(t *testing.T) {
	re

	got := Perimeter(10.0, 10.0)	
	want := 40.0
	if (want != got) {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0
	if (want != got) {
		t.Errorf("got %f want %f", got, want)
	}
	
}