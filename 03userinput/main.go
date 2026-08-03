package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to the user input program"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the user: ")

	// comma ok || comma error syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("The name of the user is: ", input)
	fmt.Printf("Type of the user is %T", input)
}
