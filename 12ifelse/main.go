package main

import "fmt"

func main() {

	loginCount := 11
	var result string

	if loginCount < 10 {
		result = "Not a Regular user"
	} else if loginCount <= 10 {
		result = "watch out"
	} else {
		fmt.Println("Regular user")
	}

	fmt.Println(result)

	if 9/2 == 4 {
		fmt.Println("Num is Even")
	} else {
		fmt.Println("Num is Odd")
	}

	if num := 2; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

}
