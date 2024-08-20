package main

import "fmt"

func main() {
	age := 30
	name := "Bob"
	height := 5.9
	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %f\n", height)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("name is %s\n", name)
	fmt.Printf("name is %s, age is %d, height is %0.2f", name, age, height)
}
