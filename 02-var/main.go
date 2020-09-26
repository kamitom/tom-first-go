package main

import (
	"fmt"

	"syreclabs.com/go/faker"
)

func main() {

	// Using var
	var fakerNameTest = faker.Name().FirstName()
	var fullName = faker.Name()
	// var name = "TOM: " + faker.RandomString(10)

	// var ageTest int64 = 100
	const isCool = true
	// isCool = false

	// Shorthand
	nameShorthand := "tom"
	// email := "tom@test.com"
	size := 1.238
	// var size float32 = 1.238

	name, email, newfakername := fullName, fakerNameTest+"@test.com", fakerNameTest

	fmt.Println("Hello shorthand name, ", nameShorthand)

	// fmt.Println(name, ageTest, isCool, size, email)

	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

	fmt.Println(name, email, newfakername)

	fmt.Println(HelloWorld(fakerNameTest))

}

func HelloWorld(user_name string) string {
	return fmt.Sprintf("Hi, %s ", user_name)
}
