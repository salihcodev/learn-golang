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

	for _, li := range links {
		fmt.Println(getStatus(li))
	}
}

func getStatus(li string) string {
	_, err := http.Get(li)

	if err != nil {
		return err.Error()
	}

	return li + " [STATUS:] Is up running!"
}
