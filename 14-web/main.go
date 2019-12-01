package main

import (
	"fmt"
	"net/http"

	"syreclabs.com/go/faker"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fName := faker.Name()
	fNaumber := faker.RandomInt(10, 1000)
	// testHTML := `<h1>Hello` + fName.Name()  + `</h1>`

	// try backticks
	testHTML2 := fmt.Sprintf(`
	<h1>
	Hello %s
	<hr>
	random number: %d
	</h1>
	`, fName.Name(), fNaumber)

	fmt.Fprintf(w, testHTML2)
}
func aboutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About ME</h1>")
}

func main() {

	exposePort := ":8633"

	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/about", aboutFunc)
	fmt.Println("server starting...")
	http.ListenAndServe(exposePort, nil)
}
