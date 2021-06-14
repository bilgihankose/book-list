package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init()  {
	gotenv.Load()
}

func logFatal(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	/* books = append(books,
		Book{ID: 1, Title: "Breasts and Eggs",Author: "Mieko Kawakami", Year: "2020"},
		Book{ID: 2, Title: "Where the Wild Ladies Are",Author: "Aoko Matsuda", Year: "2020"},
		Book{ID: 3, Title: "Deacon King Kong",Author: "James McBride", Year: "2020"},
		Book{ID: 4, Title: "A Burning",Author: "Megha Majumdar", Year: "2020"},
		Book{ID: 5, Title: " I Hold a Wolf by the Ears",Author: "Laura van den Berg", Year: "2020"})


	 */


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
var db *sql.DB


func getBooks(w http.ResponseWriter, r *http.Request){
	var book Book
	books = []Book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request){
	var book Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(books)
}

func addBook(w http.ResponseWriter, r *http.Request){
	var book Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books (title, author, year) values ($1,$2,$3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id - $1", params["id"])
	log.Fatal(err)

	rowsDeleted, err := result.RowsAffected()
	log.Fatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)

}
