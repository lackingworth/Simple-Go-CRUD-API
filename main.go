package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Test slice
var books []Book

// Handler functions
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _,v := range books {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

// **Not to use with dbs**
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, v := range books {
		if v.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = params["id"]
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for i, v := range books {
		if v.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Director: &Director{FirstName: "Alex", LastName: "Aston"}})
	books = append(books, Book{ID: "2", Isbn: "45455", Title: "Book Two", Director: &Director{FirstName: "John", LastName: "Boston"}})
	books = append(books, Book{ID: "3", Isbn: "486673", Title: "Book Three", Director: &Director{FirstName: "Hinata", LastName: "Kojima"}})

	r.HandleFunc("/books", getBooks).Methods("GET") // Get all the books from db/test json
	r.HandleFunc("/books/{id}", getBook).Methods("GET") // Get book by id
	r.HandleFunc("/books", createBook).Methods("POST") // Create new book entry
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT") // Update exsisting book entry
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE") // Delete book entry

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}