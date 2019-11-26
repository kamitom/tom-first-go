package main

import "fmt"

func main() {

	fmt.Println("everything in go is passed by value!!")

	a := 5

	// Assigning b to a pointer of a
	// (it is a MEMORY Address)
	// (a value stored on the system)
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)  // *int - the STAR represents a pointer

	// use * to read a value from memory address
	fmt.Println(*b)  // should get 5 
	fmt.Println(*&a)  // should get 5 

	// Change the value of a pointer
	*b = 10
	fmt.Println(a)


}
