package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("num is", num)
	fmt.Printf("Type of num is %T\n", num)
	var data float64 = float64(num)
	data += 1.99
	fmt.Println("num is", data)
	fmt.Printf("Type of num is %T\n", data)
	str := strconv.Itoa(num)
	fmt.Println("num is", str)
	fmt.Printf("Type of num is %T\n", str)
}
