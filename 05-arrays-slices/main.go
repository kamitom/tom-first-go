package main

import "fmt"

func main() {

	// //Array
	// var fruitArr [2]string

	// //Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Banana"

	//Declare and assing
	fruitArr := [3]string{"Apple", "Orange", "Lenmon"}

	//Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println("Hello", fruitArr)
	fmt.Println("Hello", fruitArr[1])
	fmt.Printf("%T\n", fruitArr)
	fmt.Println("Hello", fruitSlice)
	fmt.Printf("%T\n", fruitSlice)
	fmt.Println(fruitSlice[1:3]) //range: [Orange Grape] start at 1(Orange), stop at 3(Cherry)

}
