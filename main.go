package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
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
	fmt.Println("getBooks")
}

func getBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("getBook")
}

func addBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("addBook")
}

func updateBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("updateBook")
}

func removeBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("removeBook")
}
