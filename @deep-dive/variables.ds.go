package main

import "fmt"

func main() {

	/*
		there is two ways to define a vars in go
	*/

	// var card string = newCard() // #1
	card := newCard() // #2

	// print out the output
	fmt.Println(card)
}

func newCard() string {
	return "wohhoooo, New card has been created!"
}
