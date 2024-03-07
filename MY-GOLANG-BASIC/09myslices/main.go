package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices tutorial!!")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Types of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)
	fmt.Println(highScores)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 567
	// highScores[4] = 777
	highScores = append(highScores, 555, 663, 321)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value form slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}