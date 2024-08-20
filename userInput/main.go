package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Your name is", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') //read till new line and save it to name
	fmt.Println("Your name is", name)
}
