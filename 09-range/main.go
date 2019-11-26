package main

import "fmt"

func main() {

	fmt.Println("hello, range!!")

	ids := []int{1, 22, 33, 55, 29, 88, 123, 92}

	// loop through ids
	fmt.Println("loop through ids")
	for idKey, idValue := range ids {
		fmt.Printf("%d - ID: %d\n", idKey, idValue)
	}

	// not using index
	for _, idKeyValue := range ids {
		fmt.Printf("ID %d\n", idKeyValue)
	}

	// Add ids together
	fmt.Println("Add ids together!")
	Sum := 0
	for _, idKeyValue2 := range ids {
		Sum += idKeyValue2
	}
	fmt.Println("Sum:", Sum)

	// Range with map
	//Declare and add kv
	fmt.Println("Range with map")
	emails := map[string]string{"tom": "tom@p1.com", "hellen": "hellen@p1.com", "jack": "jack@p1.com"}
	emails["mike"] = "mike@p1.com"

	for kData, valueData := range emails {
		fmt.Printf("kye is %s: %s\n", kData, valueData)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}

}
