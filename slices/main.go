package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 4, 5)
	fmt.Println("Slice:", numbers)
	fmt.Println("Length of Slice is:", len(numbers))
	name := []string{}
	fmt.Println("name:", name)
	name = append(name, "Karan")
	fmt.Println("name:", name)
}
