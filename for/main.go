package main

import "fmt"

func main() {
	var x int
	x = 10
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	counter := 0
	for {
		if counter == 10 {
			break
		}
		counter++
	}
	numbers := []int{1, 2, 3, 5}
	for index, value := range numbers {
		fmt.Printf("value at index %d is %d\n", index, value)
	}
}
