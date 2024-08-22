package main

import "fmt"

func main() {
	x := 0
	if x == 0 {
		fmt.Println("Well done")
	} else if x < 10 {
		fmt.Println("So so")
	} else {
		fmt.Println("value of x:-", x)
	}
}
