package main

import "fmt"

func main() {
	fmt.Println("functions is golang")

	result := add(2, 3)
	fmt.Println(result)
}

func add(a int, b int) int {
	return a * b
}
