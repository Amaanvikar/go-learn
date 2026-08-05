package main

import "fmt"

func main() {

	fmt.Println("Defer in golang")
	defer fmt.Println("world")
	fmt.Println("hello")
	deferFunc()

}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
