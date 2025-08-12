package main

import (
	"fmt"
	"net/http"

	"github.com/emekaony/backend_programming/src/utils"
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

	idx := 10

	for idx > 0 {
		fmt.Println("Index:", idx)
		idx--
	}

	result := utils.Add(1, 2)
	fmt.Println(result)
}
