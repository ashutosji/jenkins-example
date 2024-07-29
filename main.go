package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// Default router when application will start
	r.HandleFunc("/", handler)

	// This is where the router is useful, it allows us to declare methods that
	// this path will be valid for
	r.HandleFunc("/hello", handlers)

	// We can then pass our router (after declaring all our routes) to this method
	// (where previously, we were leaving the secodn argument as nil)
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting!")
}
