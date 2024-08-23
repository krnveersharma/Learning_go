package main

import "fmt"

func main() {
	var num int
	num = 2
	var ptr *int = &num
	fmt.Println(*ptr)
}
