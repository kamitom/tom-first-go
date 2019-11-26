package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	// firstName string
	// lastName  string
	// age       int
	// city      string
	// gender    string

	firstName, lastName, city, gender string
	age                               int
}

// greet method (value reciever)
func (p Person) greet() string {
	return "hello, " + p.firstName + " " + p.lastName + " you are now " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLaseName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLaseName
	}
}

func main() {

	//Init person using struct
	person1 := Person{firstName: "hellen", lastName: "smith", age: 25, city: "Boston", gender: "f"}
	// person1 := Person{"tom", "smith", 25, "Boston", "f"}

	fmt.Println(person1)
	fmt.Println(person1.city)
	person1.age++

	fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()

	person1.getMarried("Williams")

	fmt.Println(person1.greet())

	fmt.Println("hello structs!!")
}
