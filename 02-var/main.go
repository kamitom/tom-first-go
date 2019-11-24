package main

import "fmt"
import "syreclabs.com/go/faker"

func main() {

	var fakerNameTest = faker.Name().String
	// Using var
	// var name = "TOM"
	fmt.Println("Hello, ", fakerNameTest())
}
