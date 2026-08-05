package main

import "fmt"

func main() {

	fmt.Println("Methods is golang")

	user := User{"Aman", "aman@gmail.com", true, 24}
	fmt.Println(user)
	fmt.Println("user details are: ", user)
	fmt.Printf("user details are: %+v\n", user)
	fmt.Println(user.Name, "\n", user.Email, "\n", user.isActive, "\n", user.Age)
	user.GetStatus()
	user.NewMail()
	fmt.Println("Email of this user is: ", user.Email)
}

type User struct {
	Name     string
	Email    string
	isActive bool
	Age      int
}

func (u User) GetStatus() string {
	fmt.Println("User is active: ", u.isActive)
	return "Active"
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
