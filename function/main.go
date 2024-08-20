package main

import "fmt"

func simpleFunction() {
	fmt.Println("THis is a simple function")
}
func add(a int, b int) int {
	return a + b
}
func add2(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("we are learning functions in  Golang")
	simpleFunction()
	ans := add(3, 6)
	ans2 := add2(3, 6)
	fmt.Println("Addition of two numbers are:", ans)
	fmt.Println("Addition of two numbers are:", ans2)
}
