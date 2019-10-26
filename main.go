package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init Book var as a slice Book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "Shadab", Lastname: "Khan"}})
	books = append(books, Book{ID: "2", Isbn: "448734", Title: "Book Two", Author: &Author{Firstname: "Adarsh", Lastname: "Theja"}})

	//Router Handler /EndPoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
