package main

import "fmt"

func main() {
	//by default map is unordered
	StudentGrades := make(map[string]int)
	StudentGrades["Karan"] = 100
	StudentGrades["Alice"] = 90
	StudentGrades["Bob"] = 90
	StudentGrades["Charlie"] = 90
	fmt.Println("marks of Alice is:-", StudentGrades["Alice"])
	delete(StudentGrades, "Alice")
	//checking if a key exist
	grades, exist := StudentGrades["Alice"]
	fmt.Println("marks of Alice is:-", grades)
	fmt.Println("Alice exists?", exist)
	for index, value := range StudentGrades {
		fmt.Printf("%d at index %s\n", value, index)
	}
}
