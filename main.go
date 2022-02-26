package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"net/http"
)

var DB *sql.DB

type Book struct {
	Title    string
	Author   string
	Category string
	Price    float32
}

func main() {
	log.Print("Main stared...")

	defer DB.Close()
	http.HandleFunc("/getbooks", ShowBooks)
	http.ListenAndServe(":8484", nil)

}

func ShowBooks(w http.ResponseWriter, req *http.Request) {
	books, err := GetAllbooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	for _, i := range books {
		fmt.Fprintf(w, " %s, %s, %s,INR%.2f\n", i.Title, i.Author, i.Category, i.Price)
	}
}

func GetAllbooks() ([]Book, error) {
	log.Debug().Msg("GetAllBooks function is running..")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "password", "shopdb")
	DB, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal().Msgf("error - %v \n", err)
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal().Msgf("Error %v - ", err)
	}
	defer rows.Close()
	var books []Book

	for rows.Next() {
		var bk Book
		err := rows.Scan(&bk.Title, &bk.Author, &bk.Category, &bk.Price)
		if err != nil {
			log.Fatal().Msgf("Error - %v ", err)
		}
		books = append(books, bk)
	}

	if err = rows.Err(); err != nil {
		log.Fatal().Msgf("error - %v", err)
	}
	return books, nil
}
