package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

func initDB() {
	var err error

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		);
		`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
	fmt.Println("Database initialized successfully")
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	book := Book{Title: title}
	json.NewDecoder(r.Body).Decode(&book)

	fmt.Fprintf(w, "Book created: %+v\n", book)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	book := Book{Title: title, Author: "Unknown"}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	book := Book{Title: title}
	json.NewDecoder(r.Body).Decode(&book)

	fmt.Fprintf(w, "Book updated: %+v\n", book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "Book deleted %s\n", title)
}

func main() {
	r := mux.NewRouter()

	//crud
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//server run
	http.ListenAndServe(":8080", r)
}
