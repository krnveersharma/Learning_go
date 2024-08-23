package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	var str string = "one two three two two four"
	count := strings.Count(str, "two")
	fmt.Println("Count:-", count)

	str = "   Hello,     Go!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)
	str1 := "Karanveer"
	str2 := "Sharma"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)
}
