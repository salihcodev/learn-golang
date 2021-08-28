package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://github.com",
		"https://stackoverflow.com",
		"https://google.com",
		"https://aslih.netlify.com",
		"https://youtube.com",
	}
	c := make(chan string)

	for _, li := range links {
		go getStatus(li, c)
	}

	// here gonna receive only one response, then the main routine gonna close the channel an exits.
	// fmt.Print(<-c)

	/*
		to make it wake for another msg that gonna be received via channel
		looping, is the way to go here:
	*/

	for range links {
		fmt.Println(<-c)
	}

}

func getStatus(li string, c chan string) {
	_, err := http.Get(li)

	if err != nil {
		c <- err.Error()
	}

	c <- li + " ==> [STATUS]: Is up running!"
}
