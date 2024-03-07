package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println(ptr)
	// fmt.Printf("Variable is of type: %T\n", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)
}
