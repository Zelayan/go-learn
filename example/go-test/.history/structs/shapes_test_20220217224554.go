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
	rectangle := Rectangle {
		10.0, 10.0,
	}
	got := rectangle.Area()
	want := 100.0
	if (want != got) {
		t.Errorf("got %f want %f", got, want)
	}

	t.Run("test Circle", func(t *testing.T) {
		rectangle := c {
			10.0, 10.0,
		}
		got := rectangle.Area()
		want := 100.0
		if (want != got) {
			t.Errorf("got %f want %f", got, want)
		}
	})
	
}