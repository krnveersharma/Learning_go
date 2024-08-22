package main

import "fmt"

func main() {
	x := 10
	switch x {
	case 1, 8, 10:
		fmt.Println("Multiply")
	case 2, 5, 7:
		fmt.Println("Monday")
	case 3, 4, 6:
		fmt.Println("Tools")
	default:
		fmt.Println("Unknown thing")
	}

}
