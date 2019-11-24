package main

import "fmt"

func main() {

	//todo: https://youtu.be/SqrbIlUwR0U
	fmt.Println("hello, maps, mean key value")

	// //Define map
	// emails := make(map[string]string)

	// //Assign key value
	// emails["tom"] = "tom@test.com"
	// emails["hellen"] = "hellen@test.com"
	// emails["peter"] = "peter@test.com"

	//Declare and add kv
	emails := map[string]string{"tom": "tom@p1.com", "hellen": "hellen@p1.com", "jack": "jack@p1.com"}
	emails["mike"] = "mike@p1.com"

	NBATeam := make(map[string]int)
	data := make(map[string]string)

	data["name"] = "james"
	data["password"] = "2019 spring"

	NBATeam["LA"] = 109
	NBATeam["NY"] = 123

	fmt.Println(data["password"])

	fmt.Printf("NBA team LA score: %d\n", NBATeam["LA"])

	fmt.Println(emails["hellen"])

	fmt.Printf("show me email type: %T\n", emails)
	fmt.Println(len(emails))

	//Delete from map
	delete(emails, "tom")
	fmt.Println(emails)
}
