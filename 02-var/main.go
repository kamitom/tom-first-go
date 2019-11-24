package main

import (
	"fmt"

	"syreclabs.com/go/faker"
)

func main() {

	// Using var
	var fakerNameTest = faker.Name().String
	// var name = "TOM: " + faker.RandomString(10)

	var ageTest int64 = 100
	const isCool = true
	// isCool = false

	// Shorthand
	// name := "tom"
	// email := "tom@test.com"
	size := 1.238
	// var size float32 = 1.238

	name, email := "jack", "jack@test.com"

	fmt.Println("Hello, ", fakerNameTest())

	fmt.Println(name, ageTest, isCool, size, email)

	fmt.Printf("%T\n", fakerNameTest)
	fmt.Printf("%T\n", size)
}
