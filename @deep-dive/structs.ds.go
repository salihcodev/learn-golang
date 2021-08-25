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
	shapeProps  shapeProps
	// only "shapeProps" reduces "shapeProps  shapeProps" AKA shorthand like so:
	// shapeProps
}

func main() {
	circle := shape{boldBorders: true, area: 100, shapeProps: shapeProps{x: 1, y: 5}}

	circle.printing()
}

// receiver function:
func (key shape) printing() {
	fmt.Println(key)
}
