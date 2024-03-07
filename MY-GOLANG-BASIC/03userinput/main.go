package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	welcome := "Welcome to user input!!"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || error ok syntax
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid input!!")
	}
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
