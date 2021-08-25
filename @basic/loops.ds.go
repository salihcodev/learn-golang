package main

import "fmt"

// :: app entry::
func main() {

	cards := []string{newCard("new card #1"), newCard("new card #2")}

	// appending some dummy data to cards slice:
	cards = append(cards, newCard("new card #3 'appended'"))

	// ----- FOR loop -----
	fmt.Println(":::::::: for loop ::::::::")
	for iterator := 0; iterator < len(cards); iterator++ {
		fmt.Println(cards[iterator])
	}

	// ----- WHILE loop -----
	iterator := 0
	fmt.Println(":::::::: while loop ::::::::")
	for iterator < len(cards) {
		fmt.Println(cards[iterator])
		iterator++
	}

	// ----- looping via RANGE -----
	fmt.Println(":::::::: looping via range ::::::::")
	for idx, value := range cards {
		fmt.Println("index:", idx, "value:", value)
	}

}

// :: card creator::
func newCard(name string) string {

	return name
}
