package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

//Rectangle
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

//Circle

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	var s Shape
	s = Rectangle{5, 10}
	fmt.Println("Rectangle Area", s.Area())
	fmt.Println("Rectangle Perimeter", s.Perimeter())

	s = Circle{10}
	fmt.Println("Circle Area", s.Area())
	fmt.Println("Circle Perimeter", s.Perimeter())
}
