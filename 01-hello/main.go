package main

import "fmt"

import "syreclabs.com/go/faker"

func main() {
	fmt.Println("hello, tom go app2 !!")

	username, time := test()
	fmt.Println("test multiple values return")
	fmt.Println(username, time)

}

// multiple value
func test() (string, int) {
	fName := faker.Name().FirstName
	fNumber := faker.RandomInt(10,100)
	
	return fName(), fNumber
}
