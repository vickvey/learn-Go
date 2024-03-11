package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// No inheritance in golang; No super or parent

	alex := User{"Alex", "alex@go.dev", true, 21}
	fmt.Println(alex)

	fmt.Printf("Alex details are: %+v\n", alex)
	fmt.Printf("Name is %v and email is %v.\n", alex.Name, alex.Email)

	alex.GetStatus()
	alex.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"

	fmt.Println("Email of this user is: ", u.Email)
}
