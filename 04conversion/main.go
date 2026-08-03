package main // main is special: it tells Go that this is an executable program.

import (
	"bufio"   // Reading input efficiently
	"fmt"     // Printing to the console
	"os"      //Accessing operating-system resources, here os.Stdin
	"strconv" // Converting strings into numbers
	"strings"
)

func main() { // main() is the entry point of your program.

	fmt.Println("Welcome to out pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err, "Please enter a valid rating")

	} else {

		fmt.Println("Thanks for rating, ", numRating)

	}

}
