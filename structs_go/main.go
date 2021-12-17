package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}


func main() {
	jim := person {
		firstName: "rajendra",
		lastName: "arora",
		contactInfo: contactInfo {
			email: "rajendra@gmail.com",
			zipCode: 324007,
		},
	}


	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(firstNewName string) {
	(*pointerToPerson).firstName = firstNewName
}


func (p person) print() {
	fmt.Printf("%+v", p)
}
