package main

import "fmt"

func main() {

	fmt.Println("welcome to arrays")

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "mango"

	fmt.Println("fruits array is:", fruits)
	fmt.Println(len(fruits))
}
