package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:-", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)
	formatted := currentTime.Format("02-01-2006,Monday,15:04:05") //it is to get dd-mm-yy,Day,hh:mm:ss
	fmt.Println("Formatted time:-", formatted)
}
