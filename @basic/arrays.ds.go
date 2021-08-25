package main

import "fmt"

func main() {

	cards := []string{newCard("new card #1"), newCard("new card #2")}

	// appending new card later on the function's body logic:
	cards = append(cards, newCard("new card #3 'appended'"))

	fmt.Println(cards)
}

func newCard(name string) string {

	return name
}
