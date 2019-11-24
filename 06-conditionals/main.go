package main

import "fmt"

func main() {

	// n1 := 5
	// n2 := 7

	x, y := 1010, 1010

	// if
	if x <= y {
		fmt.Printf("%d is less than or equal %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "green"

	if color == "Red" {
		fmt.Println("color is red")
	} else if color == "Blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT red or blue")
	}

	// switch
	switch color {
	case "Red":
		fmt.Println("color is red")
	case "Blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT blue or red")
	}

}
