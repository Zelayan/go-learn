package structs

import "math"

// 在 Go 语言中 interface resolution 是隐式的。如果传入的类型匹配接口需要的，则编译正确。
type Shape interface {
    Area() float64
}

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

type Triangle struct {
	Base   float64
    Height float64
}

func (t Triangle) Area() float64 {
	return 
}