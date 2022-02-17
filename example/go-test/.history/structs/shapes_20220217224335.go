package structs

type Rectangle struct {
    Width float64
    Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	
	return 2 * (rectangle.Height + rectangle.Width)
}

func (reRectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}