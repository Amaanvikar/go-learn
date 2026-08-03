package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("value of pointer is", ptr)

	var a = 10

	ptr = &a

	fmt.Println("value of a:", a)
	fmt.Println("value of pointer which is direct reference to memory address of a:", ptr)
	fmt.Println("value of pointer which is refering to a using &a is:", *ptr)
}
