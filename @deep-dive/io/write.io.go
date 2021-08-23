package main

import (
	"io/ioutil"
	"log"
)

func main() {
	writeFiles()
}

// ----- writeing a FILE -----
func writeFiles() {
	msg := []byte("Hi there !!")

	err := ioutil.WriteFile("testingIO", msg, 0666)

	if err != nil {
		log.Fatal(err)
	}
}
