package main

import (
	"fmt"

	"syreclabs.com/go/faker"
)

func main() {
	// fmt.Println("hello, tom go app2 !!")
	fmt.Println("new branch dev is ready.")

	returnedUsername, returnedNumber := multipleReturnValues()
	fmt.Println("test multiple values return")
	fmt.Println(returnedUsername, returnedNumber)

}

// multiple return values
func multipleReturnValues() (r1 string, r2 int) {
	if r2 == 0 {
		r2 = 2
	}
	r1 = "ok, random number is"
	// fName := (faker.Name().FirstName)
	fFullName := (faker.Name().Name)
	fRandomNumber := faker.RandomInt(10, 300) * r2

	return fFullName() + " says: " + r1, fRandomNumber
}
