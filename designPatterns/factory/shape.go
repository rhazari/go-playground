package factory

import "fmt"

// Shape interface
type Shape interface {
	Draw()
}

// Circle implements Shape
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

// Square implements Shape
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing a Square")
}