package main

import "fmt"

func main() {
	greeter()
	greeterTwo()

	result := adder(3, 4)
	proResult := proAdder(1, 2, 3, 4, 5)

	fmt.Println("Welcome to functions in golang!!")

	fmt.Println("Result is: ", result)
	fmt.Println("Pro Result is: ", proResult)
}

func adder(x int, y int) int {
	return (x + y)
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}

func greeter() {
	fmt.Println("Namastey from golang!!")
}

func greeterTwo() {
	fmt.Println("Another method!!")
}
