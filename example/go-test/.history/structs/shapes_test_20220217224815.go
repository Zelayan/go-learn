package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle {
		10.0, 10.0,
	}

	got := Perimeter(rectangle)	
	want := 40.0
	if (want != got) {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("tes stran")

	t.Run("test Circle", func(t *testing.T) {
		circle := Circle {
			10,
		}
		got := circle.Area()
		want := 314.159265
		if (want != got) {
			t.Errorf("got %f want %f", got, want)
		}
	})
	
}