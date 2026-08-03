package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{ID: 1, Name: "Aman Pathan", Email: "aman@go.dev", Age: 24}
	fmt.Println(user)
	fmt.Println(user.Name)
}
