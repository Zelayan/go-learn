package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)	
	want := x00.0
	if (want != got) {
		t.Errorf("got %f want %f", got, want)
	}
}