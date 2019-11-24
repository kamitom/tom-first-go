package main

import "fmt"

func greetings(name string) string {

	return "Hello " + name
}

func getSum(n1 ,n2 int) int {
	return n1 + n2
}

func main() {

	fmt.Println("Hello, " + greetings("TOM"))

	fmt.Println(getSum(100, 200))
}
