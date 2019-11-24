package main

import (
	"fmt"

	"math"

	"syreclabs.com/go/faker"

	"github.com/kamitom/tom-first-go/03-packages/strutil"
)

func main() {

	var testme = math.Abs(-12.223)
	testme2 := math.Floor(2.7) //2
	// testme2 = 1000
	testme3 := math.Ceil(2.7) //3

	strReverse := strutil.Reverse("TOMTESTHERE")

	fmt.Println("Hello World: ", faker.Name(), testme, testme2, testme3, strReverse)
}
