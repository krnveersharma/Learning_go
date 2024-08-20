package main

import "fmt"

func main() {
	fmt.Println("We are learnig array")
	var name [5]string
	var name2 [5]string
	name[0] = "karan"
	name[1] = "Bawa"
	name2[4] = "karan"
	fmt.Println("name of the person is:", name)
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers:", numbers)
	fmt.Println("number at index 4 is:", numbers[4])
	fmt.Printf("name2: %q\n", name2) //%q means quoted string
}
