package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You requested %s", r.URL.Path)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You requested some books from path: %s", r.URL.Path)
}

func main() {
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/books", booksHandler)

	// fmt.Println("Starting server on port 8080")
	// http.ListenAndServe(":8080", nil)

	mp := map[string]int{
		"Emeka": 23,
	}
	mp["chidera"] = 34
	fmt.Println(mp)
}
