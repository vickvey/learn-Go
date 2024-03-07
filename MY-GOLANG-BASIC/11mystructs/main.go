package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; no super or parent

	alex := User{"Alex", "alex@go.dev", true, 21}
	fmt.Println(alex)

	fmt.Printf("Alex details are: %+v\n", alex)
	fmt.Printf("Name is %v and email is %v.", alex.Name, alex.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
