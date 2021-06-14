package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Breasts and Eggs",Author: "Mieko Kawakami", Year: "2020"},
		Book{ID: 2, Title: "Where the Wild Ladies Are",Author: "Aoko Matsuda", Year: "2020"},
		Book{ID: 3, Title: "Deacon King Kong",Author: "James McBride", Year: "2020"},
		Book{ID: 4, Title: "A Burning",Author: "Megha Majumdar", Year: "2020"},
		Book{ID: 5, Title: " I Hold a Wolf by the Ears",Author: "Laura van den Berg", Year: "2020"})


	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router)) //if has any error

}

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book // holding the book records


func getBooks(w http.ResponseWriter, r *http.Request){
	log.Println("getBooks")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"]) // _ for error
	log.Println(params)

	for _, book := range(books){
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request){
	log.Println("addBook")
	var book Book
	json.NewEncoder(r.Body).Encode(&book)
	books = append(books,book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	log.Println("updateBook")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books{
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request){
	log.Println("removeBook")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"]) // _ for error

	for i, item := range books{
		if item.ID == book.ID {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
