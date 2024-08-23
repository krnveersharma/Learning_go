package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}
type Address struct {
	house int
	area  string
	state string
}
type Contact struct {
	email string
	phone string
}
type Employee struct {
	personDetails Person
	personContact Contact
	personAddress Address
}

func main() {
	// //1st method
	// var karan Person
	// karan.firstName = "Karanveer"
	// karan.lastName = "Sharma"
	// karan.age = 21
	// fmt.Println((karan))
	// //2nd method
	// person1 := Person{
	// 	firstName: "person",
	// 	lastName:  "1",
	// 	age:       999,
	// }
	// fmt.Println(person1)
	// //3rd but it will use pointer
	// var person2 = new(Person)
	// person2.firstName = "person"
	// person2.lastName = "2"
	// person2.age = 99
	// fmt.Println(person2)
	var employee1 Employee
	employee1.personDetails = Person{
		firstName: "Karanveer",
		lastName:  "Sharma",
		age:       21,
	}
	employee1.personContact = Contact{
		email: "karanveer1029y@gmail.com",
		phone: "9XXXXXXXX8",
	}
	employee1.personAddress = Address{
		house: 1,
		area:  "jjk",
		state: "Punjab",
	}
	fmt.Println(employee1)
	fmt.Println(employee1.personAddress)

}
