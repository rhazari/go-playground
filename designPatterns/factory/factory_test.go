package factory

import (
	"testing"
)

func Test_Factory(t *testing.T) {
	shape1 := ShapeFactory("circle")
	shape2 := ShapeFactory("square")

	if shape1 != nil {
		shape1.Draw() // Output: Drawing a Circle
	}
	if shape2 != nil {
		shape2.Draw() // Output: Drawing a Square
	}
}