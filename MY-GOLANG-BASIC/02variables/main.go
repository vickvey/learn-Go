package main

import "fmt"

// You make first later capital to make it public and accessible to other go files
const LoginToken = "hgsdihfs" // Public

func main() {
	var username string = "Alex"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float32 = 255.347329846239843928
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var largeFloat float64 = 255.347329846239843928
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T\n", largeFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)

	// wallrus operator
	numberOfUser := 3000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T\n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
