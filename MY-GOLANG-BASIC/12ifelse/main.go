package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regualar user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	/// TODO: complete here

}
