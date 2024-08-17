package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Define connection string
	connStr := "user=postgres dbname=wbapp password=password sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int32
		var userid int32
		var title string
		var body string

		err := rows.Scan(&id, &userid, &title, &body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id: %d User Id: %d Title: %s Body: %s", id, userid, title, body)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
