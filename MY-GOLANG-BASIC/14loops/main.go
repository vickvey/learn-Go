package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang!!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	fmt.Println(days)

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto here
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

here:
	fmt.Println("Jumping at www.google.com")

}
