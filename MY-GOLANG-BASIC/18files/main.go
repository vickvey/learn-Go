package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - www.google.com"

	file, err := os.Create("./18files/myfile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./18files/myfile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
