package main

import "fmt"

type Shape struct {
	x int
	y int
}

func main() {

	circle := Shape{x: 1, y: 9}

	fmt.Println(circle)
}
