package main

import (
	"fmt"
	"net/http"
)

func simpleHttp() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested %s", r.URL.Path)
	}

	booksHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested some books from path: %s", r.URL.Path)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/books", booksHandler)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	simpleHttp()
	// utils.Run()
}
