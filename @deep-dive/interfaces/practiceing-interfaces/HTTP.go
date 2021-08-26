package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	httpFirstCall()
}

type logWriter struct {
}

// make my first HTTP call with :: GO ::
func httpFirstCall() {
	res, err := http.Get("http://asalih.netlify.com")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// *AN OLD WAY TO GET RESULT:
	// byteSlice := make([]byte, 999999)
	// res.Body.Read(byteSlice)

	// fmt.Println(string(byteSlice))

	// *NEW WAY TO GET RESULT:
	// io.Copy(os.Stdout, res.Body)

	// write custom log writer or what ever its name is, then passing it to Copy()
	lw := logWriter{}
	io.Copy(lw, res.Body)

}

// *CUSTOM WAY TO GET THE RESULT:
// custom Write method
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	// *******
	// you can build your custom logic here, just to satisfy your need:
	// *******

	return len(bs), nil
}
