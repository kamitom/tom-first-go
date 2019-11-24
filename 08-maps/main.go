package main

import "fmt"

func main() {

	//todo: https://youtu.be/SqrbIlUwR0U
	fmt.Println("hello, maps, mean key value")

	data := make(map[string]string)
	NBATeam := make(map[string]int)
	data["name"] = "james"
	data["password"] = "2019 spring"

	NBATeam["LA"] = 109
	NBATeam["NY"] = 123

	fmt.Println(data["password"])

	fmt.Printf("NBA team LA score: %d\n", NBATeam["LA"])
}
