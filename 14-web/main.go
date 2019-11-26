package main

import (
	"fmt"
	"net/http"

	"syreclabs.com/go/faker"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fName := faker.Name()
	testHTML := "<h1>Hello " + fName.Name() + "</h1>"

	fmt.Fprintf(w, testHTML)
}
func aboutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About ME</h1>")
}

func main() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/about", aboutFunc)
	fmt.Println("server starting...")
	http.ListenAndServe(":3003", nil)
}
