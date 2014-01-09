package go_design_patterns

// import (
//   "fmt"
// )
// Pattern
// Composite
// Proxy
// Decorator

type Rectangle struct {
	x, y float64
}

func (r Rectangle) Area() float64 {
	return (r.x * r.y)
}

type Circle struct {
	diameter float64
}

func (r Circle) Area() float64 {
	return (r.diameter * 3.14159)
}

type Square struct {
	l float64
}

func (s Square) Area() float64 {
	return (s.l * s.l)
}

type Shape interface {
	Area() float64
}

type Shapes struct {
	shapes []Shape
}

func (s Shapes) Area() float64 {
	sum := 0.0
	for _, shape := range s.shapes {
		sum += shape.Area()
	}
	return sum
}
