package main

import "fmt"

type thing []string

// this function called receiver:
func (t thing) Print() {
	for idx, val := range t {
		fmt.Println("idx:", idx, "val:", val)
	}
}
