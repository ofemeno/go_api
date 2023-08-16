package main

import (
	"fmt"
	"log"
	"net/http"
)

// hello handler function
func hellohandler(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	// serving the static webpages
	fileServer := http.FileServer(http.Dir("./static"))

	// routes handler
	http.HandleFunc("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
