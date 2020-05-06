package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome!</h1>")
	fmt.Fprintf(w, "<p>This is the main page</p>")
}

func main() {
	fmt.Println("Server has now been started")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
