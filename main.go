package main

import (
	"fmt"
	"helloworld/myutil"
)

func main() {
	fmt.Println("Learning go language")
	myutil.PrintMessage("Hello world")
	var name string = "Karan"
	var money int = 68000
	var floatmoney float64 = 69000.654
	fmt.Println(name)
	fmt.Println("money: ", money)
	fmt.Println("money in float64: ", floatmoney)
	const pi = 3.741
	fmt.Println("constant pi variable: ", pi)
	person := "Karan"
	fmt.Println("person is: ", person)

	var Public = "Data is important" //capital for exporting it now it is public variable
	var private = "Data is private"  //small for not exporting it now it will be private
	fmt.Println("Public is: ", Public)
	fmt.Println("Private is: ", private)
}
