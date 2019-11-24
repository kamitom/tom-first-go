package main

import "fmt"

func main() {

	i := 1
	// long method
	for i < 10 {
		fmt.Printf("print %d\n", i)
		i++
	}

	fmt.Println("loops stop.")

	// short method
	for j := 1; j < 10; j++ {
		fmt.Printf("number is %d\n", j)
	}
	println("j loops stop.")

	// FizzBuzz
	for ii := 1; ii <= 100; ii++ {
		if ii%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if ii%3 == 0 {
			fmt.Println("Fizz")
		} else if ii%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(ii)
		}
	}

}
