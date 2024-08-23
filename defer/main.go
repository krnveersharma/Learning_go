package main

import "fmt"

func add(a int, b int, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("Starting of the program")
	data := add(6, 5, 6)
	defer fmt.Println("Data is ", data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}

//Defers operate in stack
