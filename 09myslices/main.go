package main

import "fmt"

func main() {

	var fruitlist = []string{"apple", "banana", "mango", "pineapple"}

	fmt.Println("length of fruitlist is:", len(fruitlist))

	fruitlist = append(fruitlist, "tomato", "lime")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println(highScores)
	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	// how to remove a value from a slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...) // ... is used to expand the slice

	fmt.Println(courses)

}
