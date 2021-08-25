package main

import "fmt"

// nested struct || imbedded struct
type shapeProps struct {
	x int
	y int
}

// struct is a way to descripe the data and data types.
type shape struct {
	boldBorders bool
	area        int
	shapeProps
}

func main() {
	circle := shape{boldBorders: true, area: 100, shapeProps: shapeProps{x: 1, y: 5}}

	// before update:
	fmt.Println("before update:", circle)

	sProps := shapeProps{x: 10, y: 35}
	circle.updateShape(&sProps)

	// after update:
	fmt.Println("after update:", circle)
}

// receiver function:
func (s *shape) updateShape(updatedShapeProps *shapeProps) {
	s.shapeProps = *updatedShapeProps
}
