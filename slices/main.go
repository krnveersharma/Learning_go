package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	fmt.Println("Length of numbers:-", len(numbers))
	fmt.Println("Capacity of numbers:-", cap(numbers))
	numbers = append(numbers, 3)
	fmt.Println("Length of numbers:-", len(numbers))
	fmt.Println("Capacity of numbers:-", cap(numbers))

}
