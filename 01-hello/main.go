package main

import (
	"fmt"

	"syreclabs.com/go/faker"
)

func main() {
	fmt.Println("hello, tom go app2 !!")

	username, time := test()
	fmt.Println("test multiple values return")
	fmt.Println(username, time)

}

// multiple return values
func test() (r1 string, r2 int) {
	if r2 == 0 {
		r2 = 2
	}
	r1 = "ok test me"
	fName := (faker.Name().FirstName)
	fNumber := faker.RandomInt(10, 100) * r2

	return fName() + " : " + r1, fNumber
}
