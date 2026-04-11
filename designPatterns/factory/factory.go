package factory

import "fmt"

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		fmt.Println("Unknown shape type")
		return nil
	}
}