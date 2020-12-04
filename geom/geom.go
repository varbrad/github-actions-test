package geom

import "math"

// Perimeter calculates the perimeter of a rectangle
func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle shape
type Circle struct {
	Radius float64
}

// Shape interface
type Shape interface {
	Area() float64
}
