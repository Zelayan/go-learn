package structs

import (
	"crypto/des"
	"testing"
)

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
	checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    }

	t.Run("tes strangle", func(t *testing.T) {
		rectangle := Rectangle {
			10.0, 10.0,
		}
		want := 100.0
		checkArea(t, rectangle, want)
	})

	t.Run("test Circle", func(t *testing.T) {
		circle := Circle {
			10,
		}
		want2 := 314.1592653589793
		checkArea(t, circle, want2)
	})
	
}

// 表格驱动测试
func Test(t *testing.T) {
	testCases := []struct {
		desc string
		shape Shape
		want  float64
		
	}{
		{	"rectangle",
			Rectangle{10, 10}, 100,
			
		},
		{	
			"circle",
			Circle{10}, 314.1592653589793,
		},
		{
			"triangle",Triangle{12, 6}, 36.0},
	}
	for _, tt := range testCases {
		got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %.2f want %.2f", got, tt.want)
        }
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc	string
		
	}{
		{
			desc: "",
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			
		})
	}
}