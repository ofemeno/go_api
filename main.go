package main

import (
	"fmt"
	"log"
	"net/http"
)

// form handler function

func formHandler(w http.ResponseWriter, r *http.Request) {
	// check for request errors
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParsForm() err: %v", err)
		return
	}

	fmt.Fprint(w, "POST request successful\n\n")

	// assign form data for variables
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

// hello handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// validate routes
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// validate post method
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	// serving the static webpages
	fileServer := http.FileServer(http.Dir("./static"))

	// routes handler
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
