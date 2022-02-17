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
	t.Run("tes strangle", func(t *testing.T) {
		rectangle := Rectangle {
			10.0, 10.0,
		}
		got := rectangle.Area()
		want := 100.0
		if (want != got) {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("test Circle", func(t *testing.T) {
		circle := Circle {
			10,
		}
		got2 := circle.Area()
		want2 := 314.159265
		if (want2 != got) {
			t.Errorf("got %f want %f", got2, want2)
		}
	})
	
}