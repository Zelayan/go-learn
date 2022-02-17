package structs

import "math"

type Rectangle struct {
    Width float64
    Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	
	return 2 * (rectangle.Height + rectangle.Width)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}