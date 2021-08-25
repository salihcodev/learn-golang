package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	ReadFile()
}

// ----- reading a FILE -----
func ReadFile() {
	fileContent, err := ioutil.ReadFile("testingIO")

	if err != nil { // err not nil.. so is there error occurred, then return it.
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("File contents", fileContent)
	}
}
