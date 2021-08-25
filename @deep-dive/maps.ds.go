package main

import "fmt"

func main() {
	/*
		you can define map this way:
			- testMap := make(map[string]string)

		or this way like so:
			- var testMap map[string]string
	*/
	testMap := make(map[string]string)

	// create map:
	testMap["key1"] = "value1"
	testMap["key2"] = "value2"
	testMap["key3"] = "value3"

	// deleting map element via key:
	delete(testMap, "key2")

	// iterating over maps:
	for key, value := range testMap {
		fmt.Println(key, value)
	}

	fmt.Println(testMap)
}
