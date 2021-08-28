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

	// NOT WORKING !!
	// resolved by using channels:
	for _, li := range links {
		go getStatus(li)
	}

}

func getStatus(li string) {
	_, err := http.Get(li)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(li + " [STATUS:] Is up running!")
}
