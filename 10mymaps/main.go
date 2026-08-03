package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["Go"] = "Golang"
	m["Dt"] = "Dart"
	m["Js"] = "Javascript"

	fmt.Println(m)
	fmt.Println(m["Go"])

	delete(m, "Js")
	fmt.Println(m)

	// loops through the map

	for key, value := range m {
		fmt.Println(key, value)
	}

	for _, value := range m {
		fmt.Println(value)
	}

}
