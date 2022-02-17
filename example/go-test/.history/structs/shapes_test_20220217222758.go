package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle {
		10.0, 10.0,
	}

	got := Perimeter()	
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