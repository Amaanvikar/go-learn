package main

import "fmt"

func main() {

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(days)

	for d := 1; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// for i := range days {
	// 	fmt.Println("days[i]")
	// }

	// for _, days := range days {
	// 	fmt.Println("index is and value is %v\n", days)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 3 {
			goto label
		}

		// if rogueValue == 2 {
		// 	rogueValue++
		// 	continue
		// }

		if rogueValue == 5 {
			break
		}

		fmt.Println("value is ", rogueValue)
		rogueValue++
	}

label:
	fmt.Println("jumping to label")
}
